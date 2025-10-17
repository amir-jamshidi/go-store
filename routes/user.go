package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"room-reserve/db"
	"room-reserve/models"
	"room-reserve/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type SendOTPRequest struct {
	Phone string `json:"phone" binding:"required" validate:"required"`
}

type VerifyOTPRequest struct {
	Phone string `json:"phone" binding:"required" validate:"required"`
	Code  string `json:"code" binding:"required" validate:"required"`
}

var validate = validator.New()

func verifyOTP(c *gin.Context) {
	var req VerifyOTPRequest

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request " + err.Error(),
		})
		return
	}

	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := db.GetCollection("users")
	otpCollection := db.GetCollection("otps")

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"phone": req.Phone}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var otp models.OTP
	err = otpCollection.FindOne(ctx, bson.M{"user_id": user.ID}).Decode(&otp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OTP not found or expired"})
		return
	}

	if time.Now().After(otp.ExpireAt) {
		otpCollection.DeleteOne(ctx, bson.M{"_id": otp.ID})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OTP has expired",
		})
		return
	}

	if otp.TryCount >= 3 {
		otpCollection.DeleteOne(ctx, bson.M{"user_id": user.ID})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Too many attempts. Please request a new OTP",
		})
		return
	}

	if otp.Code != req.Code {
		otpCollection.UpdateOne(
			ctx,
			bson.M{"_id": otp.ID},
			bson.M{"$inc": bson.M{"try_count": 1}},
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP code"})
		return
	}

	otpCollection.DeleteOne(ctx, bson.M{"_id": otp.ID})

	token, err := utils.GenerateJWT(user.ID.Hex(), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	userCollection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{"updated_at": time.Now()}},
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}

func sendOTP(ctx *gin.Context) {
	var req SendOTPRequest

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request " + err.Error(),
		})
		return
	}

	if err := validate.Struct(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed " + err.Error(),
		})
		return
	}

	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
	// 	return
	// }

	if !isValidPhoneNumber(req.Phone) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	contx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollections := db.GetCollection("users")
	otpCollections := db.GetCollection("otps")

	var user models.User

	err := userCollections.FindOne(contx, bson.M{"phone": req.Phone}).Decode(&user)

	if err != nil {
		newUser := models.User{
			Phone:     req.Phone,
			Role:      "user",
			Score:     0,
			CreateAt:  time.Now(),
			UpdatedAt: time.Now(),
		}

		result, err := userCollections.InsertOne(ctx, newUser)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		user.ID = result.InsertedID.(bson.ObjectID)
		user.Phone = newUser.Phone
	}

	otpCode := generateOTP()

	otpCollections.DeleteMany(contx, bson.M{"user_id": user.ID})

	newOTP := models.OTP{
		Code:      otpCode,
		UserID:    user.ID,
		TryCount:  0,
		ExpireAt:  time.Now().Add(5 * time.Minute),
		CreatedAt: time.Now(),
	}

	_, err = otpCollections.InsertOne(contx, newOTP)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save OTP"})
		return
	}

	fmt.Printf("OTP for %s: %s \n", req.Phone, otpCode)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP sned successfully",
		"otp":     otpCode,
	})
}

func isValidPhoneNumber(phone string) bool {
	matched, _ := regexp.MatchString(`^09[0-9]{9}$`, phone)
	return matched
}

func generateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

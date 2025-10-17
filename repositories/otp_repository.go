package repositories

import (
	"context"
	"room-reserve/db"
	"room-reserve/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OTPRepository struct {
	collection *mongo.Collection
}

func NewOTPRepository() *OTPRepository {
	return &OTPRepository{
		collection: db.GetCollection("otps"),
	}
}

func (r *OTPRepository) FindByUserID(ctx context.Context, userID bson.ObjectID) (*models.OTP, error) {
	var otp models.OTP
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&otp)
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *OTPRepository) Create(ctx context.Context, otp *models.OTP) error {
	result, err := r.collection.InsertOne(ctx, otp)
	if err != nil {
		return err
	}
	otp.ID = result.InsertedID.(bson.ObjectID)
	return nil
}

func (r *OTPRepository) DeleteByUserID(ctx context.Context, userID bson.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"user_id": userID})
	return err
}

func (r *OTPRepository) Delete(ctx context.Context, otpID bson.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": otpID})
	return err
}

func (r *OTPRepository) IncrementTryCount(ctx context.Context, otpID bson.ObjectID) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": otpID},
		bson.M{"$inc": bson.M{"try_count": 1}},
	)
	return err
}

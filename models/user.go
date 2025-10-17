package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type SellerInfo struct {
	StoreName       string    `json:"store_name" bson:"store_name"`
	NationalCode    string    `json:"national_code" bson:"national_code"`
	BankAccount     string    `json:"bank_account" bson:"bank_account"`
	IBAN            string    `json:"iban" bson:"iban"`
	Address         string    `json:"address" bson:"address"`
	IsVerified      bool      `json:"is_verified" bson:"is_verified"`
	VerifiedAt      time.Time `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	BusinessLicense string    `json:"business_license" bson:"business_license"`
	IsActive        bool      `json:"is_active" bson:"is_active"`
}

type User struct {
	ID         bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Phone      string        `json:"phone" bson:"phone" binding:"required"`
	Fullname   string        `json:"fullname" bson:"fullname"`
	Role       string        `json:"role" bson:"role"`
	Score      int           `json:"score" bson:"score"`
	CreateAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
	SellerInfo *SellerInfo   `json:"seller_info,omitempty" bson:"seller_info,omitempty"`
}

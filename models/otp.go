package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OTP struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Code      string        `json:"code" bson:"code"`
	UserID    bson.ObjectID `json:"user_id" bson:"user_id"`
	TryCount  int           `json:"try_count" bson:"try_count"`
	ExpireAt  time.Time     `json:"expire_at" bson:"expire_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

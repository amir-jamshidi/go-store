package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Phone     string        `json:"phone" bson:"phone" binding:"required"`
	Fullname  string        `json:"fullname" bson:"fullname"`
	Role      string        `json:"role" bson:"role"`
	Score     int           `json:"score" bson:"score"`
	CreateAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

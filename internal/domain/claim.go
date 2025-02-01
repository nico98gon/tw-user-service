package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	ID       primitive.ObjectID 	`bson:"_id" json:"_id,omitempty"`
	Email    string              	`bson:"email" json:"email"`
	jwt.RegisteredClaims
}
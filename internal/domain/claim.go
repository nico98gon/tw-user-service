package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	ID       primitive.ObjectID 	`json:"_id,omitempty" bson:"_id"`
	Email    string              	`json:"email" bson:"email"`
	jwt.RegisteredClaims
}
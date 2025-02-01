package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       		primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Name     		string 							`bson:"name" json:"name,omitempty"`
	LastName 		string 							`bson:"last_name" json:"last_name,omitempty"`
	Email    		string 							`bson:"email" json:"email"`
	Birthdate 	time.Time 					`bson:"birthdate" json:"birthdate,omitempty"`
	Password 		string 							`bson:"password" json:"password,omitempty"`
	Avatar 			string 							`bson:"avatar" json:"avatar,omitempty"`
	Banner			string							`bson:"banner" json:"banner,omitempty"`
	Bio 				string 							`bson:"bio" json:"bio,omitempty"`
	WebSite 		string 							`bson:"web_site" json:"web_site,omitempty"`
	Location 		string 							`bson:"location" json:"location,omitempty"`
	CreatedAt 	time.Time 					`bson:"created_at" json:"created_at,omitempty"`
}
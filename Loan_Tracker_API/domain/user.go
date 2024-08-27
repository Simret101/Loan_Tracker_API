package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName   string             `json:"username" bson:"username"`
	Email      string             `json:"email" bson:"email"`
	Is_Admin   bool               `json:"is_admin" bson:"is_admin"`
	Password   string             `json:"password,omitempty" bson:"password,omitempty"`
	IsVerified bool               `json:"is_verified" bson:"is_verified"`
}

type ResponseUser struct {
	ID       string `json:"_id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type UpdateUser struct {
	UserName string `json:"username" bson:"username"`
	Bio      string `json:"bio,omitempty" bson:"bio,omitempty"`
}

type LogINUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type RegisterUser struct {
	UserName string `json:"username" bson:"username"`
	Bio      string `json:"bio,omitempty" bson:"bio,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type UpdatePassword struct {
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password"`
}

type VerifyEmail struct {
	Email string `json:"email" bson:"email"`
}

// from actual user model to response model to be done in usecase
func CreateResponseUser(user User) ResponseUser {
	return ResponseUser{
		ID:       user.ID.Hex(),
		UserName: user.UserName,
		Email:    user.Email,
	}
}
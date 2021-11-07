package user

import (
	"context"
)

var users = []User{
	{
		Email:    "oskarmikus@protonmail.com",
		Password: "asd",
	},
	{
		Email:    "osk@gmail.com",
		Password: "qwerty",
	},
	{
		Email:    "hehehe@pl",
		Password: "=(",
	},
}

func GetUserByEmail(ctx context.Context, email string) (result interface{}) {
	var user User
	for _, u := range users {
		if email == u.Email {
			user = u
		}
	}
	return user
}

func GetUserList(ctx context.Context, limit int) (result interface{}) {
	return users
}

func AddUser(ctx context.Context, user User) User {
	_ = append(users, user)
	return user
}

func UpdateUser(ctx context.Context, user User) {

}

func DeleteUser(ctx context.Context, email string) {

}

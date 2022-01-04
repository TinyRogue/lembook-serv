package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	UsersCollection *mongo.Collection
}

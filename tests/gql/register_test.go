package gql

import (
	"context"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	service "github.com/TinyRogue/lembook-serv/pkg/mongo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

func deleteUsers(t *testing.T) {
	uh := service.DB.Collection(service.UsersCollectionName)
	filter := bson.D{{"username", os.Getenv("TEST_REGISTER1")}}
	_, err := uh.DeleteOne(context.Background(), filter)
	if err != nil {
		t.Errorf("Could not clean up resource: %v", err.Error())
	}
}

func TestRegister(t *testing.T) {
	_ = godotenv.Load("../../.env")
	service2.InitDb()

	registerTests := []struct {
		testName string
		username string
		password string
		mess     string
	}{
		{testName: "OK", username: os.Getenv("TEST_REGISTER1"), password: os.Getenv("TEST_RPASS1"), mess: "user created"},
		{testName: "User exists", username: os.Getenv("TEST_REGISTER1"), password: os.Getenv("TEST_RPASS1"), mess: "[{\"message\":\"password does not meet its requirements\",\"path\":[\"register\"]}]"},
		{testName: "Insufficient password", username: os.Getenv("TEST_REGISTER2"), password: os.Getenv("TEST_RPASS2"), mess: "[{\"message\":\"password does not meet its requirements\",\"path\":[\"register\"]}]"},
		{testName: "Invalid email", username: os.Getenv("TEST_REGISTER3"), password: os.Getenv("TEST_RPASS3"), mess: "[{\"message\":\"password does not meet its requirements\",\"path\":[\"register\"]}]"},
	}

	q := `
  		mutation register($register: Registration!) {
    		register(input: $register) {
      			res
    		}
  		}`

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))
	var res map[string]map[string]interface{}

	for i, tt := range registerTests {
		t.Run(tt.testName, func(t *testing.T) {
			registerData := client.Var("register", model.Registration{
				Username: registerTests[i].username,
				Password: registerTests[i].password,
			})
			err := c.Post(q, &res, registerData)
			t.Log(res["register"]["res"])
			if err != nil && err.Error() != tt.mess {
				t.Errorf("Wanted: %s, Got: %s", tt.mess, err.Error())
				return
			}

			if ok := res["register"]["res"] == tt.mess; !ok {
				t.Errorf("Wanted: %s, Got: %s", tt.mess, res["register"]["res"])
			}
		})
	}

	t.Cleanup(func() {
		deleteUsers(t)
	})
}

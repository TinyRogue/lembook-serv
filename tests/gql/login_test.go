package gql

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/TinyRogue/lembook-serv/src/cmd/gql/graph"
	"github.com/TinyRogue/lembook-serv/src/cmd/gql/graph/generated"
	"github.com/TinyRogue/lembook-serv/src/cmd/gql/graph/generated/model"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	_ = godotenv.Load("../../.env")
	loginTests := []struct {
		testName string
		username string
		password string
		resMess  string
	}{
		{testName: "valid credentials 1", username: os.Getenv("TEST_LOGIN1"), password: os.Getenv("TEST_PASS1"), resMess: "Could not get writer"},
		{testName: "invalid credentials 1", username: os.Getenv("TEST_LOGIN2"), password: os.Getenv("TEST_PASS2"), resMess: "Could not get writer"},
		{testName: "valid credentials 2", username: os.Getenv("TEST_LOGIN3"), password: os.Getenv("TEST_PASS3"), resMess: "Could not get writer"},
		{testName: "valid credentials 3", username: os.Getenv("TEST_LOGIN4"), password: os.Getenv("TEST_PASS4"), resMess: "Could not get writer"},
		{testName: "invalid credentials 1", username: os.Getenv("TEST_LOGIN5"), password: os.Getenv("TEST_PASS5"), resMess: "Could not get writer"},
	}

	q := `
  		mutation login($login: Login!) {
    		login(input: $login) {
      			res
    		}
  		}`

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))
	var res string

	for i, tt := range loginTests {
		t.Run(tt.testName, func(t *testing.T) {
			loginData := client.Var("login", model.Login{
				Username: loginTests[i].username,
				Password: loginTests[i].password,
			})

			err := c.Post(q, &res, loginData)
			if ok := err == nil; !ok {

			} else {
				t.Error("Should fail due to internal server error")
			}
		})
	}
}

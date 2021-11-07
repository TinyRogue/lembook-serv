package main

import (
	"github.com/TinyRogue/lembook-serv/gql/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Post_t struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("Unexpected errors: %v\n", result.Errors)
	}
	return result
}

func main() {
	godotenv.Load()

	var mode string
	if len(os.Args) > 1 && os.Args[1] == "--prod" {
		mode = "release"
	} else {
		mode = "debug"
	}

	gin.SetMode(mode)
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/graphql", func(ctx *gin.Context) {
		res := executeQuery(ctx.Request.URL.Query().Get("query"), user.Schema)
		ctx.JSON(http.StatusOK, res)
	})

	router.POST("/graphql", func(ctx *gin.Context) {
		var postData Post_t
		if ctx.BindJSON(&postData) != nil {
			ctx.JSON(http.StatusBadRequest, "Improper request body.")
			return
		}

		res := graphql.Do(graphql.Params{
			Context:        ctx.Request.Context(),
			Schema:         user.Schema,
			RequestString:  postData.Query,
			VariableValues: postData.Variables,
			OperationName:  postData.Operation,
		})
		ctx.JSON(http.StatusOK, res)
	})

	err := router.Run()
	if err != nil {
		log.Fatalln("Couldn't start the server: ", err)
	}
}

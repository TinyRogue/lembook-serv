package user

import (
	"context"
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        UserType,
				Description: "Get user",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					email, ok := p.Args["email"].(string)
					if ok {
						result = GetUserByEmail(context.Background(), email)
					}
					return result, nil
				},
			},

			"list": &graphql.Field{
				Type:        graphql.NewList(UserType),
				Description: "Get users list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					limit, _ := params.Args["limit"].(int)
					result = GetUserList(context.Background(), limit)
					return result, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        UserType,
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user := User{
					Email:    params.Args["email"].(string),
					Password: params.Args["password"].(string),
				}
				AddUser(context.Background(), user)
				return user, nil
			},
		},

		"update": &graphql.Field{
			Type:        UserType,
			Description: "Update user by email",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user := User{}
				if email, emailOk := params.Args["email"].(string); emailOk {
					user.Email = email
				}

				if password, passwordOk := params.Args["password"].(string); passwordOk {
					user.Password = password
				}

				UpdateUser(context.Background(), user)
				return user, nil
			},
		},

		"delete": &graphql.Field{
			Type:        UserType,
			Description: "Delete user by email",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				email, _ := params.Args["email"].(string)
				DeleteUser(context.Background(), email)
				return email, nil
			},
		},
	},
})

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

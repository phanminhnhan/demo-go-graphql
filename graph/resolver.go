package graph

import (
	"graphql_postrgres/database"
	"graphql_postrgres/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	LinkRepo database.LinkRepo
	UserRepo database.UserRepo
	Links []*model.Link
	Users []*model.User
}

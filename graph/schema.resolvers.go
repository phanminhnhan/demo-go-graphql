package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql_postrgres/graph/generated"
	"graphql_postrgres/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	return nil, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	return links, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var users = []*model.User{
	{ID: "1",
		Username: "nhan1",
		Email:    "pw1",
		Link:     links,
	},
	{ID: "2",
		Username: "nhan2",
		Email:    "pw2",
	},
	{ID: "3",
		Username: "nhan3",
		Email:    "pw3",
	},
}
var links = []*model.Link{
	{
		ID:   "1",
		Name: "link1",
		Desc: "desc1",
	},
	{
		ID:   "2",
		Name: "link2",
		Desc: "desc2",
	},
	{
		ID:   "3",
		Name: "link3",
		Desc: "desc3",
	},
}

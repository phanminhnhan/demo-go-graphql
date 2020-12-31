package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphql_postrgres/graph/generated"
	"graphql_postrgres/graph/model"
)

func (r *mutationResolver) UpdateLink(ctx context.Context, id string, input model.UpdateLink) (*model.Link, error) {
	link, err := r.LinkRepo.GetbyId(id)
	if err != nil || link == nil {
		return nil, errors.New("link is not exist")
	}
	updated := false
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is too short to update")
		}
		link.Name = *input.Name
		updated = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is too short to update")
		}
		link.Description = *input.Description
		updated = true
	}
	if !updated {
		return nil, errors.New("not updated yet ")
	}
	link, err = r.LinkRepo.Updatelink(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	if len(input.Name) < 5 {
		return nil, errors.New("name too short")
	}
	if len(input.Description) < 5 {
		return nil, errors.New("description too short too short")
	}

	newLink := &model.Link{
		Name:        input.Name,
		Description: input.Description,
		Userid:      input.Userid,
	}
	return r.LinkRepo.CreateLink(newLink)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	if len(input.Username) < 5 {
		return nil, errors.New("username too short")
	}
	if len(input.Email) < 5 {
		return nil, errors.New("description too short too short")
	}
	newUser := &model.User{
		Username: input.Username,
		Email:    input.Email,
	}
	return r.UserRepo.CreateUser(newUser)
}

func (r *mutationResolver) DeleteLink(ctx context.Context, id string) (bool, error) {
	link, err := r.LinkRepo.GetbyId(id)
	if err != nil || link == nil {
		return false, errors.New("link is not exist")
	}
	err = r.LinkRepo.DeleteLink(link)
	if err != nil {
		return false, errors.New("not deleted yet")
	}
	return true, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	return r.LinkRepo.GetAllLinks()
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserRepo.GetAllUsers()
}

func (r *queryResolver) UserLinks(ctx context.Context) ([]*model.UserLinks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetuserbyID(ctx context.Context, id string) (*model.User, error) {
	return r.UserRepo.GetUserById(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

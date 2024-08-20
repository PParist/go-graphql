package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	handler "github.com/PParist/go-graphql/graph/handlers"
	"github.com/PParist/go-graphql/graph/model"
)

// CreatePlayer is the resolver for the createPlayer field.
func (r *mutationResolver) CreatePlayer(ctx context.Context, req model.NewPlayer) (*model.Player, error) {
	return handler.CreatePlayer(&req), nil
}

// GetPlayers is the resolver for the getPlayers field.
func (r *queryResolver) GetPlayers(ctx context.Context) ([]*model.Player, error) {
	return handler.GetPlayer(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

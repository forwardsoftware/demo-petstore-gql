package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"petstore-gql/graph/model"
)

// PetCreate is the resolver for the petCreate field.
func (r *mutationResolver) PetCreate(ctx context.Context, pet model.PetInput) (*model.Pet, error) {
	panic(fmt.Errorf("not implemented: PetCreate - petCreate"))
}

// Pets is the resolver for the pets field.
func (r *queryResolver) Pets(ctx context.Context, limit *int32) ([]*model.Pet, error) {
	panic(fmt.Errorf("not implemented: Pets - pets"))
}

// Pet is the resolver for the pet field.
func (r *queryResolver) Pet(ctx context.Context, id string) (*model.Pet, error) {
	panic(fmt.Errorf("not implemented: Pet - pet"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

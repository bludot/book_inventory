package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/bludot/tempmee/user/graph/generated"
	"github.com/bludot/tempmee/user/graph/model"
	"github.com/bludot/tempmee/user/internal/resolvers"
)

// UserAPI is the resolver for the userApi field.
func (r *apiInfoResolver) UserAPI(ctx context.Context, obj *model.APIInfo) (*model.UserAPI, error) {
	return resolvers.AnimeAPI(r.Config)
}

// Register is the resolver for the register field.
func (r *queryResolver) Register(ctx context.Context, input model.RegisterInput) (*model.User, error) {
	return resolvers.Register(ctx, r.UserService, input.Name, input.Email, input.Password)
}

// APIInfo is the resolver for the apiInfo field.
func (r *queryResolver) APIInfo(ctx context.Context) (*model.APIInfo, error) {
	return resolvers.APIInfo(r.Config)
}

// SignIn is the resolver for the signIn field.
func (r *queryResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.User, error) {
	return resolvers.SignIn(ctx, r.UserService, input.Email, input.Password)
}

// ApiInfo returns generated.ApiInfoResolver implementation.
func (r *Resolver) ApiInfo() generated.ApiInfoResolver { return &apiInfoResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type apiInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

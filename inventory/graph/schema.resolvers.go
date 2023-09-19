package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"

	"github.com/bludot/tempmee/inventory/graph/generated"
	"github.com/bludot/tempmee/inventory/graph/model"
)

// InventoryAPI is the resolver for the inventoryApi field.
func (r *apiInfoResolver) InventoryAPI(ctx context.Context, obj *model.APIInfo) (*model.InventoryAPI, error) {
	panic(fmt.Errorf("not implemented: InventoryAPI - inventoryApi"))
}

// APIInfo is the resolver for the apiInfo field.
func (r *queryResolver) APIInfo(ctx context.Context) (*model.APIInfo, error) {
	panic(fmt.Errorf("not implemented: APIInfo - apiInfo"))
}

// ApiInfo returns generated.ApiInfoResolver implementation.
func (r *Resolver) ApiInfo() generated.ApiInfoResolver { return &apiInfoResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type apiInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
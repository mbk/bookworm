package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mbk/bookworm/graph/generated"
	"github.com/mbk/bookworm/graph/model"
	"github.com/mbk/bookworm/graph/models"
)

func (r *mutationResolver) AddBook(ctx context.Context, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditBook(ctx context.Context, id *int, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context, search *string) ([]*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Authors(ctx context.Context, search *string) ([]*models.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Exists(ctx context.Context, input model.NewAuthor) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

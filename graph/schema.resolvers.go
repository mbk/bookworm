package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
    "context"

    "github.com/mbk/bookworm/graph/generated"
    "github.com/mbk/bookworm/graph/model"
    "github.com/mbk/bookworm/graph/models"
)

func (r *mutationResolver) AddBook(ctx context.Context, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
    db := models.FetchConnection()
    defer db.Close()

    //create book using input struct
    book := models.Book{
        Name:     input.Name,
        Category: input.Category,
    }

    book.Author = make([]*models.Author, len(author))

    for index, item := range author {
        book.Author[index] = &models.Author{Firstname: item.Firstname, Lastname: item.Lastname}
    }

    db.Create(&book)
    return &book, nil
}

func (r *mutationResolver) EditBook(ctx context.Context, id *int, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
  return nil, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id *int) ([]*models.Book, error) {
 return nil, nil

}

func (r *queryResolver) Books(ctx context.Context, search *string) ([]*models.Book, error) {
  return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


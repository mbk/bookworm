package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"log"

	"github.com/jinzhu/gorm"
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
		a := &models.Author{Firstname: item.Firstname, Lastname: item.Lastname}
		book.Author[index] = a
		log.Println(*a)
		db.Create(a)
	}

	db.Create(&book)
	return &book, nil
}

func (r *mutationResolver) EditBook(ctx context.Context, id *int, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	var book models.Book

	//find book based on ID
	db = db.Preload("Authors").Where("id = ?", *id).First(&book).Update("name", input.Name)
	if input.Category != "" {
		db.Update("category", *&input.Category)
	}

	//update author
	book.Author = make([]*models.Author, len(author))
	for index, item := range author {
		book.Author[index] = &models.Author{Firstname: item.Firstname, Lastname: item.Lastname}
	}

	db.Save(&book)

	return &book, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id *int) ([]*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	var book models.Book

	//fetch based on ID and delete
	db.Where("id = ?", *id).First(&book).Delete(&book)

	//preload and fetch all recipe
	var books []*models.Book
	db.Preload("Author").Find(&books)

	return books, nil
}

func (r *queryResolver) Books(ctx context.Context, search *string) ([]*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	var books []*models.Book

	//preload loads the author relationship into each book
	db.Preload("Author").Find(&books)

	return books, nil
}

func (r *queryResolver) Authors(ctx context.Context, search *string) ([]*models.Author, error) {
	db := models.FetchConnection()
	defer db.Close()

	var authors []*models.Author

	//preload loads the author relationship into each book
	db.Find(&authors)

	return authors, nil
}

func (r *queryResolver) Exists(ctx context.Context, input model.NewAuthor) (bool, error) {
	db := models.FetchConnection()
	defer db.Close()

	//auth := models.Author{Firstname: input.Firstname, Lastname: input.Lastname}
	log.Println(input)
	result := db.Where("firstname = ? AND lastname = ?", input.Firstname, input.Lastname).Find(&models.Author{})
	log.Println("result.Value = ")
	log.Println(result.Value)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}

func (r *mutationResolver) DeleteAllBooks(ctx context.Context, alsoAuthours *bool) (*bool, error) {
	db := models.FetchConnection()
	if *alsoAuthours {
		db.Exec("DELETE FROM authors;")
	}
	db.Exec("DELETE FROM books;")
	//success
	result := true
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

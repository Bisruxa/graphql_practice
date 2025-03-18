package main

import (
	"context"
	"github.com/machinebox/graphql"
)

// Book struct represents a book object
type Book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

const hasuraEndpoint = "http://localhost:8080/v1/graphql"

// FetchBooks retrieves all books from Hasura
func FetchBooks(ctx context.Context) ([]Book, error) {
	client := graphql.NewClient(hasuraEndpoint)

	query := `
		query {
			books {
				id
				name
				genre
			}
		}
	`

	req := graphql.NewRequest(query)
	req.Header.Set("Content-Type", "application/json")

	var responseData struct {
		Books []Book `json:"books"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return nil, err
	}
	return responseData.Books, nil
}

// FetchBookByID retrieves a single book by ID
func FetchBookByID(ctx context.Context, bookID string) (*Book, error) {
	client := graphql.NewClient(hasuraEndpoint)

	query := `
		query getBook($id: uuid!) {
			books(where: {id: {_eq: $id}}) {
				id
				name
				genre
			}
		}		
	`

	req := graphql.NewRequest(query)
	req.Var("id", bookID)
	req.Header.Set("Content-Type", "application/json")

	var responseData struct {
		Books []Book `json:"books"`
	}

	if len(responseData.Books) == 0 {
		return nil, nil // Return nil if no book found
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return nil, err
	}
	return &responseData.Books[0], nil
}

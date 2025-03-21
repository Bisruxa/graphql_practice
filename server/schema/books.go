package schema

import (
	"context"
	"github.com/machinebox/graphql"
	"fmt"
)

// Book struct represents a book object
type Book struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

const hasuraEndpoint = "http://localhost:8080/v1/graphql"
const hasuraAdminSecret = "mysecret"

// createGraphQLClient initializes the GraphQL client with proper headers.
func createGraphQLClient() *graphql.Client {
	client := graphql.NewClient(hasuraEndpoint)
	client.Log = func(s string) { println(s) } // Log for debugging
	return client
}

// FetchBooks retrieves all books from Hasura using a GraphQL query.
func FetchBooks(ctx context.Context) ([]Book, error) {
	client := createGraphQLClient()

	query := `
		query {
			Books {
				uuid
				title
				genre
			}
		}
	`

	req := graphql.NewRequest(query)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret) // Add admin secret

	var responseData struct {
		Books []Book `json:"books"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return nil, err
	}

	return responseData.Books, nil
}

// FetchBookByID retrieves a single book by its UUID from Hasura.
func FetchBookByID(ctx context.Context, bookID string) (*Book, error) {
	client := createGraphQLClient()

	query := `
		query getBook($uuid: uuid!) {
			Books(where: {uuid: {_eq: $uuid}}) {
				uuid
				title
				genre
			}
		}		
	`

	req := graphql.NewRequest(query)
	req.Var("uuid", bookID) // Adjusted to pass UUID
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret) // Add admin secret

	var responseData struct {
		Books []Book `json:"books"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return nil, err
	}

	if len(responseData.Books) == 0 {
		return nil, nil // Return nil if no book found
	}

	return &responseData.Books[0], nil
}

// InsertBook inserts a new book into the database using a GraphQL mutation.
func InsertBook(ctx context.Context, book Book) error {
	client := createGraphQLClient()

	mutation := `
		mutation insertBook($uuid: uuid!, $title: String!, $genre: String!) {
			insert_Books(objects: {uuid: $uuid, title: $title, genre: $genre}) {
				returning {
					uuid
					title
					genre
				}
			}
		}
	`

	req := graphql.NewRequest(mutation)
	req.Var("uuid", book.Uuid)
	req.Var("title", book.Title)
	req.Var("genre", book.Genre)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret) // Add admin secret

	var responseData struct {
		InsertBooks struct {
			Returning []Book `json:"returning"`
		} `json:"insert_Books"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return err
	}

	if len(responseData.InsertBooks.Returning) == 0 {
		return fmt.Errorf("failed to insert book")
	}

	return nil
}

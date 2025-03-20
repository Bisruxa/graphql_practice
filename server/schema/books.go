package schema

import (
	"context"
	"github.com/machinebox/graphql"
)

// Book struct represents a book object
type Book struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

const hasuraEndpoint = "http://localhost:8080/v1/graphql"
const hasuraAdminSecret = "mysecret"

// createGraphQLClient initializes a new GraphQL client
func createGraphQLClient() *graphql.Client {
	client := graphql.NewClient(hasuraEndpoint)
	client.Log = func(s string) { println(s) }
	return client
}

// InsertBook inserts a new book into the database
func InsertBook(ctx context.Context, book Book) error {
	client := createGraphQLClient()

	mutation := `
		mutation insertBooks($uuid: uuid!, $title: String!, $genre: String!) {
			insert_Books_one(object: {uuid: $uuid, title: $title, genre: $genre}) {
				uuid
			}
		}
	`

	req := graphql.NewRequest(mutation)
	req.Var("uuid", book.Uuid)
	req.Var("title", book.Title)
	req.Var("genre", book.Genre)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	return client.Run(ctx, req, nil)
}

// FetchBooks retrieves all books
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
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	var responseData struct {
		Books []Book `json:"Books"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return nil, err
	}
	return responseData.Books, nil
}

// FetchBookByID retrieves a single book by ID
func FetchBookByID(ctx context.Context, bookID string) (*Book, error) {
	client := createGraphQLClient()

	query := `
		query getBook($id: uuid!) {
			Books(where: {uuid: {_eq: $id}}) {
				uuid
				title
				genre
			}
		}		
	`

	req := graphql.NewRequest(query)
	req.Var("id", bookID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	var responseData struct {
		Books []Book `json:"Books"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return nil, err
	}

	if len(responseData.Books) == 0 {
		return nil, nil
	}

	return &responseData.Books[0], nil
}

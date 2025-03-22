package schema

import (
	"context"
	"github.com/machinebox/graphql"
	"fmt"
)

// Book struct represents a book object
type Book struct {
	Uuid     string `json:"uuid"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	AuthorID string `json:"author_id"` // Added for the foreign key reference
}

// Author struct represents an author object
type Author struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

const hasuraEndpoint = "http://localhost:8080/v1/graphql"
const hasuraAdminSecret = "mysecret"

// we use the graphql client to connect graphiql with hasura 
func createGraphQLClient() *graphql.Client {
	client := graphql.NewClient(hasuraEndpoint)
	client.Log = func(s string) { println(s) } 
	return client
}

// FetchBooks fetches a list of all books
func FetchBooks(ctx context.Context) ([]Book, error) {
	client := createGraphQLClient()

	query := `
		query {
			Books {
				uuid
				title
				genre
				author_id
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

// InsertAuthor inserts an author into the database
func InsertAuthor(ctx context.Context, author Author) (string, error) {
	client := createGraphQLClient()

	mutation := `
		mutation insertAuthor($uuid: uuid!, $name: String!, $age: Int!) {
			insert_Author(objects: {uuid: $uuid, name: $name, age: $age}) {
				returning {
					uuid
					name
				}
			}
		}
	`

	req := graphql.NewRequest(mutation)
	req.Var("uuid", author.Uuid)
	req.Var("name", author.Name)
	req.Var("age", author.Age)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret) // Add admin secret

	var responseData struct {
		InsertAuthor struct {
			Returning []Author `json:"returning"`
		} `json:"insert_Author"`
	}

	if err := client.Run(ctx, req, &responseData); err != nil {
		return "", err
	}

	if len(responseData.InsertAuthor.Returning) == 0 {
		return "", fmt.Errorf("failed to insert author")
	}

	return responseData.InsertAuthor.Returning[0].Uuid, nil
}

// InsertBook inserts a new book into the database using a GraphQL mutation, now with author_id
func InsertBook(ctx context.Context, book Book) error {
	client := createGraphQLClient()

	mutation := `
		mutation insertBook($uuid: uuid!, $title: String!, $genre: String!, $author_id: uuid!) {
			insert_Books(objects: {uuid: $uuid, title: $title, genre: $genre, author_id: $author_id}) {
				returning {
					uuid
					title
					genre
					author_id
				}
			}
		}
	`

	req := graphql.NewRequest(mutation)
	req.Var("uuid", book.Uuid)
	req.Var("title", book.Title)
	req.Var("genre", book.Genre)
	req.Var("author_id", book.AuthorID)
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

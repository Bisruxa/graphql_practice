package schema

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

// Define Author struct
type Author struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define Book struct
type Book struct {
	Uuid   string `json:"uuid"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Author Author `json:"author"`
}

const hasuraEndpoint = "http://localhost:8080/v1/graphql"
const hasuraAdminSecret = "mysecret"

// Create GraphQL client
func createGraphQLClient() *graphql.Client {
	client := graphql.NewClient(hasuraEndpoint)
	client.Log = func(s string) { println(s) }
	return client
}

// Insert an author into Hasura
func InsertAuthor(ctx context.Context, author Author) error {
	client := createGraphQLClient()

	mutation := `
		mutation insertAuthor($uuid: uuid!, $name: String!, $age: Int!) {
			insert_Author_one(object: {uuid: $uuid, name: $name, age: $age}) {
				uuid
			}
		}
	`

	req := graphql.NewRequest(mutation)
	req.Var("uuid", author.Uuid)
	req.Var("name", author.Name)
	req.Var("age", author.Age)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	if err := client.Run(ctx, req, nil); err != nil {
		return fmt.Errorf("error inserting author: %v", err)
	}

	fmt.Println("Author inserted successfully:", author.Uuid)
	return nil
}

// Insert a book and associate it with an existing author
func InsertBook(ctx context.Context, book Book, authorUUID string) error {
	client := createGraphQLClient()

	mutation := `
		mutation insertBook($uuid: uuid!, $title: String!, $genre: String!, $author_uuid: uuid!) {
			insert_Books_one(object: {
				uuid: $uuid,
				title: $title,
				genre: $genre,
				author_id: $author_uuid 
			}) {
				uuid
			}
		}
	`

	req := graphql.NewRequest(mutation)
	req.Var("uuid", book.Uuid)
	req.Var("title", book.Title)
	req.Var("genre", book.Genre)
	req.Var("author_uuid", authorUUID) // Use the correct variable for author UUID

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	if err := client.Run(ctx, req, nil); err != nil {
		return fmt.Errorf("error inserting book: %v", err)
	}

	fmt.Println("Book inserted successfully:", book.Uuid)
	return nil
}

// Fetch all books with their respective authors
func FetchBooks(ctx context.Context) ([]Book, error) {
	client := createGraphQLClient()

	query := `
		query {
			Books {
				uuid
				title
				genre
				author_uuid {
					uuid
					name
					age
				}
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

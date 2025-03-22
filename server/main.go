package main

import (
	"context"
	"fmt"
	"log"
	"github.com/Bisruxa/graphql_practice/schema"
	"github.com/google/uuid"
)

func main() {
	// Insert an author first
	authorID := uuid.New().String() // Generate a unique UUID for the author
	author := schema.Author{
		Uuid:  authorID,
		Name:  "J.K. Rowling",
		Age:   55,
	}

	// Insert the author into the database
	authorUUID, err := schema.InsertAuthor(context.Background(), author)
	if err != nil {
		log.Fatalf("Error inserting author: %v", err)
	}
	fmt.Printf("Inserted author: %s UUID: %s\n", author.Name, authorUUID)

	// Insert books associated with the author
	dummyBooks := []schema.Book{
		{
			Uuid:     uuid.New().String(), // Generate a unique UUID
			Title:    "Harry Potter and the Sorcerer's Stone",
			Genre:    "Fantasy",
			AuthorID: authorUUID, // Associate with the author using their UUID
		},
		{
			Uuid:     uuid.New().String(), // Generate a unique UUID
			Title:    "Harry Potter and the Chamber of Secrets",
			Genre:    "Fantasy",
			AuthorID: authorUUID, // Associate with the author using their UUID
		},
	}

	// Insert each book into the database
	for _, book := range dummyBooks {
		err := schema.InsertBook(context.Background(), book)
		if err != nil {
			log.Fatalf("Error inserting book: %v", err)
		}
		fmt.Printf("Inserted book: %s UUID: %s\n", book.Title, book.Uuid)
	}

	// Fetch all books
	books, err := schema.FetchBooks(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Print all books
	fmt.Println("Books:")
	for _, book := range books {
		fmt.Printf("UUID: %s, Title: %s, Genre: %s, Author ID: %s\n", book.Uuid, book.Title, book.Genre, book.AuthorID)
	}
}

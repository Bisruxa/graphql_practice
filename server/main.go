package main

import (
	"context"
	"fmt"
	"log"
	"github.com/Bisruxa/graphql_practice/schema"
	"github.com/google/uuid"
)

func main() {
	// Insert some dummy books with unique UUIDs
	dummyBooks := []schema.Book{
		{
			Uuid:  uuid.New().String(), // Generate a unique UUID
			Title: "The Hobbit",
			Genre: "Fantasy",
		},
		{
			Uuid:  uuid.New().String(), // Generate a unique UUID
			Title: "1984",
			Genre: "Dystopian",
		},
		{
			Uuid:  uuid.New().String(), // Generate a unique UUID
			Title: "To Kill a Mockingbird",
			Genre: "Classic",
		},
		{
			Uuid:  uuid.New().String(), // Generate a unique UUID
			Title: "Harry Potter",
			Genre: "Fantasy",
		},
	}

	for _, book := range dummyBooks {
		// Insert each book into the database
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
		fmt.Printf("UUID: %s, Title: %s, Genre: %s\n", book.Uuid, book.Title, book.Genre)
	}

	// Fetch a book by UUID (example UUID)
	bookID := dummyBooks[0].Uuid // Use the UUID of the first inserted book
	book, err := schema.FetchBookByID(context.Background(), bookID)
	if err != nil {
		log.Fatal(err)
	}

	if book != nil {
		fmt.Printf("\nFetched Book: UUID: %s, Title: %s, Genre: %s\n", book.Uuid, book.Title, book.Genre)
	} else {
		fmt.Println("\nNo book found with the given UUID.")
	}
}

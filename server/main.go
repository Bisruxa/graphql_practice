package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Bisruxa/graphql_practice/schema"
	"github.com/google/uuid"
)

func main() {
	// Fetch existing books
	books, err := schema.FetchBooks(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Books:")
	for _, book := range books {
		fmt.Printf("ID: %s, Name: %s, Genre: %s\n", book.Uuid, book.Title, book.Genre)
	}

	// Insert dummy books with generated UUIDs
	ctx := context.Background()
	dummyBooks := []schema.Book{
		{Uuid: uuid.New().String(), Title: "The Hobbit", Genre: "Fantasy"},
		{Uuid: uuid.New().String(), Title: "1984", Genre: "Dystopian"},
		{Uuid: uuid.New().String(), Title: "To Kill a Mockingbird", Genre: "Classic"},
	}

	for _, book := range dummyBooks {
		if err := schema.InsertBook(ctx, book); err != nil {
			fmt.Println("Error inserting book:", err)
		} else {
			fmt.Println("Inserted book:", book.Title, "UUID:", book.Uuid)
		}
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Bisruxa/graphql_practice/schema"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	// Fetch and display existing books
	books, err := schema.FetchBooks(ctx)
	if err != nil {
		log.Fatal("Error fetching books:", err)
	}

	fmt.Println("Existing Books:")
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Genre: %s, Author: %s\n",
			book.Uuid, book.Title, book.Genre, book.Author.Name)
	}

	// Create dummy authors and books
	dummyAuthors := []struct {
		Name string
		Age  int
		Book schema.Book
	}{
		{Name: "J.K. Rowling", Age: 57, Book: schema.Book{Uuid: uuid.New().String(), Title: "Harry Potter", Genre: "Fantasy"}},
		{Name: "George Orwell", Age: 46, Book: schema.Book{Uuid: uuid.New().String(), Title: "1984", Genre: "Dystopian"}},
		{Name: "Harper Lee", Age: 89, Book: schema.Book{Uuid: uuid.New().String(), Title: "To Kill a Mockingbird", Genre: "Classic"}},
	}

	// Insert authors and books
	for _, data := range dummyAuthors {
		// Insert Author
		authorUUID := uuid.New().String()
		author := schema.Author{
			Uuid: authorUUID,
			Name: data.Name,
			Age:  data.Age,
		}

		if err := schema.InsertAuthor(ctx, author); err != nil {
			fmt.Println("Error inserting author:", err)
			continue
		}
		fmt.Println("Inserted author:", author.Name, "UUID:", author.Uuid)

		// Insert Book with correct author UUID
		if err := schema.InsertBook(ctx, data.Book, author.Uuid); err != nil {
			fmt.Println("Error inserting book:", err)
		} else {
			fmt.Println("Inserted book:", data.Book.Title, "UUID:", data.Book.Uuid)
		}
	}
}

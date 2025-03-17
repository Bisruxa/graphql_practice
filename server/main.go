package main

import (
    "context"
    "fmt"
    "log"
		"github.com/machinebox/graphql"
)

func main() {
    hasuraEndpoint := "http://localhost:8080/v1/graphql"
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
        Books []struct {
            ID    string `json:"id"`
            Name  string `json:"name"`
            Genre string `json:"genre"`
        } `json:"books"`
    }

    err := client.Run(context.Background(), req, &responseData)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Books:")
    for _, book := range responseData.Books {
        fmt.Printf("ID: %s, Name: %s, Genre: %s\n", book.ID, book.Name, book.Genre)
    }
}

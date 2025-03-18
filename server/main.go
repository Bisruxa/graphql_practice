package main

import (
    "context"
    "fmt"
    "log"
    "github.com/Bisruxa/graphql_practice/schema"
)

func main() {
   books,err := schma.FetchBooks(context.Background())
   if err!= nil{
    log.Fatal(err)
   }
   fmt.Println("Books:")
   for _,book:= range books{
    fmt.Printf("ID: %s,Name: %s,Genre: %s\n",book.ID,book.Name,book.Genre)  
   }
}
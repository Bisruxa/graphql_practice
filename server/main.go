package main

import (
    "context"
    "fmt"
    "log"
)

func main() {
   books,err := FetchBooks(context.Background())
   if err!= nil{
    log.Fatal(err)
   }
   fmt.Println("Books:")
   for _,book:= range books{
    fmt.Printf("ID: %s,Name: %s,Genre: %s\n",book.ID,book.Name,book.Genre)  
   }
}
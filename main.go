package main

import (
	"database/sql"
	"fmt"
	db2 "go-hexagonal/adapters/db"
	"go-hexagonal/application"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Example", 30)
	if err != nil {
        fmt.Println("Error creating product:", err)
        return
    }

    fmt.Println("Product created successfully")
	fmt.Printf("Product ID: %s\n", product.GetId())
	productService.Enable(product)
}
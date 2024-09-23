package main

import (
	"API"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 라우팅 설정
	// GET : /api/item
	http.HandleFunc("/api/item", API.GetItemsHandler)
	// POST: /api/item/create
	http.HandleFunc("/api/item/create", API.CreateItemHandler)

	port := ":8080"
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Cloud not start server: %s\n", err.Error())
	}

}

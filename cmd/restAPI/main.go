package main

import (
	"cloudLabGo/pkg/restAPI"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 라우팅 설정
	// GET : /api/item
	http.HandleFunc("/api/item", restAPI.GetItemsHandler)
	// POST: /api/create
	http.HandleFunc("/api/create", restAPI.CreateItemHandler)

	port := ":8080"
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Cloud not start server: %s\n", err.Error())
	}

}

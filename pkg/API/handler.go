package API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item
var idCounter = 1

// GetItemsHandler GET: /api/item  아이템 목록 가져오기
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// CreateItemHandler POST: /api/item/create - 아이템 생성
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
	}

	newItem.ID = idCounter
	idCounter++
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
	fmt.Fprintf(w, "Item created successfully!")
}

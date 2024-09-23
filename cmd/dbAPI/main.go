package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	InitDB()
	http.HandleFunc("/api/items", GetItemsHandler)
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := DB.Query("SELECT * FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	var items []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	for rows.Next() {
		var item struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)

	}
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/cloudLabGo")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Mysql!")

}

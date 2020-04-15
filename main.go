package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	Id          int
	Title       string
	Description string
}

var (
	items []Item
)

func init() {
	items = append(items, Item{Id: 1, Title: "Title1", Description: "Description1"})
	items = append(items, Item{Id: 2, Title: "Title2", Description: "Description2"})
	items = append(items, Item{Id: 3, Title: "Title3", Description: "Description3"})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/item", getItems)

	log.Fatal(http.ListenAndServe(":0333", r))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getItems")

	json.NewEncoder(w).Encode(items)

}

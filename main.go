package main

import (
	"GoWebServer/models"
	"GoWebServer/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	items []models.Item
)

func init() {
	items = append(items, models.Item{Id: 1, Title: "Title1", Description: "Description1"})
	items = append(items, models.Item{Id: 2, Title: "Title2", Description: "Description2"})
	items = append(items, models.Item{Id: 3, Title: "Title3", Description: "Description3"})
}

func main() {
	r := mux.NewRouter()

	s := r.PathPrefix("/item").Subrouter()

	s.HandleFunc("/", getItems).Methods("GET")
	s.HandleFunc("/{id}", getItemByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":0333", r))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getItems")

	json.NewEncoder(w).Encode(items)

}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, error := strconv.Atoi(vars["id"])
	if error != nil {
		return
	}

	fmt.Printf("getItemByID %d \n", id)

	item, error := utils.FindItemByID(id, items)
	if error != nil {
		return
	}

	json.NewEncoder(w).Encode(item)

}

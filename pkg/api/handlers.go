package api

import (
	"encoding/json"
	"kubernetes-go-app/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

var items = make(map[string]*models.Item)

func GetItems(w http.ResponseWriter, r *http.Request) {
	itemsList := make([]*models.Item, 0, len(items))
	for _, item := range items {
		itemsList = append(itemsList, item)
	}
	json.NewEncoder(w).Encode(itemsList)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if item, ok := items[params["id"]]; ok {
		json.NewEncoder(w).Encode(item)
		return
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	items[item.ID] = &item
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := items[params["id"]]; ok {
		var item models.Item
		json.NewDecoder(r.Body).Decode(&item)
		item.ID = params["id"]
		items[params["id"]] = &item
		json.NewEncoder(w).Encode(item)
		return
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := items[params["id"]]; ok {
		delete(items, params["id"])
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

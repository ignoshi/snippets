package snippets

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emostafa/jopher"
	"github.com/gorilla/mux"
)

func ListSnippets(w http.ResponseWriter, r *http.Request) {
	m := SnippetsManager{}
	items, err := m.Find(nil)
	if err != nil {
		jopher.NotFound(w, err)
		return
	}
	jopher.Success(w, items)
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	item := Snippet{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		jopher.Error(w, 400, err)
		return
	}
	log.Println("snippet", item.Title, item.Body)
	if ok, errs := item.IsValid(); !ok {
		log.Println("item is valid?", ok, errs)
		jopher.Error(w, 400, errors.New("Invalid Body"))
		return
	}
	item.Save()
	jopher.Success(w, item)
}

func GetSnippet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := bson.M{"_id": vars["id"]}
	m := SnippetsManager{}
	item, err := m.FindOne(q)
	if err != nil {
		jopher.NotFound(w, err)
		return
	}
	jopher.Success(w, item)
}

package controller

import (
	"ProjectsAPI/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	entry, err := json.Marshal(&model.AuthorsCellection)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(entry)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, author := range model.AuthorsCellection {
		if author.ID == id {
			entry, err := json.Marshal(author)
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(entry)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author model.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	model.AuthorsCellection = append(model.AuthorsCellection, author)
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, author := range model.AuthorsCellection {
		if author.ID == id {
			model.AuthorsCellection = append(model.AuthorsCellection[:index], model.AuthorsCellection[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, authorToUpdate := range model.AuthorsCellection {
		if authorToUpdate.ID == id {
			err := json.NewDecoder(r.Body).Decode(&authorToUpdate)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			model.AuthorsCellection = append(append(model.AuthorsCellection[:index], authorToUpdate), model.AuthorsCellection[index+1:]...)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(authorToUpdate)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	store := store.GetStore()
	toDos := store.GetToDos()
	tmpl, err := template.ParseFiles("view.html")

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, toDos)
}

func HandleCreateToDo(w http.ResponseWriter, r *http.Request) {
	store := store.GetStore()
	toDo := r.FormValue("toDo")

	if toDo == "" {
		http.Error(w, "toDo cannot be empty", http.StatusBadRequest)
		return
	}

	store.Add(toDo)
	http.Redirect(w, r, "/todos", http.StatusFound)

}

func HandleUpdateStatus(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse id"))
		return
	}

	var payload struct {
		Status bool `json:"completed"`
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&payload)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse request body"))
		return
	}

	store := store.GetStore()
	updateErr := store.UpdateToDo(id, payload.Status)

	if updateErr != nil {
		//TODO: Think I need to change this or Oli might put me in the bin :)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("To Do with ID %d not found", id)))
		return
	}

	w.WriteHeader(http.StatusOK)
}

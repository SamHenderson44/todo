package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	store := store.GetStore()
	toDos := store.GetToDos()
	tmpl, err := template.ParseFiles("/Users/sam.henderson/bench/go/todo/internal/routes/view.html")

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
	fmt.Printf("rec %v ", toDo)
	http.Redirect(w, r, "/todos", http.StatusFound)

}

package main

import (
	store "github.com/SamHenderson44/todo/internal/storePackage"
)

func HandleSaveNewToDo(newToDoItem string) {
	store := store.GetStore()
	store.Add(newToDoItem)
}

func HandleGetToDos() []store.ToDo {
	store := store.GetStore()
	return store.GetToDos()
}

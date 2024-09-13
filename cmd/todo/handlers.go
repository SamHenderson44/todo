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

func HandleGetToDo(ID int) (store.ToDo, error) {
	store := store.GetStore()
	todo, err := store.GetToDo(ID)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func HandleUpdateToDo(ID int, completed bool) error {
	store := store.GetStore()
	return store.UpdateToDo(ID, completed)
}

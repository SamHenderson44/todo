package store

import (
	"fmt"
	"slices"
	"sync"
)

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

type Store struct {
	toDos []ToDo
}

var once sync.Once
var instance *Store

const ToDoNotFoundError = "no to do found with ID"

func GetStore() *Store {
	once.Do(func() {
		instance = &Store{
			toDos: []ToDo{},
		}
	})
	return instance
}

func (s *Store) Add(title string) {
	newToDo := ToDo{
		ID:        len(s.toDos) + 1,
		Title:     title,
		Completed: false,
	}
	s.toDos = append(s.toDos, newToDo)
}

func (s *Store) GetToDos() []ToDo {
	return s.toDos
}

func (s *Store) GetToDo(ID int) (ToDo, error) {
	index, err := getIndex(ID, s.toDos)

	if err != nil {
		return ToDo{}, err
	}

	return s.toDos[index], nil

}

func (s *Store) DeleteToDo(ID int) error {
	index, err := getIndex(ID, s.toDos)

	if err != nil {
		return err
	}

	s.toDos = append(s.toDos[:index], s.toDos[index+1:]...)
	return nil
}

func (s *Store) UpdateToDo(ID int, completed bool) error {
	index, err := getIndex(ID, s.toDos)

	if err != nil {
		return err
	}

	s.toDos[index].Completed = completed
	return nil
}

func getIndex(ID int, toDos []ToDo) (int, error) {
	index := slices.IndexFunc(toDos, func(todo ToDo) bool {
		return todo.ID == ID
	})

	if index < 0 {
		return -1, fmt.Errorf(ToDoNotFoundError+" %d", ID)
	}

	return index, nil

}

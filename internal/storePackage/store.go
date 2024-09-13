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
	ToDos []ToDo
}

var once sync.Once
var instance *Store

const ToDoNotFoundError = "no to do found with ID"

func GetStore() *Store {
	once.Do(func() {
		instance = &Store{
			ToDos: []ToDo{},
		}
	})
	return instance
}

func (s *Store) Add(title string) {
	newToDo := ToDo{
		ID:        len(s.ToDos) + 1,
		Title:     title,
		Completed: false,
	}
	s.ToDos = append(s.ToDos, newToDo)
}

func (s *Store) GetToDos() []ToDo {
	return s.ToDos
}

func (s *Store) GetToDo(ID int) (ToDo, error) {
	index, err := getIndex(ID, s.ToDos)

	if err != nil {
		return ToDo{}, err
	}

	return s.ToDos[index], nil

}

func (s *Store) DeleteToDo(ID int) error {
	index, err := getIndex(ID, s.ToDos)

	if err != nil {
		return err
	}

	s.ToDos = append(s.ToDos[:index], s.ToDos[index+1:]...)
	return nil
}

func (s *Store) UpdateToDo(ID int, completed bool) error {
	index, err := getIndex(ID, s.ToDos)

	if err != nil {
		return err
	}

	s.ToDos[index].Completed = completed
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

func (s *Store) ResetStore() {
	s.ToDos = nil
}

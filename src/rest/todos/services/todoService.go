package services

import (
	"github.com/kamva/mgm/v3"
	"rest-api/src/models"
)

type TodoService struct{}

func (todoService TodoService) Create(todo *models.Todo) error {
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func findTodo(id string) (*models.Todo, error) {
	foundTodo := &models.Todo{}
	err := mgm.Coll(foundTodo).FindByID(id, foundTodo)
	if err != nil {
		return nil, err
	}

	return foundTodo, nil
}

func (todoService TodoService) ToggleCompleted(id string) (*models.Todo, error) {
	todo, err := findTodo(id)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	mongoErr := mgm.Coll(todo).Update(todo)
	if mongoErr != nil {
		return nil, mongoErr
	}

	return todo, nil
}

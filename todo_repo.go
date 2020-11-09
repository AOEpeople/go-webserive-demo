package main

import (
	"errors"
)

type (
	// InMemoryRepo is an example implementation
	// it is not concurrent safe.
	InMemoryRepo struct {
		todos map[int]Todo
	}
)

var (
	// ErrTodoNotFound is returned on access with a non-existing ID
	ErrTodoNotFound = errors.New("todo not found")
)

// Get the entry with the given id
func (r *InMemoryRepo) Get(id int) (Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return Todo{}, ErrTodoNotFound
	}

	return todo, nil
}

// All entries as unsorted list
func (r *InMemoryRepo) All() ([]Todo, error) {
	todos := make([]Todo, len(r.todos))
	i := 0
	for _, todo := range r.todos {
		todos[i] = todo
		i++
	}

	return todos, nil
}

// Save the given instance under the ID inside
func (r *InMemoryRepo) Save(todo Todo) error {
	if r.todos == nil {
		r.todos = make(map[int]Todo)
	}

	r.todos[todo.ID] = todo

	return nil
}

// Delete the entry with the given ID
// Delete is no-op if the ID doesn't exist
func (r *InMemoryRepo) Delete(id int) error {
	delete(r.todos, id)
	return nil
}

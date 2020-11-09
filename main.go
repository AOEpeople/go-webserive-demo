package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type (
	// Todo domain model
	Todo struct {
		ID      int
		Message string
	}

	// TodoRepository defines persistence actions
	TodoRepository interface {
		All() ([]Todo, error)
		Get(id int) (Todo, error)
		Save(todo Todo) error
		Delete(id int) error
	}
)

func main() {
	var todoRepository TodoRepository = new(InMemoryRepo)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello"))
	})

	_ = todoRepository.Save(Todo{
		ID:      1,
		Message: "Todo 1",
	})

	http.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		todos, err := todoRepository.All()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		todoJSON, err := json.Marshal(todos)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(todoJSON)
	})

	http.HandleFunc("/todo/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			idString := strings.TrimPrefix(request.URL.Path, "/todo/")
			id, err := strconv.Atoi(idString)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			todo, err := todoRepository.Get(id)
			switch {
			case errors.Is(err, ErrTodoNotFound):
				http.Error(writer, err.Error(), http.StatusNotFound)
				return
			case err != nil:
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			data, err := json.Marshal(todo)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = writer.Write(data)
		case http.MethodDelete:
			idString := strings.TrimPrefix(request.URL.Path, "/todo/")
			id, err := strconv.Atoi(idString)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			err = todoRepository.Delete(id)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = writer.Write([]byte("deleted"))
		case http.MethodPost,
			http.MethodPatch:
			data, err := ioutil.ReadAll(request.Body)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			todo := new(Todo)
			err = json.Unmarshal(data, todo)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			err = todoRepository.Save(*todo)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = writer.Write([]byte("saved"))
		}
	})

	fmt.Println("serving on :1111 ...")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		panic(err)
	}
}

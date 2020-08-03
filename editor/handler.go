package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
	"task-editor/scheduling"
)

type taskLister interface {
	ListTasks() ([]scheduling.Task, error)
}

func listTasks(l taskLister) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tasks, err := l.ListTasks()
		if err != nil {
			http.Error(writer, "error listing tasks: " + err.Error(), http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(writer).Encode(tasks)
	}
}

type taskAdder interface {
	AddTask(task scheduling.Task) error
}

func createTask(a taskAdder) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var task scheduling.Task
		err := json.NewDecoder(request.Body).Decode(&task)
		if err != nil {
			http.Error(writer, "error decoding body as json", http.StatusBadRequest)
			return
		}
		err = a.AddTask(task)
		if err != nil {
			http.Error(writer, "error writing new task: " + err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

type taskDeleter interface {
	DeleteTask(id string) error
}

func deleteTask(d taskDeleter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")
		err := d.DeleteTask(id)
		if err != nil && strings.Contains(err.Error(), "not found") {
			writer.WriteHeader(http.StatusOK)
			return
		} else if err != nil {
			http.Error(writer, "error deleting item: " + err.Error(), http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)
	}
}

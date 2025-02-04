package router

import (
	"barebone-go-crud/src/handler"
	"net/http"
)

func NewRouter(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandler.CreateUser(w, r)
			return
		}
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandler.GetUserById(w, r)
			return
		}
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			userHandler.UpdateUser(w, r)
			return
		}
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			userHandler.DeleteUser(w, r)
			return
		}
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	})

	return mux
}

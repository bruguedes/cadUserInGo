package api

import (
	"cadUser/model"
	"cadUser/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Variável global para validação
var validate = validator.New()

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(JsonMiddleware)

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", handleCreateUser)
		// r.Get("/", handleAllUsers)
		// r.Get("/{id}", handleGetUser)
		// r.Put("/{id}", handleUpdateUser)
		// r.Delete("/{id}", handleDeleteUser)

	})

	return r

}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSON(w, model.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		return
	}

	// Valida os campos da struct User
	if err := validate.Struct(user); err != nil {
		msg := "Please provide FirstName LastName and bio for the user"

		utils.SendJSON(w, model.Response{Message: msg}, http.StatusBadRequest)
		return
	}

	result, err := user.Insert()

	if err != nil {
		utils.SendJSON(w, model.Response{Error: "failed to insert user"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, model.Response{Data: result}, http.StatusCreated)

}

// func handleAllUsers(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func handleGetUser(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
// 	return
// }
// func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
// 	return
// }

package api

import (
	"cadUser/model"
	"cadUser/utils"
	"encoding/json"
	"errors"
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
		r.Get("/", handleGetAllUsers)
		r.Get("/{id}", handleUserGetById)
		r.Delete("/{id}", handleDeleteUser)
		r.Put("/{id}", handleUpdateUser)

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

		utils.SendJSON(w, model.Response{Error: "invalid input: " + err.Error()}, http.StatusBadRequest)
		return
	}

	result, err := user.Insert()

	if err != nil {

		utils.SendJSON(w, model.Response{Error: "failed to insert user: " + err.Error()}, http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, model.Response{Data: result}, http.StatusCreated)

}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	//busca uma chave especifica na query string
	// e retorna o valor dela
	filtersParam := r.URL.Query().Get("filters")

	var userParams map[string]string

	if filtersParam != "" {

		if err := json.Unmarshal([]byte(filtersParam), &userParams); err != nil {
			utils.SendJSON(w, model.Response{Error: "invalid filters"}, http.StatusBadRequest)
			return
		}

	}

	users, err := user.FindAll(userParams)
	if err != nil {
		utils.SendJSON(w, model.Response{Error: "failed to find users"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, model.Response{Data: users}, http.StatusOK)
}

func handleUserGetById(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	id := chi.URLParam(r, "id")

	result, err := user.GetByID(id)

	if err != nil {
		if errors.Is(err, model.ErrInvalidUserID) {
			utils.SendJSON(w, model.Response{Error: "invalid user ID"}, http.StatusBadRequest)
			return
		}
		if errors.Is(err, model.ErrNotFound) {
			utils.SendJSON(w, model.Response{Message: "The user with the specified ID does not exist"}, http.StatusNotFound)
			return
		}

		utils.SendJSON(w, model.Response{Message: "The user information could not be retrieved"}, http.StatusInternalServerError)
		return

	}

	utils.SendJSON(w, model.Response{Data: result}, http.StatusOK)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user := model.User{}

	deletedUser, err := user.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, model.ErrInvalidUserID):
			utils.SendJSON(w, model.Response{Error: "invalid user ID"}, http.StatusBadRequest)
			return
		case errors.Is(err, model.ErrNotFound):
			utils.SendJSON(w, model.Response{Message: "The user with the specified ID does not exist"}, http.StatusNotFound)
			return

		default:
			utils.SendJSON(w, model.Response{Message: "The user could not be deleted"}, http.StatusInternalServerError)
			return
		}

	}

	utils.SendJSON(w, model.Response{Data: deletedUser}, http.StatusOK)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSON(w, model.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		return
	}

	// Valida os campos da struct User
	if err := validate.Struct(user); err != nil {

		utils.SendJSON(w, model.Response{Error: "invalid input: " + err.Error()}, http.StatusBadRequest)
		return
	}

	updatedUser, err := user.Update(id)

	if err != nil {
		switch {
		case errors.Is(err, model.ErrInvalidUserID):
			utils.SendJSON(w, model.Response{Error: "invalid user ID"}, http.StatusBadRequest)
			return
		case errors.Is(err, model.ErrNotFound):
			utils.SendJSON(w, model.Response{Message: "The user with the specified ID does not exist"}, http.StatusNotFound)
			return

		default:
			utils.SendJSON(w, model.Response{Message: "The user could not be deleted"}, http.StatusInternalServerError)
			return
		}

	}

	utils.SendJSON(w, model.Response{Data: updatedUser}, http.StatusOK)
}

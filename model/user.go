package model

import (
	"errors"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	// ID        uuid.UUID `json:"id" validate:"required"`                       // UUID, obrigatório
	FirstName string `json:"first_name" validate:"required,min=2,max=20"`  // String, obrigatório, tamanho entre 2 e 20
	LastName  string `json:"last_name" validate:"required,min=2,max=20"`   // String, obrigatório, tamanho entre 2 e 20
	Biography string `json:"biography" validate:"required,min=20,max=450"` // String, obrigatório, tamanho entre 20 e 450

}

type application struct {
	mu   sync.Mutex // Mutex para controle de concorrência
	data map[uuid.UUID]User
}

var App = application{
	mu:   sync.Mutex{},             // Inicializa o mutex
	data: make(map[uuid.UUID]User), // Inicializa o mapa vazio
}

var (
	ErrInvalidUserID = errors.New("invalid user ID")
	ErrNotFound      = errors.New("not found")
	ErrInsertUser    = errors.New("failed to insert user")
)

func (u *User) Insert() (map[string]string, error) {

	App.mu.Lock()
	defer App.mu.Unlock()

	// Gera um novo UUID
	id := getUUID(&App)

	App.data[id] = *u

	if value, exists := App.data[id]; !exists || value != *u {
		return nil, ErrInsertUser
	}

	return map[string]string{
		"id":         id.String(),
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"biography":  u.Biography,
	}, nil
}

func FindAll(params map[string]string) ([]map[string]string, error) {
	App.mu.Lock()
	defer App.mu.Unlock()

	users := make([]map[string]string, 0)

	if len(App.data) == 0 {
		return users, nil
	}

	if len(params) == 0 {
		for key, value := range App.data {

			users = append(users, buildUser(key, value))

		}

		return users, nil
	}

	if params["first_name"] != "" && params["last_name"] != "" {

		for key, value := range App.data {
			if strings.EqualFold(value.FirstName, params["first_name"]) && strings.EqualFold(value.LastName, params["last_name"]) {
				users = append(users, buildUser(key, value))
				continue
			}
		}
		return users, nil
	}

	if params["first_name"] != "" || params["last_name"] != "" {

		for key, value := range App.data {

			if strings.EqualFold(value.FirstName, params["first_name"]) {

				users = append(users, buildUser(key, value))
				continue
			}

			if strings.EqualFold(value.LastName, params["last_name"]) {

				users = append(users, buildUser(key, value))
				continue
			}
		}

	}

	return users, nil

}

func GetByID(id string) (map[string]string, error) {
	App.mu.Lock()
	defer App.mu.Unlock()

	parsedId, err := uuid.Parse(id)

	if err != nil {

		return nil, ErrInvalidUserID
	}

	result, exists := App.data[parsedId]

	if !exists {
		return nil, ErrNotFound
	}

	return map[string]string{
		"id":         id,
		"first_name": result.FirstName,
		"last_name":  result.LastName,
		"biography":  result.Biography,
	}, nil

}

func getUUID(App *application) uuid.UUID {
	// Verifica se o ID já existe no mapa
	// Se o ID já existe, gera um novo ID
	// Se o ID não existe, retorna o ID

	for {
		id := uuid.New()
		if _, exists := App.data[id]; !exists {
			return id // Retorna o ID se ele não existir no mapa
		}
		// Caso contrário, o loop continua gerando novos UUIDs
	}

}

func buildUser(id uuid.UUID, user User) map[string]string {
	return map[string]string{
		"id":         id.String(),
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"biography":  user.Biography,
	}
}

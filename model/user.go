package model

import (
	"errors"
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

func (u *User) Insert() (map[string]string, error) {

	App.mu.Lock()
	defer App.mu.Unlock()

	// Gera um novo UUID
	id := getUUID(&App)

	App.data[id] = *u

	if value, exists := App.data[id]; !exists || value != *u {
		return nil, errors.New("failed to insert user")
	}

	return map[string]string{
		"id":         id.String(),
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"biography":  u.Biography,
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

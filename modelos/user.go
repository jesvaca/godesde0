package modelos

import (
	"time"
)

// Definición de la estructura
type Usuario struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// Definición de la clase
func (usuario *Usuario) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}

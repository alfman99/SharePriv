package models

import "time"

type Clave struct {
	Clave         string    `gorm:"primary_key"`
	FechaCreacion time.Time `gorm:"default:now()"`
	GruposUsando  []Grupo
}

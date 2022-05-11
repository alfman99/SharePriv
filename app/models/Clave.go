package models

import "time"

type Clave struct {
	Clave         string    `gorm:"primary_key;type:char(32);"`
	FechaCreacion time.Time `gorm:"default:now()"`
	GruposUsando  []Grupo
}

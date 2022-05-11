package models

import (
	"time"
)

type Usuario struct {
	Username      string    `gorm:"primary_key; type:varchar(50); unique_index; not null"`
	Password      string    `gorm:"type:varchar(50); not null"`
	FechaCreacion time.Time `gorm:"default:now()"`
	Grupos        []Grupo   `gorm:"many2many:usuarios_grupos"`

	PropietarioGrupo   Grupo   `gorm:"foreignkey:PropietarioGrupoUuid"`
	PropietarioArchivo Archivo `gorm:"foreignkey:PropietarioArchivoUuid"`

	InvitacionRegistroCodigo string `gorm:"not null; type:varchar(50);"`
}

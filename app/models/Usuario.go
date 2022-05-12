package models

import (
	"time"
)

type Usuario struct {
	Username      string    `gorm:"primary_key; type:varchar(50); unique_index; not null"`
	Password      string    `json:"-" gorm:"type:varchar(50); not null"`
	FechaCreacion time.Time `gorm:"default:now()"`
	Grupos        []Grupo   `gorm:"many2many:usuarios_grupos"`

	InvitacionesRegistroCreadas []InvitacionRegistro `json:"-" gorm:"foreignkey:Propietario"`
	InvitacionesGrupoCreadas    []InvitacionGrupo    `json:"-" gorm:"foreignkey:Propietario"`

	PropietarioGrupos  []Grupo   `gorm:"foreignKey:PropietarioUsername"`
	PropietarioArchivo []Archivo `gorm:"foreignkey:PropietarioArchivo"`

	InvitacionRegistroCodigo string `json:"-" gorm:"not null; type:varchar(50);"`
}

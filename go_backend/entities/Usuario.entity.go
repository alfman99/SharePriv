package entities

import (
	"time"
)

type Usuario struct {
	Username      string    `gorm:"primary_key; type:varchar(50); unique_index; not null"`
	Password      string    `json:"-" gorm:"type:varchar(50); not null"`
	FechaCreacion time.Time `gorm:"default:now();not null;"`
	Grupos        []Grupo   `gorm:"many2many:usuarios_grupos"` // Asociativa

	InvitacionesRegistroCreadas []InvitacionRegistro `json:"-" gorm:"foreignkey:Propietario"`
	InvitacionesGrupoCreadas    []InvitacionGrupo    `json:"-" gorm:"foreignkey:Propietario"`

	PropietarioGrupos []Grupo `gorm:"foreignKey:Propietario"`

	PropietarioArchivoPublico []ArchivoPublico `gorm:"foreignkey:Propietario"`
	PropietarioArchivoGrupo   []ArchivoGrupo   `gorm:"foreignkey:Propietario"`

	InvitacionRegistroCodigo string `json:"-" gorm:"type:varchar(50); not null;"`
}

func (u *Usuario) TableName() string {
	return "usuario"
}

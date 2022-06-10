package entities

import (
	"time"
)

type Invitacion struct {
	Codigo         string    `gorm:"primary_key; type:varchar(50);"`
	FechaCreacion  time.Time `gorm:"default:now();not null;"`
	FechaCaducidad time.Time
	MaximoUsos     uint `gorm:"not null;"`
	Usos           uint `gorm:"default:0;not null;"`

	Propietario string `gorm:"not null"`
}

type InvitacionGrupo struct {
	Invitacion

	GrupoId string `gorm:"not null;"`
}

type InvitacionRegistro struct {
	Invitacion

	UsuariosRegistrados []Usuario `gorm:"foreignkey:InvitacionRegistroCodigo"`
}

func (u *InvitacionGrupo) TableName() string {
	return "invitacion_grupo"
}

func (u *InvitacionRegistro) TableName() string {
	return "invitacion_registro"
}

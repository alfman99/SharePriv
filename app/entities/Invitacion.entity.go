package entities

import (
	"time"
)

type Invitacion struct {
	Codigo         string    `gorm:"primary_key; type:varchar(50);"`
	FechaCreacion  time.Time `gorm:"default:now()"`
	FechaCaducidad time.Time
	MaximoUsos     uint
	Usos           uint `gorm:"default:0"`

	Propietario string `gorm:"foreignkey:PropietarioID"`
}

type InvitacionGrupo struct {
	Invitacion

	GrupoId string `gorm:"not null;"`
}

type InvitacionRegistro struct {
	Invitacion

	// UsuariosInvitadosRegistro []Usuario `json:"-"`
}

func (u *InvitacionGrupo) TableName() string {
	return "invitacion_grupo"
}

func (u *InvitacionRegistro) TableName() string {
	return "invitacion_registro"
}

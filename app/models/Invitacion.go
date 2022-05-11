package models

import "time"

type Invitacion struct {
	Codigo         string    `gorm:"primary_key; type:varchar(50);"`
	FechaCreacion  time.Time `gorm:"default:now()"`
	FechaCaducidad time.Time
	MaximoUsos     uint
	Usos           uint `gorm:"default:0"`
}

type InvitacionGrupo struct {
	Invitacion

	GrupoUuid string `gorm:"not null; type:uuid"`
}

type InvitacionRegistro struct {
	Invitacion

	UsuariosInvitadosRegistro []Usuario
}

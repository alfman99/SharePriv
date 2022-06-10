package entities

import (
	"time"
)

type Grupo struct {
	Id                string    `gorm:"primary_key;default:md5(concat(random()::text, clock_timestamp()::text));not null;"`
	Nombre            string    `gorm:"type:varchar(50); unique_index"`
	FechaCreacion     time.Time `gorm:"default:now();not null"`
	Usuarios          []Usuario `gorm:"many2many:usuarios_grupos"` // Asociativa
	Archivos          []ArchivoGrupo
	InvitacionesGrupo []InvitacionGrupo
	Propietario       string `gorm:"not null; type:varchar(50);"`
}

type UsuariosGrupos struct {
	Grupo_Id         string    `gorm:"primary_key;not null"`
	Usuario_username string    `gorm:"primary_key; type:varchar(50);not null;"`
	FechaRegistro    time.Time `gorm:"default:now();not null;"`
}

func (u *Grupo) TableName() string {
	return "grupo"
}

func (ug *UsuariosGrupos) TableName() string {
	return "usuarios_grupos"
}

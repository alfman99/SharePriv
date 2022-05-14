package entities

import (
	"time"
)

type Grupo struct {
	Uuid                string    `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Nombre              string    `gorm:"type:varchar(50); unique_index"`
	FechaCreacion       time.Time `gorm:"default:now()"`
	Usuarios            []Usuario `gorm:"many2many:usuarios_grupos"` // Asociativa
	Archivos            []ArchivoGrupo
	InvitacionesGrupo   []InvitacionGrupo
	PropietarioUsername string `gorm:"not null; type:varchar(50);"`
}

type UsuariosGrupos struct {
	Grupo_uuid       string    `gorm:"primary_key; type:uuid; not null"`
	Usuario_username string    `gorm:"primary_key; type:varchar(50); not null"`
	FechaRegistro    time.Time `gorm:"default:now(); not null"`
}

func (u *Grupo) TableName() string {
	return "grupo"
}

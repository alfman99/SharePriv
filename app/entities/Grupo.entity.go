package entities

import (
	"time"
)

type Grupo struct {
	Id                  string    `gorm:"primary_key;default:md5(concat(random()::text, clock_timestamp()::text))"`
	Nombre              string    `gorm:"type:varchar(50); unique_index"`
	FechaCreacion       time.Time `gorm:"default:now()"`
	Usuarios            []Usuario `gorm:"many2many:usuarios_grupos"` // Asociativa
	Archivos            []ArchivoGrupo
	InvitacionesGrupo   []InvitacionGrupo
	PropietarioUsername string `gorm:"not null; type:varchar(50);"`
}

type UsuariosGrupos struct {
	Grupo_Id         string    `gorm:"primary_key; not null"`
	Usuario_username string    `gorm:"primary_key; type:varchar(50); not null"`
	FechaRegistro    time.Time `gorm:"default:now(); not null"`
}

func (u *Grupo) TableName() string {
	return "grupo"
}

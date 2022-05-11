package models

import "time"

type Grupo struct {
	Uuid                 string    `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Nombre               string    `gorm:"type:varchar(50); unique_index"`
	FechaCreacion        time.Time `gorm:"default:now()"`
	Usuarios             []Usuario `gorm:"many2many:usuarios_grupos"` // Asociativa
	Archivos             []ArchivoGrupo
	InvitacionesGrupo    []InvitacionGrupo
	PropietarioGrupoUuid string `gorm:"not null; type:uuid"`
	ClaveClave           string `gorm:"not null;"`
}

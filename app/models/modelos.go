package models

import "time"

type Usuario struct {
	Uuid          string    `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Username      string    `gorm:"type:varchar(50); unique_index"`
	Password      string    `gorm:"type:varchar(50)"`
	FechaCreacion time.Time `gorm:"default:now()"`
	Grupos        []Grupo   `gorm:"many2many:usuarios_grupos"`

	PropietarioGrupo   Grupo   `gorm:"foreignkey:PropietarioGrupoUuid"`
	PropietarioArchivo Archivo `gorm:"foreignkey:PropietarioArchivoUuid"`

	InvitacionRegistroCodigo string `gorm:"not null; type:varchar(50);"`
}

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

type Archivo struct {
	Uuid             string    `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Data             []byte    `gorm:"type:bytea"`
	Visualizaciones  uint      `gorm:"default:0"`
	FechaPublicacion time.Time `gorm:"default:now()"`

	PropietarioArchivoUuid string `gorm:"not null; type:uuid"`
}

type ArchivoGrupo struct {
	Archivo
	GrupoUuid string
}

type ArchivoPublico struct {
	Archivo
	Password string `gorm:"type:varchar(50)"`
}

type Clave struct {
	Clave         string    `gorm:"primary_key"`
	FechaCreacion time.Time `gorm:"default:now()"`
	GruposUsando  []Grupo
}

type Invitacion struct {
	Codigo         string    `gorm:"primary_key; type:varchar(50);"`
	FechaCreacion  time.Time `gorm:"default:now()"`
	FechaCaducidad time.Time
	MaximoUsos     uint
	Usos           uint
}

type InvitacionGrupo struct {
	Invitacion

	GrupoUuid string `gorm:"not null; type:uuid"`
}

type InvitacionRegistro struct {
	Invitacion

	UsuariosInvitadosRegistro []Usuario
}

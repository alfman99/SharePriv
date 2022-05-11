package models

import "time"

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

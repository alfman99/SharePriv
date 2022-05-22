package entities

import "time"

type Archivo struct {
	Id               string    `gorm:"primary_key;default:md5(concat(random()::text, clock_timestamp()::text))"`
	Data             []byte    `gorm:"type:bytea"`
	Visualizaciones  uint      `gorm:"default:0"`
	FechaPublicacion time.Time `gorm:"default:now()"`
	Mime             string    `gorm:"not null"`

	PropietarioArchivo string `gorm:"not null;type:varchar(50);"`
}

type ArchivoGrupo struct {
	Archivo

	GrupoId string
}

type ArchivoPublico struct {
	Archivo
}

func (u *ArchivoGrupo) TableName() string {
	return "archivo_grupo"
}

func (u *ArchivoPublico) TableName() string {
	return "archivo_publico"
}
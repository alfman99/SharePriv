package entities

import "time"

type Archivo struct {
	Uuid             string    `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Data             []byte    `gorm:"type:bytea"`
	Visualizaciones  uint      `gorm:"default:0"`
	FechaPublicacion time.Time `gorm:"default:now()"`
	Mime             string    `gorm:"not null"`

	PropietarioArchivo string `gorm:"not null;type:varchar(50);"`
}

type ArchivoGrupo struct {
	Archivo

	GrupoUuid string
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

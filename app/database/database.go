package database

import (
	"log"
	"sharepriv/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var InstanciaDB *gorm.DB

func ConnectDB() {
	connectionString := "postgres://postgres:postgrespw@localhost:49153"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Conexión a la base de datos correcta")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Hacer migraciones")

	err = db.SetupJoinTable(&entities.Usuario{}, "Grupos", &entities.UsuariosGrupos{})

	if err != nil {
		panic(err)
	}

	// Hacer migraciones
	db.AutoMigrate(
		&entities.Grupo{},
		&entities.InvitacionGrupo{},
		&entities.InvitacionRegistro{},
		&entities.ArchivoGrupo{},
		&entities.ArchivoPublico{},
		&entities.Usuario{},
	)

	InstanciaDB = db
}

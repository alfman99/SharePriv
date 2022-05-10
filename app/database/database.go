package database

import (
	"log"
	"sharepriv/models"

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

	// Hacer migraciones
	db.AutoMigrate(
		&models.Clave{},
		&models.Grupo{},
		&models.InvitacionGrupo{},
		&models.InvitacionRegistro{},
		&models.ArchivoGrupo{},
		&models.ArchivoPublico{},
		&models.Usuario{},
	)

	InstanciaDB = db
}

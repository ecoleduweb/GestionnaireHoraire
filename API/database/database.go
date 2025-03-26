package database

import (
	"llio-api/useful"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	mysqld "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Charger les variables d'environnement
	useful.LoadEnv()

	log.Println("Exécution des migrations...")
	// Run migrations first
	err := useful.RunMigrationCommand("up")
	if err != nil {
		log.Printf("Attention lors des migrations: %v", err)
		// Continue anyway since it might be an ErrNoChange or other non-fatal error
	}

	// Then connect to DB
	db, err := gorm.Open(mysqld.Open(useful.GetGormDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données : ", err)
	} else {
		log.Println("Connexion à la base de données effectuée avec succès")
	}

	DB = db
}

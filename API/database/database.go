package database

import (
	"fmt"
	"llio-api/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {

	// Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	// Lire les variables d'environnement
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Construire la chaîne de connexion DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données : ", err)
	} else {
		log.Println("Connexion effectué")
	}

	// Migrer le modèle
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Erreur lors de la migration des modèles : ", err)
	} else {
		log.Println("Migrations effectuées")
	}

	DB = db
}

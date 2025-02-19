package database

import (
	"fmt"
	"llio-api/models/DAOs"
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
	env := os.Getenv("ENV")
	println(env)
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")

	var databaseName string

	// Définir le host et le nom de la base de données en fonction de l'environnement
	switch env {
	case "TEST":
		databaseName = os.Getenv("DB_NAME_TEST")
	case "DEV", "PROD":
		databaseName = os.Getenv("DB_NAME_RUN")
	default:
		log.Fatalf("Environnement non pris en charge : %s", env)
	}

	// Construire la chaîne de connexion DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, databaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données : ", err)
	} else {
		log.Println("Connexion à la base de données effectuée avec succès")
	}

	// Migrer le modèle
	err = db.AutoMigrate(&DAOs.Task{})
	if err != nil {
		log.Fatal("Erreur lors de la migration des modèles : ", err)
	} else {
		log.Println("Migrations effectuées avec succès")
	}

	DB = db
}

package useful

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// GetDSN returns migration-compatible DSN
func GetDSN() string {
	// Charger les variables d'environnement
	LoadEnv()

	env := os.Getenv("ENV")
	var databaseName string

	switch env {
	case "TEST":
		databaseName = os.Getenv("DB_NAME_TEST")
	case "DEV", "PROD":
		databaseName = os.Getenv("DB_NAME_RUN")
	default:
		log.Fatalf("Environnement non pris en charge : %s", env)
	}

	// DSN for migrations
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		databaseName,
	)

	return dsn
}

// GetGormDSN returns GORM-compatible DSN
func GetGormDSN() string {
	// Charger les variables d'environnement
	LoadEnv()

	env := os.Getenv("ENV")
	var databaseName string

	switch env {
	case "TEST":
		databaseName = os.Getenv("DB_NAME_TEST")
	case "DEV", "PROD":
		databaseName = os.Getenv("DB_NAME_RUN")
	default:
		log.Fatalf("Environnement non pris en charge : %s", env)
	}

	// DSN for GORM
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		databaseName,
	)

	return dsn
}

func GetMigrationPath() string {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erreur lors de l'obtention du répertoire de travail : %v\n", err)
		os.Exit(1)
	}
	migrationPath := filepath.Join(workDir, "database", "migrations")
	if os.PathSeparator == '\\' {
		migrationPath = strings.ReplaceAll(migrationPath, "\\", "/")
	}
	return migrationPath
}

func RunMigrationCommand(direction string) error {
	LoadEnv()

	migrationsPath := GetMigrationPath()
	log.Printf("Chemin des migrations : %s", migrationsPath)

	// Vérifier si le répertoire existe
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		log.Printf("Le répertoire de migrations n'existe pas. Création...")
		if err := os.MkdirAll(migrationsPath, 0755); err != nil {
			return fmt.Errorf("erreur lors de la création du répertoire de migrations : %v", err)
		}
	}

	// Créer la migration
	m, err := migrate.New(
		"file://"+migrationsPath,
		GetDSN(),
	)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la migration : %v", err)
	}

	// Exécuter la migration
	switch direction {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		return fmt.Errorf("direction de migration non prise en charge : %s", direction)
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Println(err)
		return fmt.Errorf("erreur lors de l'exécution des migrations : %v", err)
	}

	log.Println("Migrations appliquées avec succès")
	return nil
}

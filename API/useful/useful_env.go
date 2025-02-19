package useful

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Aucun fichier .env trouvé, utilisation des variables d'environnement système")
	}
}

func GetEnvInt(key string, defaultVal int) int {
	if os.Getenv(key) == "" {
		log.Fatal("La variable d'environnement " + key + " n'est pas définie")
	}
	if val, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return val
	}
	return defaultVal
}

func GetEnvBool(key string, defaultVal bool) bool {
	if os.Getenv(key) == "" {
		log.Fatal("La variable d'environnement " + key + " n'est pas définie")
	}
	if key == "DEV" {
		return false
	} else if key == "PROD" {
		return true
	}
	return defaultVal
}

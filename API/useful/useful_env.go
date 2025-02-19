package useful

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement système")
	}
}

func getEnvInt(key string, defaultVal int) int {
    if val, err := strconv.Atoi(os.Getenv(key)); err == nil {
        return val
    }
    return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val, err := strconv.ParseBool(os.Getenv(key)); err == nil {
		return val
	}
	return defaultVal
}
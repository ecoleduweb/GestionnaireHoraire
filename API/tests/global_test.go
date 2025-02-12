package tests

import (
	"fmt"
	"log"
	"os"
)

func changeCurrentDiretory() {
	// Obtenir le répertoire courant actuel
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erreur lors de la récupération du répertoire courant :", err)
		return
	}
	fmt.Println("Répertoire courant actuel :", currentDir)

	// Remonter d'un niveau dans l'arborescence des répertoires pour trouver le .env
	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("Erreur lors du changement de répertoire :%v", err)
		return
	}

	// Vérifier le nouveau répertoire courant
	updatedDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erreur lors de la récupération du répertoire courant :", err)
		return
	}
	fmt.Println("Nouveau répertoire courant :", updatedDir)
}

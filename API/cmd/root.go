package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "llio-api",
	Short: "LLIO API - L'API de l'application LLIO",
	Long:  `LLIO API est l'API backend pour l'application LLIO.`,
}

// Execute ajoute toutes les commandes enfants au rootCmd et configure correctement les flags.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Ajouter ici les flags globaux et les configurations de commande
}

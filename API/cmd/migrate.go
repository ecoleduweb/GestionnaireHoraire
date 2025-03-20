package cmd

import (
	"fmt"
	"llio-api/useful"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Gère les migrations de base de données",
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Exécute toutes les migrations",
	Run: func(cmd *cobra.Command, args []string) {
		useful.RunMigrationCommand("up")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Annule la dernière migration",
	Run: func(cmd *cobra.Command, args []string) {
		useful.RunMigrationCommand("down")
	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create [nom]",
	Short: "Crée une nouvelle migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createMigration(args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateCreateCmd)
}

func createMigration(name string) {
	timestamp := time.Now().Format("20060102150405")
	migrationName := fmt.Sprintf("%s_%s", timestamp, name)

	migrationsPath := useful.GetMigrationPath()

	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		fmt.Printf("Le répertoire de migrations n'existe pas. Création...\n")
		if err := os.MkdirAll(migrationsPath, 0755); err != nil {
			fmt.Printf("Erreur lors de la création du répertoire de migrations : %v\n", err)
			os.Exit(1)
		}
	}

	// Créer les fichiers de migration manuellement
	upFileName := filepath.Join(migrationsPath, fmt.Sprintf("%s.up.sql", migrationName))
	downFileName := filepath.Join(migrationsPath, fmt.Sprintf("%s.down.sql", migrationName))

	// Créer des fichiers vides
	if err := os.WriteFile(upFileName, []byte("-- Migration Up\n"), 0644); err != nil {
		fmt.Printf("Erreur lors de la création du fichier de migration up : %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(downFileName, []byte("-- Migration Down\n"), 0644); err != nil {
		fmt.Printf("Erreur lors de la création du fichier de migration down : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Migration créée : %s\n", migrationName)
	fmt.Printf("Fichiers créés :\n- %s\n- %s\n", upFileName, downFileName)
}

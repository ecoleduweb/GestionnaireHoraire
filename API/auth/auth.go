package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azureadv2"
)

func NewAuth() {
	// Chargement des variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement des variables d'environnement")
	}

	// Config de l'authentification Microsoft
	config := struct {
		AureAdClientID  string
		AzureAdSecret   string
		AzureAdCallback string
	}{
		AureAdClientID:  os.Getenv("AZUREAD_CLIENT_ID"),
		AzureAdSecret:   os.Getenv("AZUREAD_CLIENT_SECRET"),
		AzureAdCallback: os.Getenv("AZUREAD_CALLBACK_URL"),
	}

	// Vérification des variables d'environnement Microsoft
	if config.AureAdClientID == "" || config.AzureAdSecret == "" || config.AzureAdCallback == "" {
		log.Fatal("Les variables d'environnement AzureAD ne sont pas définies")
	}

	sessionKey := os.Getenv("SESSION_SECRET")
	// Vérification de la clé de session
	if sessionKey == "" {
		log.Fatal("La clé de session n'est pas définie")
	}

	// Configuration de la session
	store := sessions.NewCookieStore([]byte(sessionKey))
	store.MaxAge(86400 * 30)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false // Mettre a true en production (https)

	gothic.Store = store

	// Configuration de goth
	goth.UseProviders(
		azureadv2.New(
			config.AureAdClientID,
			config.AzureAdSecret,
			config.AzureAdCallback,
			azureadv2.ProviderOptions{
				Tenant: azureadv2.OrganizationsTenant,
				Scopes: []azureadv2.ScopeType{
					azureadv2.OpenIDScope,
					azureadv2.OfflineAccessScope,
					azureadv2.UserReadScope,
				},
			},
		),
	)
}

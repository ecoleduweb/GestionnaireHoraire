package auth

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azureadv2"
)

func NewAuth() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement des variables d'environnement")
	}

	authenticationConfig := struct {
		AureAdClientID  string
		AzureAdSecret   string
		AzureAdCallback string
	}{
		AureAdClientID:  os.Getenv("AZUREAD_CLIENT_ID"),
		AzureAdSecret:   os.Getenv("AZUREAD_CLIENT_SECRET"),
		AzureAdCallback: os.Getenv("AZUREAD_CALLBACK_URL"),
	}
	if authenticationConfig.AureAdClientID == "" || authenticationConfig.AzureAdSecret == "" || authenticationConfig.AzureAdCallback == "" {
		log.Fatal("Les variables d'environnement AzureAD ne sont pas définies")
	}

	sessionsConfig := struct {
		SessionKey string
		SessionMaxAge int
		HttpOnly bool
		IsProduction bool
	}{
		SessionKey: os.Getenv("SESSION_SECRET"),
		SessionMaxAge: 86400 * 30,
		HttpOnly: true,
		IsProduction: false,
	}
	if sessionsConfig.SessionKey == "" {
		log.Fatal("La variable d'environnement SESSION_SECRET n'est pas définie")
	}


	// Configuration de la session
	store := sessions.NewCookieStore([]byte(sessionsConfig.SessionKey))
	store.MaxAge(sessionsConfig.SessionMaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = sessionsConfig.HttpOnly
	store.Options.Secure = sessionsConfig.IsProduction
	store.Options.SameSite = http.SameSiteLaxMode

	gothic.Store = store


	goth.UseProviders(
		azureadv2.New(
			authenticationConfig.AureAdClientID,
			authenticationConfig.AzureAdSecret,
			authenticationConfig.AzureAdCallback,
			azureadv2.ProviderOptions{
				Tenant: azureadv2.OrganizationsTenant,
				Scopes: []azureadv2.ScopeType{
					azureadv2.OpenIDScope,
				},
			},
		),
	)
}

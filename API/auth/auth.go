package auth

import (
	"llio-api/useful"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azureadv2"
)

func NewAuth() {
	useful.LoadEnv()
	authenticationConfig := struct {
		AzureAdClientID string
		AzureAdSecret   string
		AzureAdCallback string
	}{
		AzureAdClientID: os.Getenv("AZUREAD_CLIENT_ID"),
		AzureAdSecret:   os.Getenv("AZUREAD_CLIENT_SECRET"),
		AzureAdCallback: os.Getenv("AZUREAD_CALLBACK_URL"),
	}
	if authenticationConfig.AzureAdClientID == "" || authenticationConfig.AzureAdSecret == "" || authenticationConfig.AzureAdCallback == "" {
		log.Fatal("Les variables d'environnement AzureAD ne sont pas définies")
	}

	sessionsConfig := struct {
		SessionKey    string
		SessionMaxAge int
		HttpOnly      bool
		IsProduction  bool
	}{
		SessionKey:    os.Getenv("SESSION_SECRET"),
		SessionMaxAge: useful.GetEnvInt("SESSION_MAX_AGE", 86400),
		HttpOnly:      true,
		IsProduction:  useful.GetEnvBool("ENV", false),
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
			authenticationConfig.AzureAdClientID,
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

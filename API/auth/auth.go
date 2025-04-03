package auth

import (
	"llio-api/models/enums"
	"llio-api/useful"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azureadv2"
)

func AuthWithAzure() {
	useful.LoadEnv()
	// Structure de configuration pour AzureAD
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

	// Structure de configuration pour les sessions
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

	// Configuration du fournisseur AzureAD (ajouter d'autres fournisseurs si nécessaire)
	goth.UseProviders(
		azureadv2.New(
			authenticationConfig.AzureAdClientID,
			authenticationConfig.AzureAdSecret,
			authenticationConfig.AzureAdCallback,
			azureadv2.ProviderOptions{
				Tenant: azureadv2.OrganizationsTenant,
				Scopes: []azureadv2.ScopeType{
					azureadv2.OpenIDScope,
					azureadv2.EmailScope,
					azureadv2.CalendarsReadScope,
				},
			},
		),
	)
}

func CreateJWTToken(userEmail string, fisrtName string, lastName string, expiresAt time.Time, userRole enums.UserRole) (string, error) {
	// Define the claims for the token
	secretKey := os.Getenv("JWT_SECRET_KEY")
	claims := jwt.MapClaims{
		"email":      userEmail,
		"first_name": fisrtName,
		"last_name":  lastName,
		"exp":        expiresAt.Unix(),
		"role":       userRole, // TODO aller chercher le role dans la base de données et le passer en paramètre ici
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

package middleware

import (
	"context"
	"fmt"
	"os"
	"llio-api/useful"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)


func ValidateMicrosoftToken(accessToken string) (*jwt.Token, error) {
	useful.LoadEnv()
	tenantID := os.Getenv("AZUREAD_TENANT_ID")
	jwksURL := fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/v2.0/keys", tenantID)

	// Récupérer les clés JWK (clés publiques pour vérifier la signature du token) depuis Microsoft
	keySet, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		log.Printf("Erreur lors de la récupération des clés JWK: %v", err)
	}

	// Validation du token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("méthode de signature inattendue: %v", token.Header["alg"])
		}

		// Récupérer du kid (identifiant de la clé) du token 
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid manquant dans le header du token")
		}

		// Utiliser LookupKeyID pour trouver la clé correspondante avec le kid du token
		key, found := keySet.LookupKeyID(kid)
		if !found {
			return nil, fmt.Errorf("clé JWK non trouvée pour le kid: %s", kid)
		}

		// Extraire la clé
		var rawKey interface{}
		if err := key.Raw(&rawKey); err != nil {
			return nil, fmt.Errorf("erreur lors de l'extraction de la clé JWK: %v", err)
		}

		return rawKey, nil
	})

	if err != nil {
		log.Printf("Erreur lors de la validation du token: %v", err)
		return nil, err
	}

	return token, nil
}

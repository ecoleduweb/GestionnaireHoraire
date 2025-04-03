package middleware

import (
	"crypto/rsa"
	"fmt"
	"llio-api/useful"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
)

func ValidateMicrosoftToken(accessToken string, c *gin.Context) (*jwt.Token, error) {
	useful.LoadEnv()
	tenantID := os.Getenv("AZUREAD_TENANT_ID")
	jwksURL := fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/v2.0/keys", tenantID)
	keySet, err := jwk.Fetch(c, jwksURL)

	log.Printf("VALIDATING TOKEN")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwa.RS256.String() {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid header not found")
		}

		keys, ok := keySet.LookupKeyID(kid)
		if !ok {
			return nil, fmt.Errorf("key %v not found", kid)
		}

		publickey := &rsa.PublicKey{}
		err = keys.Raw(publickey)
		if err != nil {
			return nil, fmt.Errorf("could not parse pubkey")
		}

		return publickey, nil
	})

	if err != nil {
		fmt.Errorf("could not parse pubkey %v", err)
		return nil, err
	}
	return token, nil
}

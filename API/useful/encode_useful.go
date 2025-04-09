package useful

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetupAuthProvider(c *gin.Context) {
	provider := c.Param("provider")
	log.Printf("Configuration de la requête avec provider %s", provider)
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
}
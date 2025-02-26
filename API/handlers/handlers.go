package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func ApiStatus(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Pong",
		})
	})

	r.GET("/health/status", func(c *gin.Context) {
		// Calculer la durée écoulée depuis le démarrage
		uptime := time.Since(startTime)

		// Supprimer les millisecondes
		uptime = uptime.Truncate(time.Second)

		// Formater la durée en heures, minutes et secondes
		hours := int(uptime.Hours())
		minutes := int(uptime.Minutes()) % 60
		seconds := int(uptime.Seconds()) % 60

		// Créer une chaîne formatée
		formattedUptime := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

		c.JSON(http.StatusOK, gin.H{
			"status": "Healthy",
			"uptime": formattedUptime,
		})
	})
}

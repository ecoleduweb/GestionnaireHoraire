package services

import (
	"fmt"
	"llio-api/models/DTOs"
	"llio-api/useful"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Fonction qui permet de valider la structure et le contenu du JSON attendu
// Traite les erreurs de validation de champs, de parsing de date et autres erreurs
func VerifyJSON(c *gin.Context, dto interface{}) []DTOs.FieldErrorDTO {
	var fieldErrors []DTOs.FieldErrorDTO
	if err := c.ShouldBindJSON(dto); err != nil {
		// Toutes les listes d'erreurs possibles
		var validationErrors validator.ValidationErrors
		var parseErrors []*time.ParseError
		var otherErrors []error

		// Utiliser une assertion de type pour accumuler toutes les erreurs
		switch e := err.(type) {
		case validator.ValidationErrors:
			validationErrors = e
		case *time.ParseError:
			parseErrors = append(parseErrors, e)
		default:
			otherErrors = append(otherErrors, e)
		}

		// traitement des erreurs de validation
		for _, err := range validationErrors {
			fiedlName := err.Field()
			camelCaseFieldName := useful.ToCamelCase(fiedlName)
			fieldErrors = append(fieldErrors, DTOs.FieldErrorDTO{
				Field:   camelCaseFieldName,
				Message: fmt.Sprintf("Le champ %s est invalide ou manquant", camelCaseFieldName),
			})
		}

		// Traitement des erreurs de parsing de date
		for _, err := range parseErrors {
			fieldErrors = append(fieldErrors, DTOs.FieldErrorDTO{
				Field:   "date",
				Message: fmt.Sprintf("Erreur de parsing de la date : %s", err.Error()),
			})
		}

		// Traitement des autres erreurs
		for _, err := range otherErrors {
			fieldErrors = append(fieldErrors, DTOs.FieldErrorDTO{
				Field:   "unknown",
				Message: fmt.Sprintf("Erreur inconnue : %s", err.Error()),
			})
		}
	}
	return fieldErrors
}

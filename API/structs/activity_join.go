package structs

import (
	"llio-api/models/DAOs"
)

// Structure temporaire pour recevoir les résultats du JOIN
type ActivityJoin struct {
	DAOs.Activity
	ProjectName string 
}
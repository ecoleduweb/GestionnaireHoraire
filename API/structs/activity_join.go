package structs

import (
	"llio-api/models/DAOs"
)

type ActivityJoin struct {
	DAOs.Activity
	ProjectName string
}

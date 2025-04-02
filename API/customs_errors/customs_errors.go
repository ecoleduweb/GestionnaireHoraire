package customs_errors

import (
	"errors"
)

// Types d'erreurs personnalisées pour la base de données
var (
	ErrDuplicateEntry = errors.New("duplication")
	ErrNotFound       = errors.New("ressource introuvable")
	ErrDatabase       = errors.New("erreur de la BD")
)

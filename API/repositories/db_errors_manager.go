package repositories

import (
	"errors"
	"llio-api/customs_errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func DBErrorManager(err error) error {
	if err == nil {
		return nil
	}

	// Vérification spécifique pour MySQL
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		// Code 1062 est pour "Duplicate entry"
		if mysqlErr.Number == 1062 {
			return customs_errors.ErrDuplicateEntry
		}
	}

	// Pour les erreurs "not found" (selon votre ORM, cette implémentation peut varier)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customs_errors.ErrNotFound
	}

	// Erreur par défaut
	return customs_errors.ErrDatabase
}

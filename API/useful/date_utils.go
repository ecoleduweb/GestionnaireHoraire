package useful

import (
	"llio-api/customs_errors"
	"time"
)

// DateToISOString Formatter la date dans le format RFC3339 (YYYY-MM-DDThh:mm:ssZ), notamment utile pour l'API Outlook.
func DateToISOString(date time.Time) string {
	return date.Format(time.RFC3339)
}

func ToStartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func ToEndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
}

// AsTorontoThenUTC prends une heure existante, peu importe son fuseau horaire, et la considère comme étant du fuseau
// horaire local. Elle la reconverti alors en heure UTC, ce qui peut être utile pour certains services externes comme
// Graph API.
func AsTorontoThenUTC(date time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		return time.Time{}, customs_errors.ErrCantLoadLocalTimezone
	}

	// Reprends les mêmes chiffres que la date originale, mais lui applique le fuseau horaire America/Toronto plutôt que
	// celui existant/celui du serveur
	toronto := time.Date(
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second(), date.Nanosecond(),
		loc,
	)

	return toronto.UTC(), nil
}

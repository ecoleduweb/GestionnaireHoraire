package services

import (
	"llio-api/models/DTOs"
	"llio-api/useful"

	"encoding/json"
	"fmt"
	"time"
)

const GraphApiBaseUrl = "https://graph.microsoft.com/v1.0"

func GetCalendarEvents(accessToken string, date time.Time) ([]DTOs.GraphEvent, error) {
	startOfDay := useful.ToStartOfDay(date)
	endOfDay := useful.ToEndOfDay(date)

	startOfDayUTC, err := useful.AsTorontoThenUTC(startOfDay)
	if err != nil {
		return nil, err
	}

	endOfDayUTC, err := useful.AsTorontoThenUTC(endOfDay)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"%s/me/calendarView?startDateTime=%s&endDateTime=%s",
		GraphApiBaseUrl,
		useful.DateToISOString(startOfDayUTC),
		useful.DateToISOString(endOfDayUTC),
	)

	body, err := GraphApiGetRequest(url, accessToken)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []DTOs.GraphEvent `json:"value"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Value, nil
}

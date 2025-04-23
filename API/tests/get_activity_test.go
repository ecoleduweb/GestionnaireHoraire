package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUsersActivities(t *testing.T) {
	w := sendRequest(router, "GET", "/activities/me", nil)
	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetActivity(t *testing.T) {
	w := sendRequest(router, "GET", fmt.Sprintf("/activity/%d", doNotDeleteActivity.Id), nil)
	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetNotFoundActivity(t *testing.T) {
	w := sendRequest(router, "GET", "/activity/0", nil)
	assertResponse(t, w, http.StatusNotFound, nil)
	assert.NotNil(t, w.Body)
}

func TestGetFromToActivities(t *testing.T) {
	startDate := time.Now().Format("2006-01-02")
	endDate := time.Now().AddDate(0, 0, 7).Format("2006-01-02")

	url := fmt.Sprintf("/activities/%d/%s/%s", doNotDeleteUser.Id, startDate, endDate)

	w := sendRequest(router, "GET", url, nil)
	assertResponse(t, w, http.StatusNotFound, nil)
	assert.NotNil(t, w.Body)
}


package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	w := sendRequest(router, "GET", "/activities", nil)
	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetCategory(t *testing.T) {
	w := sendRequest(router, "GET", fmt.Sprintf("/activity/%d", doNotDeleteActivity.Id), nil)
	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetNotFoundCategorie(t *testing.T) {
	w := sendRequest(router, "GET", "/activity/0", nil)
	assertResponse(t, w, http.StatusNotFound, nil)
}

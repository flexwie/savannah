package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	router "felixwie.com/savannah/api"
	"github.com/stretchr/testify/assert"
)

func TestReceiveWebhook(t *testing.T) {
	router := router.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/webhook/123", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

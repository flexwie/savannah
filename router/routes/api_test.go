package routes_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"felixwie.com/savannah/config"
	"felixwie.com/savannah/router"
	"felixwie.com/savannah/router/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	router := router.GetRouter()
	routes.Cfg.Source = []config.Source{{Name: "test", Branch: "test", Folder: "test", URL: "test"}}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/repositories", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	log.Printf("%v", string(w.Body.Bytes()))
	var data []config.Source
	json.Unmarshal(w.Body.Bytes(), data)

	assert.Equal(t, "test", data[0].Name)
}

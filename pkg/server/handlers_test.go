package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/nirvana-assessment/pkg/configs"
	"github.com/ariel17/nirvana-assessment/pkg/services"
)

func TestCoalesceHandler(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		statusCode int
		response   *services.Response
		successful bool
	}{
		{"correct member_id", "/?member_id=1", http.StatusOK, &services.Response{1066, 11000, 5666}, true},
		{"empty member_id", "/", http.StatusBadRequest, nil, true},
		{"non-numeric member_id", "/?member_id=x", http.StatusBadRequest, nil, true},
		{"error by falling api call", "/?member_id=1", http.StatusInternalServerError, nil, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.Default()
			r.GET("/", CoalesceHandler)
			configs.ShouldAPIFail = !tc.successful

			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)
			assert.Equal(t, tc.statusCode, rr.Code)

			if tc.statusCode == http.StatusOK {
				b, _ := json.Marshal(tc.response)
				assert.Equal(t, string(b), rr.Body.String())
			}
		})
	}
}
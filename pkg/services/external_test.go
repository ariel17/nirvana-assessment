package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/nirvana-assessment/pkg/configs"
)

func TestGetAPI(t *testing.T) {
	testCases := []struct{
		name string
		f func() Response
		expected Response
	}{
		{"api 1", GetAPI1, Response{1000, 10000, 5000}},
		{"api 2", GetAPI2, Response{1200, 13000, 6000}},
		{"api 3", GetAPI3, Response{1000, 10000, 6000}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			response := tc.f()
			duration := time.Since(start)

			assert.LessOrEqual(t, time.Duration(configs.MaxMockedResponseTimeInMillis), duration)
			assert.Equal(t, tc.expected, response)
		})
	}
}
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/nirvana-assessment/pkg/configs"
)

func TestStrategies(t *testing.T) {
	responses := []*Response{
		&Response{1000, 500, 100},
		&Response{2000, 400, 110},
		&Response{3000, 300, 120},
	}

	t.Run("average", func(t *testing.T) {
		as := averageStrategy{}
		r := as.Coalesce(responses)
		assert.Equal(t, 2000, r.Deductible)
		assert.Equal(t, 400, r.StopLoss)
		assert.Equal(t, 110, r.OopMax)
	})

	t.Run("sum", func(t *testing.T) {
		ss := sumStrategy{}
		r := ss.Coalesce(responses)
		assert.Equal(t, 6000, r.Deductible)
		assert.Equal(t, 1200, r.StopLoss)
		assert.Equal(t, 330, r.OopMax)
	})
}

func TestCoalesceAPIResponses(t *testing.T) {
	testCases := []struct{
		name string
		strategy string
		expected Response
	}{
		{"average", configs.AverageCoalesceStrategy, Response{1066, 11000, 5666}},
		{"sum", configs.SumCoalesceStrategy, Response{3200, 33000, 17000}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			configs.CoalesceStrategyToUse = tc.strategy
			r, err := CoalesceAPIResponses(0)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, *r)
		})
	}
}
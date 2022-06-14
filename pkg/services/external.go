package services

import (
	"math/rand"
	"time"

	"github.com/ariel17/nirvana-assessment/pkg/configs"
)

// Response is the structure response from mocked APIs.
type Response struct {
	Deductible int `json:"deductible"`
	StopLoss   int `json:"stop_loss"`
	OopMax     int `json:"oop_max"`
}

// GetAPI1 returns a mocked external API response as #1.
func GetAPI1() Response{
	return getAPI(1000, 10000, 5000)
}

// GetAPI2 returns a mocked external API response as #2.
func GetAPI2() Response{
	return getAPI(1200, 13000, 6000)
}

// GetAPI3 returns a mocked external API response as #3.
func GetAPI3() Response{
	return getAPI(1000, 10000, 6000)
}

func getAPI(deductible, stopLoss, oopMax int) Response {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(configs.MaxMockedResponseTimeInMillis)
	time.Sleep(time.Duration(n)*time.Millisecond)
	return Response{
		Deductible: deductible,
		StopLoss: stopLoss,
		OopMax: oopMax,
	}
}
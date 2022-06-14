package services

import (
	"github.com/ariel17/nirvana-assessment/pkg/configs"
)

type API func(int) (*Response, error)

// CoalesceAPIResponses collects external API responses in async way to coalesce
// the result into a single response based on the configured strategy.
func CoalesceAPIResponses(memberID int) (*Response, error) {
	apis := []API{GetAPI1, GetAPI2, GetAPI3}
	inputChan := make(chan API, len(apis))

	errorsChan := make(chan error, len(apis))
	defer close(errorsChan)

	responsesChan := make(chan *Response, len(apis))
	defer close(responsesChan)

	// spawning workers
	for i := 0; i < configs.WorkersCount; i++ {
		go func() {
			for f := range inputChan {
				r, err := f(memberID)
				responsesChan <- r
				errorsChan <- err
			}
		}()
	}

	// provisioning
	for _, f := range apis {
		inputChan <- f
	}
	close(inputChan)

	// collecting async data
	var responses []*Response
	for range apis {
		if err := <-errorsChan; err != nil {
			return nil, err
		}
		responses = append(responses, <-responsesChan)
	}

	s := newStrategy()
	return s.Coalesce(responses), nil
}

type strategy interface {
	Coalesce([]*Response) *Response
}

func newStrategy() strategy {
	if configs.CoalesceStrategyToUse == configs.AverageCoalesceStrategy {
		return &averageStrategy{}
	}
	return &sumStrategy{}
}

type averageStrategy struct{}

func (as *averageStrategy) Coalesce(responses []*Response) *Response {
	var deductible, stopLoss, oopMax int
	for _, r := range responses {
		deductible += r.Deductible
		stopLoss += r.StopLoss
		oopMax += r.OopMax
	}
	return &Response{
		Deductible: deductible / len(responses),
		StopLoss:   stopLoss / len(responses),
		OopMax:     oopMax / len(responses),
	}
}

type sumStrategy struct{}

func (as *sumStrategy) Coalesce(responses []*Response) *Response {
	var deductible, stopLoss, oopMax int
	for _, r := range responses {
		deductible += r.Deductible
		stopLoss += r.StopLoss
		oopMax += r.OopMax
	}
	return &Response{
		Deductible: deductible,
		StopLoss:   stopLoss,
		OopMax:     oopMax,
	}
}
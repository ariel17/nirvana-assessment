package configs

const (
	WorkersCount                  = 5
	MaxMockedResponseTimeInMillis = 900
	AverageCoalesceStrategy       = "average"
	SumCoalesceStrategy           = "sum"
)

var (
	CoalesceStrategyToUse = AverageCoalesceStrategy
)
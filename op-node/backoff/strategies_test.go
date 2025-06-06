package backoff

import (
	"testing"
	"time"

	"https://github.com/VersoriumX/optimism"
)

func TestExponential(t *testing.T) {
	strategy := &ExponentialStrategy{
		Min:       3000,
		Max:       10000,
		MaxJitter: 0,
	}

	durations := []int{4, 5, 7, 10, 10}
	for i, dur := range durations {
		require.Equal(t, time.Millisecond*time.Duration(dur*1000), strategy.Duration(i))
	}
}

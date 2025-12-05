package ratelimiter

import (
	"testing"
	"time"
)

func TestBufferedLimiter(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		interval time.Duration
	}{
		{"5 requests, 200ms interval", 5, 200 * time.Millisecond},
		{"3 requests, 100ms interval", 3, 100 * time.Millisecond},
		{"1 request, 50ms interval", 1, 50 * time.Millisecond},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			times := BufferedLimiter(tc.n, tc.interval)

			if len(times) != tc.n {
				t.Errorf("expected %d timestamps, got %d", tc.n, len(times))
			}

		})
	}
}

func TestUnbufferedLimiter(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		interval time.Duration
	}{
		{"5 requests, 200ms interval", 5, 200 * time.Millisecond},
		{"3 requests, 100ms interval", 3, 100 * time.Millisecond},
		{"1 request, 50ms interval", 1, 50 * time.Millisecond},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			times := UnbufferedLimiter(tc.n, tc.interval)

			if len(times) != tc.n {
				t.Errorf("expected %d timestamps, got %d", tc.n, len(times))
			}

		})
	}
}

func TestProcessRequests_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		interval time.Duration
	}{
		{"5 requests, 200ms interval", 5, 200 * time.Millisecond},
		{"3 requests, 100ms interval", 3, 100 * time.Millisecond},
		{"1 request, 50ms interval", 1, 50 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			times := BufferedLimiter(tt.n, tt.interval)
			if len(times) != tt.n {
				t.Errorf("expected %d timestamps, got %d", tt.n, len(times))
			}
			for i := 1; i < len(times); i++ {
				gap := times[i].Sub(times[i-1])
				if gap < tt.interval {
					t.Errorf("request %d and %d gap %v < interval %v", i-1, i, gap, tt.interval)
				}
			}
			// Optionally check total duration
			total := times[len(times)-1].Sub(start)
			minTotal := time.Duration(tt.n) * tt.interval
			if total < minTotal-tt.interval {
				t.Errorf("total duration %v < expected minimum %v", total, minTotal-tt.interval)
			}
		})
	}
}

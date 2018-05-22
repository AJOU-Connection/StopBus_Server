package main

import (
	"net/http"
	"sync"
	"testing"
)

func TestTargetObserver(t *testing.T) {
	var wg sync.WaitGroup

	tt := []struct {
		routeID    string
		stationID  string
		httpStatus int
		resultCode int
	}{
		{"223000100", "203000066", http.StatusOK, 0},
		{"234000026", "203000066", http.StatusOK, 0},
		{"234000024", "203000066", http.StatusOK, 0},
		{"200000053", "203000066", http.StatusOK, 0},
		{"200000110", "203000066", http.StatusOK, 0},
		{"200000112", "203000066", http.StatusOK, 0},
		{"200000144", "203000066", http.StatusOK, 0},
	}
	for _, tc := range tt {
		wg.Add(1)
		defer wg.Done()
		go TargetObserver(tc.routeID, tc.stationID)
	}

	wg.Wait()
}

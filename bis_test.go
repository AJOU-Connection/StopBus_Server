package StopBus

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchForStation(t *testing.T) {
	tt := []struct {
		keyword    string
		status     int
		resultCode int
	}{
		{"아주대학교입구", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := SearchForStation(tc.keyword)
		fmt.Println(resultData)
	}
}

func TestSearchForRoute(t *testing.T) {
	tt := []struct {
		keyword    string
		status     int
		resultCode int
	}{
		{"3007", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := SearchForRoute(tc.keyword)
		fmt.Println(resultData)
	}
}

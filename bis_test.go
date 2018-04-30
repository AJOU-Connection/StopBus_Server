package StopBus

import (
	"testing"
)

func TestSearchForStation(t *testing.T) {
	SearchForStation("아주대학교입구")
}

func TestSearchForRoute(t *testing.T) {
	SearchForRoute("14")
}

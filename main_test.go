package main

import (
	"fmt"
	"testing"
)

func TestSetUpConfig(t *testing.T) {
	setUpConfig()
	if (configuration{}) == config {
		fmt.Println("HERE")
	}
}

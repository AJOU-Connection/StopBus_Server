package main

import "testing"

func TestGetInAlert(t *testing.T) {
	GetInAlert("234000026", "203000066")
}

func TestGetOutAlert(t *testing.T) {
	GetOutAlert("234000026", "203000066", "경기00바0000")
}

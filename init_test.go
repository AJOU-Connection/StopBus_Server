package main

import (
	"os"
	"testing"
)

func TestSetUpConfig(t *testing.T) {
	expectedErrorMessage := [3]string{"directory not exists: ./configs", "config file not exists: ./configs/config.json", "invalid JSON file: ./configs/config.json"}

	err := setUpConfig()
	if err == nil {
		return
	}

	if err.Error() == expectedErrorMessage[0] {
		os.Mkdir("configs", os.ModeDir)
	} else if err.Error() == expectedErrorMessage[1] {
		fi, _ := os.OpenFile("configs/config.json", os.O_RDWR|os.O_CREATE, 0755)
		fi.Close()
	} else if err.Error() == expectedErrorMessage[2] {
		os.Rename("configs/config.json", "configs/_config.json")
		fo, _ := os.OpenFile("configs/config.json", os.O_RDWR|os.O_CREATE, 0755)
		fo.Write(([]byte)("{\"serviceKey\":\"testServiceKey\"}"))
		fo.Close()
	} else {
		t.Errorf("not expected error: %v", err)
	}

}

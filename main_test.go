package StopBus

import "testing"

func TestSetUpConfig(t *testing.T) {
	setUpConfig()
	if config == (configuration{}) {
		t.Errorf("serviceKey is not.")
	}
}

func TestMain(t *testing.T) {
	main()
}

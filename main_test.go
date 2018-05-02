package StopBus

import "testing"

func TestSetUpConfig(t *testing.T) {
	setUpConfig()
	if config == (configuration{}) {
		t.Logf("serviceKey is not.")
	}
}

func TestMain(t *testing.T) {
	main()
}

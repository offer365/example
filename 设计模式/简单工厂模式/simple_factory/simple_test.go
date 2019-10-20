package simple_factory

import "testing"

func TestNewAPI(t *testing.T) {
	NewAPI(1).Say("tom")
	NewAPI(2).Say("tom")
}

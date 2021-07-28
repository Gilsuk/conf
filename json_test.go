package conf_test

import (
	"testing"

	"github.com/gilsuk/conf"
)

type configuration struct {
	Name string
	Age  int
}

func TestLoadNonExistingJsonFile(t *testing.T) {
	config := &configuration{}
	err := conf.LoadJson("nonexisting.json", config)
	if err == nil {
		t.Fail()
	}
}

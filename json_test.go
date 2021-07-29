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

func TestLoadJsonCompletly(t *testing.T) {
	config := &configuration{}
	err := conf.LoadJson("./testdata/configuration.json", config)
	if err != nil {
		t.Errorf("fail when load file")
	}

	if config.Name != "Alice" || config.Age != 19 {
		t.Fail()
	}
}

func TestOverrideDefault(t *testing.T) {
	config := &configuration{
		Name: "Bob",
	}

	err := conf.LoadJson("./testdata/configuration.json", config)
	if err != nil {
		t.Errorf("fail when load file")
	}

	if config.Name != "Alice" || config.Age != 21 {
		t.Fail()
	}
}

func TestLoadDefault(t *testing.T) {
	config := &configuration{
		Name: "Bob",
	}

	err := conf.LoadJson("./testdata/configuration-with-default.json", config)
	if err != nil {
		t.Errorf("fail when load file")
	}

	if config.Name != "Bob" || config.Age != 21 {
		t.Fail()
	}
}

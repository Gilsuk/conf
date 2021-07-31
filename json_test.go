package conf_test

import (
	"path/filepath"
	"testing"

	"github.com/gilsuk/conf"
)

type configuration struct {
	Name string
	Age  int
}

func TestLoadNonExistingJsonFile(t *testing.T) {
	config := &configuration{}

	err := conf.Load(conf.JSON, "nonexisting.json", config)
	if err == nil {
		t.Fail()
	}
}

func TestLoadJsonCompletly(t *testing.T) {
	config := &configuration{}
	loadJsonHelper(t, "configuration.json", config)

	if config.Name != "Alice" || config.Age != 19 {
		t.Fail()
	}
}

func TestOverrideDefault(t *testing.T) {
	config := &configuration{
		Name: "Bob",
	}

	loadJsonHelper(t, "configuration.json", config)

	if config.Name != "Alice" || config.Age != 19 {
		t.Fail()
	}
}

func TestLoadDefault(t *testing.T) {
	config := &configuration{
		Name: "Bob",
	}

	loadJsonHelper(t, "configuration-with-default.json", config)

	if config.Name != "Bob" || config.Age != 21 {
		t.Fail()
	}
}

func loadJsonHelper(t *testing.T, fileName string, config interface{}) {
	t.Helper()
	err := conf.Load(conf.JSON, filepath.Join("testdata", fileName), config)
	if err != nil {
		t.Errorf("fail when load file")
	}
}

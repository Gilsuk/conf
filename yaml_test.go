package conf_test

import (
	"path/filepath"
	"testing"

	"github.com/gilsuk/conf"
)

func TestLoadNonExistingYamlFile(t *testing.T) {
	config := &configuration{}

	err := conf.Load(conf.YAML, "nonexisting.yaml", config)
	if err == nil {
		t.Fail()
	}
}

func TestLoadYamlCompletly(t *testing.T) {
	config := &configuration{}
	loadYamlHelper(t, "configuration.yaml", config)

	if config.Name != "Alice" || config.Age != 19 {
		t.Fail()
	}
}

func TestOverrideYamlDefault(t *testing.T) {
	config := &configuration{
		Name: "Bob",
	}

	loadYamlHelper(t, "configuration.yaml", config)

	if config.Name != "Alice" || config.Age != 19 {
		t.Fail()
	}
}

func TestLoadYamlDefault(t *testing.T) {
	config := &configuration{
		Name: "Bob",
	}

	loadYamlHelper(t, "configuration-with-default.yaml", config)

	if config.Name != "Bob" || config.Age != 21 {
		t.Fail()
	}
}

func TestOutOverrideExistingYamlOutFile(t *testing.T) {
	configs := []configuration{
		{
			Name: "Cindy",
			Age:  18,
		},
		{
			Name: "Lauper",
			Age:  20,
		},
	}

	outFileName := "out.yaml"
	outPath := filepath.Join("testdata", outFileName)

	for i := 0; i < len(configs); i++ {
		if err := conf.Out(conf.YAML, outPath, configs[i]); err != nil {
			t.Fail()
		}
		config := &configuration{}
		loadYamlHelper(t, outFileName, config)

		if config.Name != configs[i].Name || config.Age != configs[i].Age {
			t.Fail()
		}
	}
}

func loadYamlHelper(t *testing.T, fileName string, config interface{}) {
	t.Helper()
	err := conf.Load(conf.YAML, filepath.Join("testdata", fileName), config)
	if err != nil {
		t.Error(err)
	}
}

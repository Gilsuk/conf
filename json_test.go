package conf_test

import (
	"os"
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

func TestOutOverrideExistingOutFile(t *testing.T) {
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

	outFileName := "out.json"
	outPath := filepath.Join("testdata", outFileName)

	for i := 0; i < len(configs); i++ {
		if err := conf.Out(conf.JSON, outPath, configs[i]); err != nil {
			t.Fail()
		}
		config := &configuration{}
		loadJsonHelper(t, outFileName, config)

		if config.Name != configs[i].Name || config.Age != configs[i].Age {
			t.Fail()
		}
	}
}

func loadJsonHelper(t *testing.T, fileName string, config interface{}) {
	t.Helper()
	err := conf.Load(conf.JSON, filepath.Join("testdata", fileName), config)
	if err != nil {
		t.Error(err)
	}
}

func tearDown() {
	trashes := []string{
		"out.json",
		"out.yaml",
	}

	for i := 0; i < len(trashes); i++ {
		os.Remove(filepath.Join("testdata", trashes[i]))
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	tearDown()
	os.Exit(exitCode)
}

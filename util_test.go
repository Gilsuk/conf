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

func loadFileHelper(t *testing.T, fileName string, config interface{}) {
	t.Helper()
	err := conf.Load(filepath.Join("testdata", fileName), config)
	if err != nil {
		t.Error(err)
	}
}

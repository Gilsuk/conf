package conf

import (
	"path/filepath"
	"strings"
)

func Load(path string, confStruct interface{}) error {
	worker := workerByFormat(path)
	return worker.load(path, confStruct)
}

func Out(path string, confStruct interface{}) error {
	worker := workerByFormat(path)
	return worker.out(path, confStruct)
}

func workerByFormat(path string) worker {
	switch ext := strings.ToLower(filepath.Ext(path))[1:]; ext {
	case "json":
		return newJsonWorker()
	case "yaml", "yml":
		return newYamlWorker()
	default:
		return newDefaultWorker(ext)
	}
}

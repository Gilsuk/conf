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
	switch ext := extension(path); ext {
	case "json":
		return newJsonWorker()
	case "yaml", "yml":
		return newYamlWorker()
	default:
		return newDefaultWorker(ext)
	}
}

func extension(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	if len(ext) > 0 {
		ext = ext[1:]
	}
	return ext
}

package conf

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type yamlWorker workerStruct

func newYamlWorker() worker {
	return new(yamlWorker)
}

func (w yamlWorker) load(path string, confStruct interface{}) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, confStruct)
}

func (w yamlWorker) out(path string, confStruct interface{}) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := yaml.Marshal(confStruct)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)

	return err
}

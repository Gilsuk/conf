package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type jsonWorker workerStruct

func newJsonWorker() worker {
	return new(jsonWorker)
}

func (w jsonWorker) load(path string, confStruct interface{}) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, confStruct)
}

func (w jsonWorker) out(path string, confStruct interface{}) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(confStruct, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)

	return err
}

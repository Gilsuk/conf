package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type jsonWorker struct{}

func (l *jsonWorker) load(path string, confStruct interface{}) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, confStruct)

	return nil
}

func (l *jsonWorker) out(path string, confStruct interface{}) error {
	return nil
}

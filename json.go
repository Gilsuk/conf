package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJson(path string, configuration interface{}) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, configuration)

	return nil
}

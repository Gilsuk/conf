package conf

import "fmt"

type defaultWorker struct{}

func (d *defaultWorker) load(path string, confStruct interface{}) error {
	return fmt.Errorf("This file format is not implemented")
}

func (d *defaultWorker) out(path string, confStruct interface{}) error {
	return fmt.Errorf("This file format is not implemented")
}

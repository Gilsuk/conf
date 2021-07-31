package conf

import "fmt"

type defaultWorker struct{}

func (w *defaultWorker) load(path string, confStruct interface{}) error {
	return fmt.Errorf("This file format is not supported yet")
}

func (w *defaultWorker) out(path string, confStruct interface{}) error {
	return fmt.Errorf("This file format is not supported yet")
}

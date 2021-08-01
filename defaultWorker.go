package conf

import "fmt"

type defaultWorker workerStruct

func newDefaultWorker() worker {
	return new(defaultWorker)
}

func (w *defaultWorker) load(path string, confStruct interface{}) error {
	return fmt.Errorf("this file format is not supported yet")
}

func (w *defaultWorker) out(path string, confStruct interface{}) error {
	return fmt.Errorf("this file format is not supported yet")
}

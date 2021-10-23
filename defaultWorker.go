package conf

import "fmt"

type defaultWorker struct {
	extenstion string
}

func newDefaultWorker(extension string) worker {
	return &defaultWorker{
		extenstion: extension,
	}
}

func (w *defaultWorker) load(path string, confStruct interface{}) error {
	return fmt.Errorf("%s format is not supported yet", w.extenstion)
}

func (w *defaultWorker) out(path string, confStruct interface{}) error {
	return fmt.Errorf("%s format is not supported yet", w.extenstion)
}

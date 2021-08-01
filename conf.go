package conf

func Load(fileFormat format, path string, confStruct interface{}) error {
	worker := workerByFormat(fileFormat)
	return worker.load(path, confStruct)
}

func Out(fileFormat format, path string, confStruct interface{}) error {
	worker := workerByFormat(fileFormat)
	return worker.out(path, confStruct)
}

func workerByFormat(fileFormat format) worker {
	switch fileFormat {
	case JSON:
		return newJsonWorker()
	default:
		return newDefaultWorker()
	}
}

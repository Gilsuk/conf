package conf

type worker interface {
	load(string, interface{}) error
	out(string, interface{}) error
}

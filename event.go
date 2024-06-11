package i18next

type Event struct {
	events map[string][]func(interface{})
}

func (e *Event) Add(key string, callback func(interface{})) error {

}

func (e *Event) Once(key string, callback func(interface{})) error {

}

func (e *Event) Emit(key string, payload ...interface{}) error {

}

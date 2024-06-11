package i18n

type EventBus struct {
	events map[string][]func(interface{})
}

func (e *EventBus) Add(key string, callback func(interface{})) error {

}

func (e *EventBus) Once(key string, callback func(interface{})) error {

}

func (e *EventBus) Emit(key string, payload ...interface{}) error {

}

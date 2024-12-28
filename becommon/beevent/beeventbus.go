package beevent

import (
	"reflect"
	"sync"
)

// BeEventBus manages event subscriptions and publishing.
type BeEventBus struct {
	subscribers map[reflect.Type][]reflect.Value
	mutex       sync.RWMutex
}

// NewEventBus creates a new EventBus instance.
func NewEventBus() *BeEventBus {
	return &BeEventBus{
		subscribers: make(map[reflect.Type][]reflect.Value),
	}
}

// Subscribe adds a new subscriber for a specific event type.
func (eb *BeEventBus) Subscribe(eventType interface{}, handler interface{}) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	evtType := reflect.TypeOf(eventType)
	handlerValue := reflect.ValueOf(handler)

	if handlerValue.Kind() != reflect.Func {
		panic("handler must be a function")
	}

	eb.subscribers[evtType] = append(eb.subscribers[evtType], handlerValue)
}

// Publish sends an event to all subscribers of the event type.
func (eb *BeEventBus) Publish(event interface{}) {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	evtType := reflect.TypeOf(event)

	if handlers, found := eb.subscribers[evtType]; found {
		for _, handler := range handlers {
			// Call the handler with the event
			handler.Call([]reflect.Value{reflect.ValueOf(event)})
		}
	}
}

package beevent

import (
	"errors"
	"reflect"
	"sync"
)

// BeEventBus manages event subscriptions and publishing.
type BeEventBus struct {
	subscribers map[reflect.Type][]reflect.Value
	mutex       sync.RWMutex
}

// NewEventBus creates a new EventBus instance.
func NewBeEventBus() *BeEventBus {
	return &BeEventBus{
		subscribers: make(map[reflect.Type][]reflect.Value),
	}
}

// Subscribe adds a new subscriber for a specific event type.
func (eb *BeEventBus) Subscribe(eventType interface{}, handler interface{}) error {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	evtType := reflect.TypeOf(eventType)
	handlerValue := reflect.ValueOf(handler)

	if handlerValue.Kind() != reflect.Func {
		return errors.New("handler must be a function")
	}

	// Check if handler has the correct signature
	if handlerValue.Type().NumIn() != 1 || handlerValue.Type().In(0) != evtType {
		return errors.New("handler must have a single parameter of the event type")
	}

	// Check if handler returns an error
	if handlerValue.Type().NumOut() != 1 || handlerValue.Type().Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return errors.New("handler must return a single error value")
	}

	eb.subscribers[evtType] = append(eb.subscribers[evtType], handlerValue)
	return nil
}

// Publish sends an event to all subscribers of the event type.
func (eb *BeEventBus) Publish(event interface{}) error {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	evtType := reflect.TypeOf(event)

	handlers, found := eb.subscribers[evtType]
	if !found {
		return errors.New("no subscribers for event type")
	}

	for _, handler := range handlers {
		// Call the handler with the event
		results := handler.Call([]reflect.Value{reflect.ValueOf(event)})
		if err := results[0].Interface(); err != nil {
			return err.(error)
		}
	}
	return nil
}

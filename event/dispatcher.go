package event

import (
	"reflect"
	"sort"
)

// DispatcherInterface interface
type DispatcherInterface interface {
	AddSubscriber(subscriber *SubscriberInterface)
	HasListener(eventName string) bool
	Dispatch(eventName string) error
}

// Dispatcher struct implements DispatcherInterface
type Dispatcher struct {
	Listeners map[string][]*Listener
}

// NewDispatcher create new Dispatcher Object
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Listeners: make(map[string][]*Listener),
	}
}

// AddSubscriber register subscriber to dispatcher
func (dp *Dispatcher) AddSubscriber(subscriber SubscriberInterface) {
	listeners := subscriber.GetSubscribedEvents()
	for _, listener := range listeners {
		eventName := listener.EventName

		en, ok := dp.Listeners[eventName]
		if !ok {
			en := make([]*Listener, 0)
			dp.Listeners[eventName] = en
		}
		dp.Listeners[eventName] = append(en, listener)
	}
}

// HasListener check if Dispatcher has listener
func (dp *Dispatcher) HasListener(eventName string) bool {
	_, ok := dp.Listeners[eventName]

	if !ok {
		return false
	}
	return true
}

// Dispatch eventName event
func (dp *Dispatcher) Dispatch(event Event) error {
	eventName := event.GetName()
	if !dp.HasListener(eventName) {
		return nil
	}

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(event)
	listeners := dp.Listeners[eventName]
	sort.Sort(ByPriority(listeners))

	for _, listener := range listeners {
		subscriber := listener.Subscriber
		methodName := listener.MethodName
		reflect.ValueOf(subscriber).MethodByName(methodName).Call(inputs)
	}

	return nil
}

// ByPriority sort listeners by priority
type ByPriority []*Listener

func (pr ByPriority) Len() int {
	return len(pr)
}

func (pr ByPriority) Swap(i, j int) {
	pr[i], pr[j] = pr[j], pr[i]
}

func (pr ByPriority) Less(i, j int) bool {
	return pr[j].Priority < pr[i].Priority
}

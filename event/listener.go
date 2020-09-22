package event

//SubscriberInterface interface
type SubscriberInterface interface {
	GetSubscribedEvents() []*Listener
}

// Listener interface
type Listener struct {
	Subscriber interface{}
	EventName  string
	MethodName string
	Priority   int
}

// NewListener creates new Listener
func NewListener(subscriber interface{}, eventName string, methodName string, priority int) *Listener {
	return &Listener{
		EventName:  eventName,
		MethodName: methodName,
		Priority:   priority,
		Subscriber: subscriber,
	}
}

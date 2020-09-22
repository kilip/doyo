package event

// M is short name fo map[string]...
type M map[string]interface{}

// Event interface
type Event interface {
	GetName() string
}

// BasicEvent a basic event struct
type BasicEvent struct {
	name string

	data map[string]interface{}
}

func (e *BasicEvent) initData() {
	e.data = make(map[string]interface{})
}

// NewBasicEvent creates a basic event instance
func NewBasicEvent(name string, data M) *BasicEvent {
	if data == nil {
		data = make(map[string]interface{})
	}
	return &BasicEvent{
		name: name,
		data: data,
	}
}

// Set event data value by key
func (e *BasicEvent) Set(key string, val interface{}) {
	if e.data == nil {
		e.initData()
	}
	e.data[key] = val
}

// Get event data value by key
func (e *BasicEvent) Get(key string) interface{} {
	if v, ok := e.data[key]; ok {
		return v
	}
	return nil
}

// GetName get event name
func (e BasicEvent) GetName() string {
	return e.name
}

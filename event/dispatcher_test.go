package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestSubscriber struct {
	Test1 bool
}

func (ts TestSubscriber) GetSubscribedEvents() []*Listener {
	listeners := make([]*Listener, 2)

	listeners[0] = NewListener(ts, "test", "OnCalled2", 10)
	listeners[1] = NewListener(ts, "test", "OnCalled1", 1)
	return listeners
}

func (ts TestSubscriber) OnCalled1(event *BasicEvent) {
	event.Set("called1", true)
}

func (ts TestSubscriber) OnCalled2(event *BasicEvent) {
	event.Set("called2", true)
}

func TestDispatcher(t *testing.T) {
	dp := NewDispatcher()
	ts := &TestSubscriber{}

	dp.AddSubscriber(ts)
	assert.True(t, dp.HasListener("test"))

	data := map[string]interface{}{
		"called1": false,
		"called2": false,
	}
	evt := NewBasicEvent("test", data)
	assert.Equal(t, false, evt.Get("called1"))
	assert.Equal(t, false, evt.Get("called2"))

	dp.Dispatch(evt)
	assert.Equal(t, true, evt.Get("called1"))
	assert.Equal(t, true, evt.Get("called2"))

}

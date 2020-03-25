package helpers

import (
	"reflect"
	"sync"
)

var eventManagerWithDataInstance *EventManagerWithData
var eventManagerWithDataSync sync.Once

// EventHandlerWithData - eventWithData handler prototype
type EventHandlerWithData func([]byte)

type eventWithData struct {
	subscribers  []EventHandlerWithData
	lastData     []byte
	happenedOnce bool
}

// EventManagerWithData - eventWithData manager
type EventManagerWithData struct {
	events     map[string]*eventWithData
	eventsSync sync.RWMutex
}

// EventsWithData - Singleton to get the events manager instance
func EventsWithData() *EventManagerWithData {
	eventManagerWithDataSync.Do(func() {
		eventManagerWithDataInstance = &EventManagerWithData{
			events:     make(map[string]*eventWithData),
			eventsSync: sync.RWMutex{},
		}
	})

	return eventManagerWithDataInstance
}

// On - Tells EventManagerWithData to add a subscriber for an eventWithData
func (thisRef *EventManagerWithData) On(eventName string, eventHandler EventHandlerWithData) {
	thisRef.addEventIfNotExists(eventName)
	thisRef.addSubscriberIfNotExists(eventName, eventHandler)

	// thisRef.eventsSync.RLock()
	// defer thisRef.eventsSync.RUnlock()

	// if thisRef.events[eventName].happenedOnce {
	// 	go eventHandler(thisRef.events[eventName].lastData)
	// }
}

// Off - Tells EventManagerWithData to remove a subscriber for an eventWithData
func (thisRef *EventManagerWithData) Off(eventName string, eventHandler EventHandlerWithData) {
	thisRef.addEventIfNotExists(eventName)

	thisRef.eventsSync.RLock()
	defer thisRef.eventsSync.RUnlock()

	var foundIndex = -1
	for index, existingEventHandler := range thisRef.events[eventName].subscribers {
		if reflect.ValueOf(eventHandler) == reflect.ValueOf(existingEventHandler) {
			foundIndex = index
			break
		}
	}

	if foundIndex != -1 {
		thisRef.events[eventName].subscribers = append(
			thisRef.events[eventName].subscribers[:foundIndex],
			thisRef.events[eventName].subscribers[foundIndex+1:]...,
		)
	}
}

// Raise - Informs all subscribers about the eventWithData
func (thisRef *EventManagerWithData) Raise(eventName string, data []byte) {
	thisRef.addEventIfNotExists(eventName)

	thisRef.eventsSync.RLock()
	defer thisRef.eventsSync.RUnlock()

	thisRef.events[eventName].happenedOnce = true
	thisRef.events[eventName].lastData = data
	for _, eventHandler := range thisRef.events[eventName].subscribers {
		go eventHandler(data)
	}
}

func (thisRef *EventManagerWithData) addEventIfNotExists(eventName string) {
	thisRef.eventsSync.Lock()
	defer thisRef.eventsSync.Unlock()

	if _, ok := thisRef.events[eventName]; !ok {
		thisRef.events[eventName] = &eventWithData{
			subscribers:  []EventHandlerWithData{},
			happenedOnce: false,
		}
	}
}

func (thisRef *EventManagerWithData) addSubscriberIfNotExists(eventName string, eventHandler EventHandlerWithData) bool {
	thisRef.eventsSync.RLock()
	defer thisRef.eventsSync.RUnlock()

	// 1. Check if delegate for the eventWithData already there, assumes map-key exists
	var alreadyThere = false

	for _, existingEventHandler := range thisRef.events[eventName].subscribers {
		alreadyThere = (reflect.ValueOf(eventHandler) == reflect.ValueOf(existingEventHandler))
		if alreadyThere {
			break
		}
	}

	// 2. Add the delegate
	if !alreadyThere {
		thisRef.events[eventName].subscribers = append(
			thisRef.events[eventName].subscribers,
			eventHandler,
		)
	}

	return alreadyThere
}

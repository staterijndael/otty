package handlers

import "fmt"

// Handler is a interface for handlers to describe otty
type Handler interface {
	GetName() string
	GetData() []byte
	SetName([]byte)
	SetData([]byte)
}

// InitHandlers initialize basic handlers
func InitHandlers() []Handler {
	routeHandler := &Route{Name: []byte("Route")}
	dataHandler := &Data{Name: []byte("Data")}

	var handlers []Handler

	handlers = append(handlers, routeHandler)
	handlers = append(handlers, dataHandler)

	return handlers
}

// FindHandlerByName returns handler which name is equal
func FindHandlerByName(handlers []Handler, name string) Handler {
	for _, handler := range handlers {
		if handler.GetName() == name {
			return handler
		}
	}
	return nil
}

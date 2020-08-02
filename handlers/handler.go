package handlers

// Handler is a interface for handlers to describe otty
type Handler interface {
	GetName() string
	GetData() []byte
	SetName([]byte)
	SetData([]byte)
}

// InitHandlers initialize basic handlers
func InitHandlers() map[string]Handler {
	routeHandler := &Route{Name: []byte("Route")}
	dataHandler := &Data{Name: []byte("Data")}

	var handlers map[string]Handler

	handlers[routeHandler.GetName()] = routeHandler
	handlers[dataHandler.GetName()] = dataHandler

	return handlers
}

// FindHandlerByName returns handler which name is equal
func FindHandlerByName(handlers map[string]Handler, name string) Handler {
	if handlers[name] != nil {
		return handlers[name]
	}

	return nil
}

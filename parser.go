package otty

import (
	"github.com/Oringik/otty/handlers"
)

// Otty ...
type Otty struct {
	RawData   []byte
	Handlers  map[string]handlers.Handler
	Endpoints map[string]func([]byte)
}

// New returns pointer to new otty structure
func New() *Otty {
	return &Otty{
		Handlers:  make(map[string]handlers.Handler),
		Endpoints: make(map[string]func([]byte)),
	}
}

// ParseOtty parsing any data and return structure with ready handlers and raw data
func (otty *Otty) ParseOtty(data []byte) {
	otty.Handlers = handlers.InitHandlers()
	otty.RawData = data

	for i := 0; i < len(data)-1; i++ {
		if isSpace(data[i]) {
			continue
		}

		if isEnd(data[i]) {
			data = []byte{}
			break
		}

		if isNewString(data[i]) {
			if len(data) > i+1 {
				data = otty.ParseHandler(data[i+1:])
				i = 0

				continue
			}
		}

		data = otty.ParseHandler(data[i:])
		i = 0

	}

}

// ParseHandler parsing handlers for otty structure
func (otty *Otty) ParseHandler(data []byte) []byte {
	handlersToSearch := otty.Handlers
	var name []byte
	var value []byte

	// Parse name of handler
	for i := 0; i < len(data)-1; i++ {
		if isSpace(data[i]) {
			continue
		}

		if isEnd(data[i]) {
			data = []byte{}
			break
		}

		if isColon(data[i]) {
			data = data[i+1:]
			break
		}

		name = append(name, data[i])

	}

	// Parse value of handler
	for i := 0; i < len(data)-1; i++ {
		if isSpace(data[i]) {
			continue
		}

		if isEnd(data[i]) {
			data = []byte{}
			break
		}

		if isNewString(data[i]) {
			data = data[i+1:]
			break
		}

		value = append(value, data[i])

	}

	handler := handlers.FindHandlerByName(handlersToSearch, string(name))

	// Set name and data for found handler
	handler.SetName(name)
	handler.SetValue(value)

	return data

}

func isColon(symbol byte) bool {
	if symbol == 58 {
		return true
	}

	return false
}

func isEnd(symbol byte) bool {
	if symbol == 0 {
		return true
	}

	return false
}

func isSpace(symbol byte) bool {
	if symbol == 32 || symbol == 9 {
		return true
	}

	return false
}

func isNewString(symbol byte) bool {
	if symbol == 10 {
		return true
	}

	return false
}

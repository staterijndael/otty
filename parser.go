package otty

import "otty/handlers"

// Otty ...
type Otty struct {
	RawData  []byte
	Handlers []handlers.Handler
}

// New returns pointer to new otty structure
func New() *Otty {
	return &Otty{}
}

// ParseOtty parsing any data and return structure with ready handlers and raw data
func ParseOtty(data []byte) *Otty {
	otty := New()
	otty.Handlers = handlers.InitHandlers()

	for i := 0; i < len(data)-1; i++ {
		if isSpace(data[i]) {
			i++
			continue
		}

		if isNewString(data[i]) {
			data = otty.ParseHandler(data[i+1:])
			i = 0
		}

	}

	return otty
}

// ParseHandler parsing handlers for otty structure
func (otty *Otty) ParseHandler(data []byte) []byte {
	handlersToSearch := otty.Handlers
	var name []byte
	var value []byte

	// Parse name of handler
	for i := 0; i < len(data)-1; i++ {
		if isSpace(data[i]) {
			i++
			continue
		}

		if isNewString(data[i]) {
			data = data[i+1:]
			break
		}

		name = append(name, data[i])

	}

	// Parse value of handler
	for i := 0; i < len(data)-1; i++ {
		if isSpace(data[i]) {
			i++
			continue
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
	handler.SetData(data)

	return data

}

func isSpace(symbol byte) bool {
	if symbol == 32 {
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

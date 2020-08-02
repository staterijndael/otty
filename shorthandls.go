package otty

import "github.com/Oringik/otty/handlers"

// Route returns route handler
func (otty *Otty) Route() *handlers.Route {
	return otty.Handlers["Route"].(*handlers.Route)
}

// Data returns data handler
func (otty *Otty) Data() *handlers.Data {
	return otty.Handlers["Data"].(*handlers.Data)
}

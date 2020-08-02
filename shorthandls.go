package otty

import "github.com/Oringik/otty/handlers"

// GetRoute returns route handler
func (otty *Otty) GetRoute() *handlers.Route {
	return otty.Handlers["Route"].(*handlers.Route)
}

// GetData returns data handler
func (otty *Otty) GetData() *handlers.Data {
	return otty.Handlers["Data"].(*handlers.Data)
}

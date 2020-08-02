package handlers

// Route is a structure for handler which contains routing path
type Route struct {
	Name []byte
	Data []byte
}

// GetName returns name of route handler
func (route *Route) GetName() string {
	return string(route.Name)
}

// GetData returns bytes array of data for this handler
func (route *Route) GetData() []byte {
	return route.Data
}

// SetName set name for route
func (route *Route) SetName(name []byte) {
	route.Name = name
}

// SetData set data for route
func (route *Route) SetData(value []byte) {
	route.Data = value
}

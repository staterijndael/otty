package handlers

// Data is a structure for handler which contains data
type Data struct {
	Name []byte
	Data []byte
}

// GetName returns name of data handler
func (data *Data) GetName() string {
	return string(data.Name)
}

// GetValue returns bytes array of data for this handler
func (data *Data) GetValue() []byte {
	return data.Data
}

// SetName set name for data
func (data *Data) SetName(name []byte) {
	data.Name = name
}

// SetValue set data for data
func (data *Data) SetValue(value []byte) {
	data.Data = value
}

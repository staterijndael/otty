package otty

import "errors"

// CreateEndpoint creates endpoint for call function by this endpoint after
func (otty *Otty) CreateEndpoint(endpoint string, function func([]byte)) error {
	if otty.Endpoints[endpoint] != nil {
		return errors.New("Endpoint already exist")
	}

	otty.Endpoints[endpoint] = function

	return nil
}

// ResolveEndpoint call function of endpoint by endpoint name
func (otty *Otty) ResolveEndpoint(endpoint []byte, data []byte) {
	endpointString := string(endpoint)

	functionToCall := otty.Endpoints[endpointString]
	if functionToCall == nil {
		return
	}

	functionToCall(data)
	return

}

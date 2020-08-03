package otty

// CreateEndpoint creates endpoint for call function by this endpoint after
func (otty *Otty) CreateEndpoint(endpoint string, function func(...interface{}) interface{}) {
	otty.Endpoint[endpoint] = function
}

// ResolveEndpoint call function of endpoint by endpoint name
func (otty *Otty) ResolveEndpoint(endpoint []byte, args ...interface{}) interface{} {
	endpointString := string(endpoint)

	functionToCall := otty.Endpoint[endpointString]

	returnValue := functionToCall(args)
	return returnValue

}

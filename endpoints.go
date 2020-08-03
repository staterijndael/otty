package otty

// CreateEndpoint creates endpoint for call function by this endpoint after
func (otty *Otty) CreateEndpoint(endpoint string, function func(...interface{}) interface{}) {
	otty.Endpoints[endpoint] = function
}

// ResolveEndpoint call function of endpoint by endpoint name
func (otty *Otty) ResolveEndpoint(endpoint []byte, args ...interface{}) interface{} {
	endpointString := string(endpoint)

	functionToCall := otty.Endpoints[endpointString]

	returnValue := functionToCall(args)
	return returnValue

}

// BasicIterator returns basic iterator
func BasicIterator(a ...interface{}) []interface{} {
	return a[0].([]interface{})
}

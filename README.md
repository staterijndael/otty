# Otty
Transfer protocol with routing
## Basic usage
```go
	data := []byte(
		`Data:{"someKey":"someValue"}
	Route:auth
	`)

	ottyStruct := otty.New()

	ottyStruct.ParseOtty(data)

	ottyStruct.CreateEndpoint("auth", func(a []byte) interface{} {
		fmt.Println(string(a))
		return nil

	})

	ottyStruct.ResolveEndpoint(ottyStruct.Route().GetValue(), ottyStruct.Data().GetValue())
  ```
In this example we are parsing our data. As you can see we are creating endpoint with name 'auth' and function that accepts byte array, after we are resolving endpoint by route name in data that we parsed and call function resolved in function CreateEndpoint which prints accepted byte array.

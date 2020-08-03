# Otty
Transfer protocol with routing
## Basic usage
```go
	data := []byte(
		`Data:{"someKey":"someValue"}
	Route:auth
	`)

	ottyStruct := otty.ParseOtty(data)

	ottyStruct.CreateEndpoint("auth", func(a ...interface{}) interface{} {
		for i := 0; i < len(a[0].([]interface{})); i++ {
			fmt.Println(a[0].([]interface{})[i])
		}
		return nil

	})

	ottyStruct.ResolveEndpoint(ottyStruct.Route().GetValue(), "value1", "value2", "value3")
  ```
In this example we are parsing our data. As you can see we are creating endpoint with name 'auth' and function that accepts any arguments, after we are resolving endpoint by route name in data that we parsed and call function resolved in function CreateEndpoint which prints all accepted arguments.

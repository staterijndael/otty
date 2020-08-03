# Otty
Transfer protocol with routing
## Basic usage
```
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
  In this example we just parsing our data, after we creating endpoint with name 'auth' and function that accept any arguments, after we resolving endpoint by route name in data we parsed and call function resolved in function CreateEndpoint that print all accepted arguments

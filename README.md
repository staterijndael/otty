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
		fmt.Println(a[0].([]interface{})[1])
		return nil

	})

	ottyStruct.ResolveEndpoint(ottyStruct.Route().GetValue(), "asdsas", "aa")
  ```
  In this example we just parsing our data, after we creating endpoint with name 'auth' and function that accept any arguments, after we resolving endpoint by route name in data we parsed and call function resolved in function CreateEndpoint

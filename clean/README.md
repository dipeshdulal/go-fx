#### Clean Architecture
 
```
.
├── bundler
├── controllers
├── main.go
├── routes
└── services
```

- Bundler combines all the modules together.
- Routes responsible for registering routes.
- Services responsible for creating services.
- Controllers responsible for creating controllers. Services are injected to Controllers.

##### To run `go run main.go` and go to `http://localhost:5000/user`
You can see that controllers are running properly and services injected are also being called properly
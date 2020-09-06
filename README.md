#### Go Dependency Injection

**WHY**: Let's not make injecting dependency not complicated for large application.

Learn about go lang dependency injection using `go-fx` DI framework.

- [initfx](/initfx) contains simple logger interface `MyAwesomeLogger` and a string `CustomString` injected in a function.

##### Checklist

- [x] First Program written in `go-fx`
- [x] Sample `Go-Gin` skeleton with dependency injection. [server](/server)
- [x] Build a sample application pseudo following clean architecture to check fx library. [clean](/clean)

##### Following [this](https://pmihaylov.com/shared-components-go-microservices/) tutorial.

- Every module dependency can be extracted separately by the use of `fx.Options` method.
- Service routes can be handled separately in the router.
- If we would like to follow clean architecture. We would simply want to provide it to the dependency tree. (Need to check by doing a sample application).
- Every module exports a `Module` variable when could then be added to `fx.Options` or `fx.New` method that would add required dependency to the tree.
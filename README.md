[Tutorial](https://www.howtographql.com/graphql-go/5-create-and-retrieve-links/)

[Tutorial](https://medium.com/@wahyubagus1910/implementation-graphql-server-with-golang-fb8f8303b4bc)

[gqlgn](https://gqlgen.com/config/)

```sh
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
Exec "go run ./server.go" to start GraphQL server
```

```sh
replace type in graph/schema.graphqls
```

```sh
created tools/tools.go
go mody tidy
```
```sh
go run github.com/99designs/gqlgen generate
```
It throws an error as old code does not get deleted in schema.resolvers.go for warning purpose. delete maually

```go
// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

```

the resolver.go file is like the context in apollo server, it maintains state accross request, it is here, dependencies such as db, pubsub, application service are passed into the graphql

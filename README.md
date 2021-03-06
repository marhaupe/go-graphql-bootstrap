# go-graphql-bootstrap

This is a very basic bootstrap for a GraphQL server powered by [gqlgen](https://github.com/99designs/gqlgen). The only thing I added is logging. Per default, every query that is not an introspect query, including its variables and its result, will be logged.

## How to use this bootstrap

Clone the repository and delete its .git folder:

```bash
git clone --depth=1 --branch=master https://github.com/marhaupe/go-graphql-bootstrap &&
rm -rf go-graphql-bootstrap/.git &&
cd go-graphql-bootstrap
```

Rename every import and usage of the repository's original name (you must have make, xargs and perl installed):

```bash
make rename domain=github.com user=peter repository=swapi-graphql
```

Download all dependencies (you must have [dep](https://github.com/golang/dep) installed): 
```bash
dep ensure
````

Instead of dep, you could just call `go get`. Howevery, I can't guarantee that this project will support the gqlgen version downloaded by `go get`, so be warned. 

Now you're set to go! Edit the schema in internal/schema.graphql and run `make regen` to run gqlgen. If you have and issues or improvements, feel free to open issues or submit a PR.

## Included scripts

- `make dev`: Runs main.go. Will be convenient e.g. if you have to set up envs or start up a database.
- `make build [output]`: Builds project to the optionally given output parameter
- `make regen`: Runs gqlgen
- `make test`: Runs every test
- `make coverage`: Runs every test and generates the test coverage to cover.html
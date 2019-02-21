# go-graphql-bootstrap

This is a very basic bootstrap for a GraphQL server powered by [gqlgen](https://github.com/99designs/gqlgen)

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

Instead of dep, you could just call `go get`. Since 99designs might introduce breaking changes in the way gqlgen works, your project might get a bit messed up.

Now you're set to go! If you have and issues or improvements, feel free to open issues or submit a PR.
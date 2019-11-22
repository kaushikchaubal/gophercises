## Context

This is the solution for this exercise: https://github.com/gophercises/urlshort

## How to run this?

* This uses the default values

```
go run main.go
```

## How to test this

* Visiting http://localhost:8080/ should print:

```
This is the default path... be more creative!
```

* Visiting http://localhost:8080/u should redirect to `https://www.google.com"`
* Visiting http://localhost:8080/y should redirect to `https://godoc.org/gopkg.in/yaml.v2"`
* Visiting http://localhost:8080/a should redirect to `https://github.com/gophercises/urlshort`
* Visiting http://localhost:8080/b should redirect to `https://github.com/kaushikchaubal/gophercises`
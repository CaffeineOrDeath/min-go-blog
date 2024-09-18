# Minimum Go Blog
(Series)[https://dev.to/caffeineordeath/from-r-to-go-4pk8] write-up

This is the bare minimum to get a minimum Go/Templ website started. By the time this is
done, it is intended to be a complete starting point for Golang web
applications.

## Packages

(Go)[https://go.dev/]

There are a few needed packages, regardless of editor.
While this repo can `go run` as it is, any modifications you make, I would
highly recommend installing `gopls`.

Gopls is the language server for Golang. 
`go install golang.org/x/tools/gopls@latest`

(Templ)[https://github.com/a-h/templ]
`go get github.com/a-h/templ/cmd/templ@latest`

Templ is our template manager. You write `*.templ` files and it generates
`*_templ.go` files.

## Future additions

(HTMX)[https://htmx.org/]
`<script src="https://unpkg.com/htmx.org@2.0.2"></script>a`

HTMX allows us to stay with a single compiled language, removing the need to
code in JS. You can if you're more comfortable with that.

(Docker)[https://www.docker.com]

Docker allows us to make clones of our application.

(Kubernetes)[https://kubernetes.io]
k8 is the container manager. This allows instances of our application to be
created when new connections are made, and also keeps them separate. A web based
sandbox generator.

# GOBURL - Simple URL obsfucator microservice written in Go

Basically a more evolved "Hello world" in Golang...

# Usage

`goburl` starts service.

Request like `GET http://localhost:8080/obsfucate?url=https://golang.org/pkg/encoding/base64/` returns body like:

`http://localhost:8080/?u=aHR0cHM6Ly9nb2xhbmcub3JnL3BrZy9lbmNvZGluZy9iYXNlNjQv`

...and when this URL is accessed, a temporary redirect reponse to the original URL is returned.

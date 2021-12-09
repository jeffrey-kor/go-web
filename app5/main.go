package main

import (
	"go-web/app5/handler"
	"net/http"
)

func main() {
	http.ListenAndServe(": 3000", handler.NewHandler())
}

package main

import (
	"fmt"
	"go-web/myapp"
	"net/http"
)


func main() {
	fmt.Println("Server is running on 3000...")
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}

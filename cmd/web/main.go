package main

import (
	"fmt"
	"net/http"

	"github.com/ssr0016/booking/cmd/pkg/handlers"
)

const portNumber = ":4000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

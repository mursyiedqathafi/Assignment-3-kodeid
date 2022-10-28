package main

import (
	"assignment-3/service"
	"fmt"
	"net/http"
)

func main() {
	port := ":8082"
	go service.AutoReload()
	http.HandleFunc("/", service.ReloadWeb)
	fmt.Println("This Auto-Reload App is listening on port", port)
	http.ListenAndServe(port, nil)
}

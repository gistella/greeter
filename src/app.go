package main

import (
	"fmt"
	"net/http"
)

func main() {

	srvA := http.NewServeMux()
	srvA.HandleFunc("/", Greet)

	srvB := http.NewServeMux()
	srvB.HandleFunc("/", Greet)

	go http.ListenAndServe(":12027", srvA)
	http.ListenAndServe(":12028", srvB)

}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s\n", r.RemoteAddr)
}

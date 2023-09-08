package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func Run(addr string) {
	http.HandleFunc("/", helloworld)
	fmt.Println("Server started and listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("hello, world server\n")
	server1()
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!server1\n")
}

func server1() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Fatal("ListenAndServe: test", err)
}

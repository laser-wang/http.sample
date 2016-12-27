package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("hello, world client\n")
	client1()
}

func client1() {
	resp, err := http.Get("http://192.168.0.10:12345/hello")
	//resp, err := http.Get("http://127.0.0.1:12345/hello")
	if err != nil {
		log.Fatal("client1:", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}

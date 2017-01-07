package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"strings"
)

func main() {
	fmt.Printf("hello, world client\n")
	client2()
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

func client2() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://192.168.0.10:12345/login?userCode=001&pwd=001", nil)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	req.Header.Set("loginToken", "xxx122311")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

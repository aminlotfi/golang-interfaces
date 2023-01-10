package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://ticateb.com")
	if err != nil {
		panic(err)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

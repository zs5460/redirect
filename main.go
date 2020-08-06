package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/caddyserver/certmagic"
	"github.com/zs5460/my"
)

type site struct {
	Domain string `json:"domain"`
	URL    string `json:"url"`
}

type config struct {
	Sites []site `json:"sites"`
}

func main() {
	cfg := config{}
	my.MustLoadConfig("config.json", &cfg)
	if len(cfg.Sites) == 0 {
		fmt.Println("sites is empty!")
		os.Exit(0)
	}

	mux := NewHandler(cfg.Sites)

	domains := []string{}
	for _, v := range cfg.Sites {
		domains = append(domains, v.Domain)
	}

	err := certmagic.HTTPS(domains, mux)
	if err != nil {
		fmt.Println(err)
	}
}

// NewHandler return a handler
func NewHandler(s []site) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			host := r.Host
			for _, v := range s {
				if host == v.Domain {
					http.Redirect(w, r, v.URL, 301)
				}
			}
			w.Write([]byte("welcome!"))
		},
	)
}

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/ore/cash_redis"
)

func handler(c *cash_redis.Cach) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var target string
		target, _ = c.GetCash(r.Host)
		// if r.Host == "web.test.user.localhost:8090" {
		// 	target = "http://localhost:8081/"
		// } else if r.Host == "test2.user.com.localhost:8090" {
		// 	target = "http://localhost:8082/"
		// } else {
		// 	target = "http://httpbin.org/"
		// }
		if target == "" {
			c.SetCach(r.Host, "http://httpbin.org/")
			target = "http://httpbin.org/"
		}
		log.Println("[request]: " + r.Host)
		log.Println("[target]: " + target)

		remote, _ := url.Parse(target)
		w.Header().Set("X-Forwarded-For", r.Host)
		p := httputil.NewSingleHostReverseProxy(remote)
		p.ServeHTTP(w, r)
	}
}

func main() {
	c := cash_redis.Init("localhost:6379", "tcp")
	c.Connect()
	http.HandleFunc("/", handler(c))
	http.ListenAndServe(":8090", nil)
}

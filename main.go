package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type request struct {
	Timestamp  string      `json:"ts"`
	ClientIP   string      `json:"clientip"`
	Method     string      `json:"method"`
	Host       string      `json:"host"`
	URL        string      `json:"url"`
	Headers    http.Header `json:"headers"`
	FormValues url.Values  `json:"formvalues"`
}

func logJSON(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		addr := r.RemoteAddr
		lastColon := strings.LastIndex(addr, ":")
		ip := addr[0:lastColon]
		r.ParseForm()
		req := &request{
			Timestamp:  now.String(),
			ClientIP:   ip,
			Method:     r.Method,
			Host:       r.Host,
			URL:        r.URL.String(),
			Headers:    r.Header,
			FormValues: r.PostForm,
		}
		reqJSON, err := json.Marshal(req)
		if err != nil {
			log.Print(err)
		} else {
			fmt.Println(string(reqJSON))
		}
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	err := http.ListenAndServe(":9090", logJSON(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

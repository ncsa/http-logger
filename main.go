package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type context map[string]interface{}

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

type blockedHandler struct {
	templates *template.Template
}

func newBlockedHandler(templatePath string) *blockedHandler {
	indexTemplate := path.Join(templatePath, "index.html")
	templates := template.Must(template.ParseFiles(indexTemplate))
	return &blockedHandler{
		templates: templates,
	}

}

func (bh *blockedHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := &context{
		"Host": req.Host,
	}
	err := bh.templates.ExecuteTemplate(w, "index.html", ctx)
	if err != nil {
		e := fmt.Sprintf("Error rendering template: %s", err.Error())
		http.Error(w, e, http.StatusInternalServerError)
		return
	}
}

func main() {
	templatePath := flag.String("template", "./", "Path to templates")
	flag.Parse()
	http.Handle("/", newBlockedHandler(*templatePath))
	err := http.ListenAndServe(":9090", logJSON(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

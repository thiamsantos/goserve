package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Printf("%s %s\n", r.Method, r.URL)
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	var port int
	var path string

	flag.IntVar(&port, "port", 8080, "Port to use")
	flag.StringVar(&path, "path", ".", "Path to serve")
	flag.Usage = func() {
		fmt.Printf("Usage\n")
		fmt.Printf("  $ goserve [options]\n")
		fmt.Printf("\n")
		fmt.Printf("Options\n")
		flag.PrintDefaults()
		fmt.Printf("\n")
		fmt.Printf("Examples\n")
		fmt.Printf("  $ goserve\n")
		fmt.Printf("  $ goserve -port 3000\n")
		fmt.Printf("  $ goserve -path /tmp/static\n")
		fmt.Printf("  $ goserve -port 8888 -path /tmp/static\n")
	}

	flag.Parse()

	fs := http.FileServer(http.Dir(path))
	http.Handle("/", withLogging(fs))
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Serving %s at http://localhost:%d\n", path, port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

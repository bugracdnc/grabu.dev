package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func servePdf(w http.ResponseWriter, r *http.Request) {
	locale := r.URL.Query().Get("locale")
	cvEn := "static/pdf/cv-grabu.pdf"
	cvTr := "static/pdf/cv-grabu-tr.pdf"
	w.Header().Set("Content-Type", "application/pdf")
	if strings.Contains(strings.ToLower(locale), "tr") {
		http.ServeFile(w, r, cvTr)
	}
	http.ServeFile(w, r, cvEn)
}

func main() {
	port := os.Args[1:]

	if len(port) == 0 || port == nil {
		fmt.Printf("usage: CV :[port]\n\ne.g.\tCV :80, CV :443\n")
		return
	}

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/cv", servePdf)

	fmt.Printf("Starting server on %s...\n\n", port[0])
	err := http.ListenAndServe(port[0], nil)
	if err != nil {
		log.Panic(err)
	}
}

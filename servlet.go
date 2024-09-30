package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
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
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/cv", servePdf)

	fmt.Println("Starting server on :443...")
	err := http.ListenAndServe(":443", nil)
	if err != nil {
		log.Panic(err)
	}
}

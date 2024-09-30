package main

import (
	"fmt"
	"net/http"
	"strings"
)

func servePdf(w http.ResponseWriter, r *http.Request) {
	locale := r.URL.Query().Get("locale")
	cvEn := "static/pdf/cv-grabu-dev-en.pdf"
	cvTr := "static/pdf/cv-grabu-dev-tr.pdf"
	w.Header().Set("Content-Type", "application/pdf")
	if strings.Contains(strings.ToLower(locale), "tr") {
		http.ServeFile(w, r, cvTr)
	}
	http.ServeFile(w, r, cvEn)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/pdf", servePdf)

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

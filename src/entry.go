package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" {
		w.Write([]byte("OK by FinalEndPoint"))
		return
	}

	if r.Method != "GET" {
		http.Error(w, "503 Bad Request - Not Allowed Method", http.StatusServiceUnavailable)
		return
	}

	if strings.HasPrefix(r.RequestURI, "/w2766/") {
		http.Error(w, "402 - You NEED pay!", http.StatusPaymentRequired)
		return
	}
	
	if strings.HasPrefix(r.RequestURI, "/renbaoshuo/") {
		http.Error(w, "402 - FUCK YOU，You can only do DDOS？idiot！", http.StatusPaymentRequired)
		return
	}

	if r.RequestURI == "/robots.txt" {
		w.Write(robotsTxtByte)
		return
	}

	if isLog {
		log.Println(r.Host + r.RequestURI)
	}

	url, ok := convFastGitToCloudFlareWorker(r.Host, r.RequestURI)
	if ok {
		http.Redirect(w, r, url, 301)
	} else {
		http.Error(w, "503 Bad Request - Not Allowed Url", http.StatusServiceUnavailable)
	}
	return
}

func main() {
	initialise()

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at %v...\n", listenAddr)
	fmt.Printf("Workers Link: %v\n", workerUrl)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}

func initialise() {
	parseArgs()

	if confPath != "" {
		m, ok:= parseFromFile(confPath)
		if ok {
			convHostMap = m
		}
	}
}


package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "17-web-server/index.html")
}

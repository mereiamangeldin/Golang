package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe("8181", nil)
	if err != nil {
		panic(err)
	}

}

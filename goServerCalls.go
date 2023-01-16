package main

import (
	"fmt"
	"log"
	"net/http"
)

func Teja(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: red'>hello.....teja!!</h1>"))
}

func Puneeth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: blue'>hello.....puneeth!!</h1>"))
}

func Random(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: green'>unkown user</h1>"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: red'>hello.....!!</h1>"))
}

func List(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<h1> available list: </h1><h3 style='color: green' ><ol> <li> teja </li><li> puneeth </li><li> hyak </li></ol></h3>"))

		return
	}

	if ok && len(keys[0]) >= 1 {
		key := keys[0]

		w.Header().Set("Content-Type", "text/html")
		if key == "teja" {
			fmt.Fprintf(w, "<h1 style='color: red'>hello.....%s!!</h1>", key)
		} else if key == "puneeth" {
			fmt.Fprintf(w, "<h1 style='color: blue'>hello.....%s!!</h1>", key)
		} else if key == "hyak" {
			fmt.Fprintf(w, "<h1 style='color: yellow'>hello.....%s!!</h1>", key)
		} else {
			fmt.Fprintf(w, "<h1 style='color: green'>hello.....%s!!</h1>", key)
		}

	}

}
func Main() {

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/names/list", List)
	http.HandleFunc("/puneeth", Puneeth)
	http.HandleFunc("/random", Random)
	http.HandleFunc("/teja", Teja)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

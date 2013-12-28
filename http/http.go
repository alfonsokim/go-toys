package main

import (
    "fmt"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func main() {
    http.Handle("/string", String("Joe Merol!"))
    http.Handle("/struct", &Struct{"Hello", ":", "Joe!!"})
    http.ListenAndServe("localhost:4000", nil)
}


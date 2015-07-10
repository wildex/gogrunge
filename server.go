package main

import (
    "net"
    "net/http/fcgi"
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    //"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
)

type Msg struct {
    Id int
    Name string
    Message string
}

func main() {
    listener, _ := net.Listen("tcp", "127.0.0.1:9001")

    rtr := mux.NewRouter()
    rtr.HandleFunc("/", indexHandler)
    rtr.HandleFunc("/test/", testHandler)

    fcgi.Serve(listener, rtr)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

    t := template.New("index.html") //create a new template
    t, _ = t.ParseFiles("templates/index.html") //open and parse a template text file

    messages := []Msg{Msg{1, "test", "test12"}, Msg{2, "test2", "test23"}}
    t.Execute(w, messages)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

    t := template.New("index.html") //create a new template
    t, _ = t.ParseFiles("templates/index.html") //open and parse a template text file

    messages := []Msg{Msg{1, "dasdas", "testdasdsa12"}, Msg{2, "1123", "dsads"}}
    t.Execute(w, messages)
}
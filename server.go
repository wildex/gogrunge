package main

import (
    "net"
    "net/http"
    "net/http/fcgi"
    "html/template"
)

type FastCGIServer struct{}

type Msg struct {
    Id int
    Name string
    Message string
}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    t := template.New("index.html") //create a new template
    t, _ = t.ParseFiles("templates/index.html") //open and parse a template text file

    messages := []Msg{Msg{1, "test", "test12"}, Msg{2, "test2", "test23"}}
    t.Execute(resp, messages)
}

func main() {
    listener, _ := net.Listen("tcp", "127.0.0.1:9001")
    srv := new(FastCGIServer)
    fcgi.Serve(listener, srv)
}

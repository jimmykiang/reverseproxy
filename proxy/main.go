package main

import (
	"fmt"
	Router "github.com/jimmykiang/reverseproxy/proxy/router"
	"net/http"
	"net/http/httputil"
	"os"
	d "github.com/jimmykiang/reverseproxy/proxy/director"
)

var rp = httputil.ReverseProxy{
	Director: d.NewDirector(),
}

func main(){

	fmt.Printf("SCHEME %v\n", os.Getenv("SCHEME"))
	fmt.Printf("HOST %v\n", os.Getenv("HOST"))
	fmt.Printf("AUTH %v\n", os.Getenv("AUTH"))
	fmt.Println("Starting proxy server")	
	
	//r := rt.NewRouter(&rp)

	srv := http.Server{
		//Handler: r,
		Handler: Router.Routes(),
		Addr: "localhost:8081",
	}
	srv.ListenAndServe()
}
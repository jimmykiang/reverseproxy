package director

import (
	"fmt"
	"net/http"
	"os"
)

func NewDirector() func(req *http.Request){
	host:= os.Getenv("HOST")
	scheme:= os.Getenv("SCHEME")
	return func(req *http.Request){
		req.URL.Host = host
		req.URL.Scheme = scheme

		// toDo: modify incoming request payLoad.
		fmt.Printf("SECRET %v\n", req.Header.Get("SECRET"))
	}
}
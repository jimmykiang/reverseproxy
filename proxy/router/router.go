package Router

import (
	d "github.com/jimmykiang/reverseproxy/proxy/director"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	return mux
}
func home(w http.ResponseWriter, r *http.Request) {
	var rp = httputil.ReverseProxy{
		Director: d.NewDirector(),
	}
	//Getting envs
	secret := os.Getenv("SECRET")
	hardCodedAuth := os.Getenv("AUTH")

	if secret!= "" {
		r.Header.Add("SECRET", secret)
	}

	auth := r.Header.Get("AUTH")

	if auth != hardCodedAuth {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "Auth needed\n")
		return
	}

	rp.ServeHTTP(w, r)
}
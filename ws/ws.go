package main

import (
	"fmt"
	"login"
	"net/http"
	"strings"
	"umt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Server", "Wayneserver (Linux/CentOS)")
	host := r.Host
	userid := login.IdentifyLoggedInUser(w, r)
	operation := r.URL.Path
	fmt.Println(userid, host, operation)
	if strings.Contains(operation, "..") {
		w.WriteHeader(404)
		return
	}
	if len(operation) >= 7 {
		if operation[:7] == "/login/" {
			login.Handler(w, r, operation[7:], userid)
			return
		}
	}
	if len(operation) >= 5 {
		if operation[:5] == "/umt/" {
			umt.Handler(w, r, operation[5:], userid)
			return
		}
	}
	if host == "ultimatemusictoy.org" {
		umt.Handler(w, r, operation[1:], userid)
		return
	}
	// This is an abbreviated version of this web server, so everything else
	// I run on my actual site has been taken out
	// (Hopefully it still builds)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":4000", nil)
	fmt.Println(err)
}

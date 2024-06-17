package profibook

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"testing"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func TestWebProfile(t *testing.T) {
	PORT := ":8003"
	r := http.NewServeMux()
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", myHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// $ go tool pprof http://localhost:8003/debug/pprof/profile
// (pprof) top
// https://www.apachelounge.com/download/
// Download httpd-2.4.59-240605-win64-VS17.zip
// $ C:\Prog\Apache24\bin\ab -k -c 10 -n 100000 "http://127.0.0.1:8003/time"
// http://127.0.0.1:8003/debug/pprof/allocs?debug=1

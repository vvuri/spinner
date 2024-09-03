package websocket

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"testing"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	listenAddr string = "127.0.0.1:8080"
	wsAddr     string = "127.0.0.1:8080"
	jsTemplate *template.Template
)

func init() {
	//flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	//flag.StringVar(&wsAddr, "ws-addr", "", "Address for WebSocket connection")
	//flag.Parse()

	var err error
	jsTemplate, err = template.ParseFiles("C:/code/angular/spinner/cmd/blackhatbook/websocket/logger.js")

	if err != nil {
		panic(err)
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}

func TestMainServer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveWS)
	r.HandleFunc("/k.js", serveFile)

	t.Fatal(http.ListenAndServe(":8080", r))

}

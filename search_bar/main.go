package main

import (
	websocketfunc "app/ds/websocketFunc"
	ws "app/ds/websocketPkg"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	wsConn := ws.NewHub()

	http.HandleFunc("/websocket_test/", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the value of the dynamic field from the request URL
		// NOTED THAT websocket link start with wss://example.com/websocket/
		dynamic := strings.TrimPrefix(r.URL.Path, "/websocket_test/")
		path := strings.Split(dynamic, "/")
		if len(path) < 2 {
			return
		}
		fmt.Println("Connection:::: ", r)
		switch path[0] {
		case "search_bar":
			//20230419 Token Refresh -------
			fmt.Println("			*****", r)
			wsConn.Handler(w, r, websocketfunc.SearchBar)
		}
	})

	go wsConn.WaitingSocketConnection()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

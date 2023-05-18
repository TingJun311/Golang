package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//"time"
	websocketfunc "app/ds/websocketFunc"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HomePage(w http.ResponseWriter, r *http.Request) {

}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var data2 string
		data := string(p)
		if err := json.Unmarshal([]byte(data), &data2); err != nil {
			log.Println(err)
		}
		log.Println("[MESSAGE]: ", data2)
		res := websocketfunc.SearchBar(data)

		if err := conn.WriteMessage(messageType, res); err != nil {
			log.Println(err)
			return
		}
	}
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected")

	reader(ws)
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/websocket_test/search_bar", WebSocket)
	//websocketFunc()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func websocketFunc() {
	// Set the HTTP request header with the Bearer token
	header := http.Header{}

	// Create a new WebSocket dialer with the HTTP request header
	//	dialer := websocket.Dialer{
	//		Proxy:           http.ProxyFromEnvironment,
	//		HandshakeTimeout: 5 * time.Second,
	//	}

	// Dial the WebSocket server with the specified URL and protocols
	conn, _, err := websocket.DefaultDialer.Dial("wss://app-dev.s-eco.com.my/websocket_test/search_bar/", header)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send a message to the WebSocket
	request := make(map[string]interface{})
	request["customer_mapping_id"] = "3218350"

	reqJson, err := json.Marshal("T")
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {

			err = conn.WriteMessage(websocket.TextMessage, reqJson)
			if err != nil {
				log.Fatal("Failed to send message to WebSocket:", err)
			}
			fmt.Println("Message sent successfully")
			time.Sleep(5 * time.Second)
		}
	}()

	// Read messages from the WebSocket server
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		fmt.Printf("received: %s\n", message)
	}
}

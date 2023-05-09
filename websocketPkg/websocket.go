package websocket

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

/*
    List of websocket errors code

    Web sockets support status codes ranging from 1000 to 4999.

    1XXX — actively used range.
    2XXX — reserved for web socket extensions
    3XXX — registered first come first serve at IANA
    4XXX — available for applications

    1000 — CLOSE NORMAL — normal socket shut down
    1001 — CLOSE_GOING_AWAY — browser tab closing
    1002 — CLOSE_PROTOCOL_ERROR — endpoint received a malformed frame
    1003 — CLOSE_UNSUPPORTED — endpoint received unsupported frame
    1004 — No Code Name — reserved
    1005 — CLOSED_NO_STATUS — expected close status, received none
    1006 — CLOSE_ABNORMAL — no close code frame has been received
    1007 — Unsupported payload — endpoint received an inconsistent message
    1008 — Policy violation — generic code used for situations other than 1003 and 1009
    1009 — CLOSE_TOO_LARGE — endpoint won’t process large frame
    1010 — Mandatory extension — client wanted an extension which server did not negotiate
    1011 — Server error — internal server error while operating
    1012 — Service restart — service is restarting
    1013 — Try again later — temporary server condition forced blocking client’s request
    1014 — Bad gateway — server acting as gateway received an invalid response
    1015 — TLS handshake fail — transport Layer Security handshake failure

    const (
        CloseNormalClosure           = 1000
        CloseGoingAway               = 1001
        CloseProtocolError           = 1002
        CloseUnsupportedData         = 1003
        CloseNoStatusReceived        = 1005
        CloseAbnormalClosure         = 1006
        CloseInvalidFramePayloadData = 1007
        ClosePolicyViolation         = 1008
        CloseMessageTooBig           = 1009
        CloseMandatoryExtension      = 1010
        CloseInternalServerErr       = 1011
        CloseServiceRestart          = 1012
        CloseTryAgainLater           = 1013
        CloseTLSHandshake            = 1015
    )
*/


var Hubs = make(map[string]*Hub)

type Hub struct {
    clients   map[string]*websocket.Conn
    Broadcast chan Message
    upgrader  websocket.Upgrader
}

type Message struct {
    Data     []byte
    SenderID string
	Callback *func(string) []byte
}

func NewHub() *Hub {
	
    newHub := &Hub{
        clients:   make(map[string]*websocket.Conn),
        Broadcast: make(chan Message),
        upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
            },
        },
    }
	
	return newHub
}

func (h *Hub) Run() {
    for {
        msg := <-h.Broadcast
        for id, client := range h.clients {
            if id == msg.SenderID {
                //err := client.WriteMessage(websocket.TextMessage, msg.Data)
                err := client.WriteMessage(websocket.TextMessage, []byte("Hellow there send back from server ID " + id))
                if err != nil {
                    log.Println(err)
                    client.Close()
                    delete(h.clients, id)
                }
            }
        }
    }
}

func (h *Hub) WaitingSocketConnection() {
	/*
		This method is a improve version of Run(), 
		instead of using loops this method uses key to get the connection instantly 0(1)
	*/

    for {
        msg := <-h.Broadcast
        if client, ok := h.clients[msg.SenderID]; ok {
			var result []byte
			if msg.Callback != nil {
				result = (*msg.Callback)(string(msg.Data))
                log.Println("       MESSAGE:: ", string(msg.Data))
			} else {
				log.Println("       [WARNING] Callback function cannot be empty")
                client.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(
                    websocket.CloseInternalServerErr, 
                    "Server error - internal server error while operating"), 
                    time.Now().Add(time.Second),
                )
			}
            err := client.WriteMessage(websocket.TextMessage, result)
            if err != nil {
                log.Println(" [ERROR] Failed to WriteMessage: ", err)
                client.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(
                    websocket.CloseInternalServerErr, 
                    "Server error - internal server error while operating"), 
                    time.Now().Add(time.Second),
                )
                client.Close()
                delete(h.clients, msg.SenderID)
            }
        }
    }
}

func (h *Hub) Handler(w http.ResponseWriter, r *http.Request, callBack func(string) []byte) {
    conn, err := h.upgrader.Upgrade(w, r, nil)
    if err != nil {
        conn.WriteControl(websocket.CloseInternalServerErr, []byte("Internal Server Error"), time.Now().Add(5 * time.Second))
        log.Println("       [ERROR] Failed to Upgrade()", err)
        return
    }
    defer conn.Close()

    var senderID string
	ok := true
	for ok {
		/* 
			For each connections we generate a unqiue ID for that connections,
			If somehow the key already exist in the map we attempt to generate again
			until the key was unqiue and not found in the current map.
		*/
		senderID = hubID()
		if _, ok = h.clients[senderID]; !ok {
			break
		} 
	}
    h.clients[senderID] = conn


    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            conn.WriteControl(1014, []byte("Bad gateway — server acting as gateway received an invalid response"), time.Now().Add(5 * time.Second))
            log.Println("		[ERROR] Failed to ReadMessage()", err)
            delete(h.clients, senderID)
            return
        }
        h.Broadcast <- Message{
			msg, 
			senderID, 
			&callBack,
		}
    }
}

func hubID() (string) {
	rand.Seed(time.Now().UnixNano())
    letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()")
    b := make([]rune, 10)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
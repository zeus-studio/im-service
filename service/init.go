package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func StartWebSocket() {
	http.HandleFunc("/im", wsPage)
	http.ListenAndServe(":8080", nil)
}

func wsPage(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	fmt.Println("websocket 建立连接：", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

}

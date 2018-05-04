package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	//"time"

	"github.com/gorilla/websocket"
	"time"
	//"net/http"
	"fmt"
	"net/http"
)

/*
 测试WebSocket功能的客户端代码，与其他功能无关
*/
var addr = flag.String("addr", "127.0.0.1:10012", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/oapi/v1/watch/projects"}
	log.Printf("connecting to %s", u.String())

	var rh http.Header
	rh = make(map[string][]string)
	rh.Set("Authorization", "Bearer v2lnb5Ia-x2GEeYZb0Bl8DZbx92oHE3m1MLgjGAliyY")
	fmt.Println(u.String())
	//time.Sleep(20*time.Second)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), rh)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {

		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}

	}
}

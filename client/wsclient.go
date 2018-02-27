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
	"net/http"
)

var addr = flag.String("addr", "10.19.14.21:10012", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/watch/projects/wutest001"}
	log.Printf("connecting to %s", u.String())

	var rh http.Header
	rh = make(map[string] []string)
	rh.Set("Authorization","Bearer fFU-kOhDhxrfGX1fwjSfHefAyW_S8dLe5KTktw8Rbj8")
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

	for{



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


package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

func main() {

	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:8080",
		Path:   "/ws",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	go sender(conn)

	go reader(conn)

	fmt.Println("Exiting the application")
}

func sender(conn *websocket.Conn) {

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text() // Get the current line of text
		if text == "" {
			break // Exit loop if an empty line is entered
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Fatal(err)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error from scanner:", err)
	}

}

func reader(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Message from server : %v\n", string(msg))
	}
}

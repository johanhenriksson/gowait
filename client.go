package gowait

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	out  chan interface{}
}

func Connect(upstream string) *Client {
	// todo: token auth
	conn, _, err := websocket.DefaultDialer.Dial(upstream, nil)
	if err != nil {
		log.Fatal("websocket dial failed:", err)
	}
	client := &Client{
		conn: conn,
		out:  make(chan interface{}),
	}

	// write pump
	go func() {
		for msg := range client.out {
			if err := conn.WriteJSON(msg); err != nil {
				log.Println("Send failed:", err)
				return
			}
		}
	}()

	// read pump
	go func() {
		for {
			t, bytes, err := conn.ReadMessage()
			if t == -1 {
				return
			}
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			fmt.Fprintln(RealStdout, "RECV:", t, string(bytes))
		}
	}()

	return client
}

func (c *Client) Close() {
	if err := c.conn.Close(); err != nil {
		log.Fatal("Close ws failed:", err)
	}
}

func (c *Client) Send(msg interface{}) {
	c.out <- msg
}

func (c *Client) SendInit(taskdef *Taskdef) {
	c.Send(InitMessage{
		Type:    InitMsg,
		ID:      taskdef.ID,
		Version: "0.3.3",
		Task:    *taskdef,
	})
}

func (c *Client) SendReturn(id string, result interface{}) {
	c.Send(ReturnMessage{
		Type:       ReturnMsg,
		ID:         id,
		Result:     result,
		ResultType: "any",
	})
}

func (c *Client) SendError(id string, err string) {
	c.Send(ErrorMessage{
		Type:  ErrorMsg,
		ID:    id,
		Error: err,
	})
}

func (c *Client) SendLog(id, file, data string) {
	c.Send(LogMessage{
		Type: LogMsg,
		ID:   id,
		File: file,
		Data: data,
	})
}

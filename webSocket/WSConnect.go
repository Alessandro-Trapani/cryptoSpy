package websocket

import (
	"cryptoSpy/config"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func ConnectToWebSocket(url string, message string) *websocket.Conn  {
  headers := http.Header{}
  headers.Add("Authorization",config.API_KEY)
	conn, _, err := websocket.DefaultDialer.Dial(url,headers)

	if err != nil {
    fmt.Printf("ERROR : \n %s",err)
	}

  err = conn.WriteMessage(websocket.TextMessage, []byte(message))
  if err != nil {
  fmt.Printf("ERROR : \n %s",err)
	
  }
  return conn
}

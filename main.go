/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cryptoSpy/cmd"
	websocket "cryptoSpy/webSocket"
	"fmt"
)

func main() {
	cmd.Execute()
  conn := websocket.ConnectToWebSocket("ws://coinbase.ws-ds.md.coinapi.io/",`{
  "type": "hello",
  "heartbeat": false,
  "subscribe_data_type": ["trade"],
  "subscribe_filter_symbol_id": [
    "COINBASE_SPOT_BTC_USD",
    "ETH-USD"
  ]
}`)
 
for {

  _, wsMessage, err := conn.ReadMessage()

  if err != nil {
  fmt.Printf("ERROR : \n %s",err)
	
  }


  fmt.Println(string(wsMessage))
}
}

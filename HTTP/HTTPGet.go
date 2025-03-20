package HTTP

import (
	"cryptoSpy/config"
	"fmt"
	"io"
	"net/http"
)

func GET(url string) {

	client := &http.Client{}

  req,err := http.NewRequest("GET",url,nil)

  if err != nil{
    fmt.Printf("ERROR : %s",err)
    return

  }
  
  req.Header.Add("Accept", "text/plain")
  req.Header.Add("X-CoinAPI-Key", config.API_KEY)

  res, err := client.Do(req)

  if(err != nil){
    fmt.Printf("ERROR : %s",err) 
    return
   }

   defer res.Body.Close()

  body, err := io.ReadAll(res.Body)

  if err != nil {
    fmt.Printf("ERROR : %s",err) 
    return
  }

  fmt.Println(string(body))
}

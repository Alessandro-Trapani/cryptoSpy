package HTTP
import (
	"cryptoSpy/config"
	"fmt"
	"io"
	"net/http"
)

func GET(url string, headers map[string]string) string {

	client := &http.Client{}

  req,err := http.NewRequest("GET",url,nil)

  if err != nil{
    fmt.Printf("ERROR : %s",err)

  }
  

  req.Header.Add("Accept", "text/plain")
  req.Header.Add("X-CoinAPI-Key", config.API_KEY)

  for _,value := range headers{

    req.Header.Add(value,headers[value])
  }


  res, err := client.Do(req)

  if(err != nil){
    fmt.Printf("ERROR : %s",err) 
   }

   defer res.Body.Close()

  body, err := io.ReadAll(res.Body)

  if err != nil {
    fmt.Printf("ERROR : %s",err) 

  }


  return string(body)
}

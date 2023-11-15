package veeamservertime

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func getServerTime(servername string, port int) {
  reqUrl := fmt.Sprintf("https://%s:%d/api/v1/serverTime", servername, port)
  req, err := http.NewRequest("GET", reqUrl, nil)
  if err != nil {
    panic(err)
  }
  req.Header.Add("x-api-version", "1.1-rev0")
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err)
  }

  fmt.Println(res)
  fmt.Println(string(body))
}
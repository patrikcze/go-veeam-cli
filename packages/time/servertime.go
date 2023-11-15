package veeamtime

import (
  "crypto/tls"
  "fmt"
  "net/http"
  "io/ioutil"
)

type ServerTimeResponse struct {
	ServerTime string `json:"serverTime"`
	TimeZone   string `json:"timeZone"`
}

func GetServerTime(servername string, port int) string {
  tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	reqUrl := fmt.Sprintf("https://%s:%d/api/v1/serverTime", servername, port)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("x-api-version", "1.1-rev0")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
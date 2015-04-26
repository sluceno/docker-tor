package main

import (
  "github.com/sluceno/socks"
  "fmt"
  "net/http"
  // "io/ioutil"
  "os"
  "time"
)

func main() {

  // proxy_host := os.Getenv("PRIVOXY_PORT_8118_TCP_ADDR")
  proxy_host := os.Getenv("DOCKERTOR_TOR_1_PORT_9050_TCP_ADDR")
  proxy_port := 9050
  proxy_connection_url := fmt.Sprintf("%s:%d", proxy_host, proxy_port)

  fmt.Printf("main - %s\n", proxy_connection_url)

  dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, proxy_connection_url)
  tr := &http.Transport{Dial: dialSocksProxy}
  httpClient := &http.Client{Transport: tr}

  for {
    TestHttpsGet(httpClient, "http://94.23.230.21/")
    // TestHttpsGet(http.Client, "http://94.23.230.21/")

    time.Sleep(2 * time.Second)
  }
}

func TestHttpsGet(c *http.Client, url string) {
// func TestHttpsGet(c *http.Client, url string) (bodyText string, err error) {
  c.Get(url)
  // if err != nil { return }
  // defer resp.Body.Close()
  // body, err := ioutil.ReadAll(resp.Body)
  // if err != nil { return }
  // bodyText = string(body)
}

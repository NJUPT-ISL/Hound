package lib

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SendToken(token string){
	config := tls.Config{
		InsecureSkipVerify:true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &config,
	}
	client := &http.Client{Transport: tr}
	_, err := client.PostForm(
		"https://"+os.Getenv("MasterUrl")+"/tokens/send",
		url.Values{"Host": {os.Getenv("hostname")}, "Token": {token}})
	if err != nil {
		log.Printf("Send Token message failed.")
		panic(err)
	}
}
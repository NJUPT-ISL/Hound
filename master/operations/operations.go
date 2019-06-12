package operations

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"
	"../models"
)

func PostOperations(Method string,NodeName string,token string,images []string){
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
	data := url.Values{}
	for _,image := range images{
		data.Add("imageName",image)
	}
	req, err := http.NewRequest("POST", "https://"+NodeName+":8081/image/"+Method, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("token", token)
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func GetOperations(Method string,NodeName string,token string){
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
	req, err := http.NewRequest("POST", "https://"+NodeName+":8081/image/"+Method, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("token", token)
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func Pull(NodeName string,images []string){
	token,err := models.TokenQuery(NodeName)
	if err != nil{
		panic(err)
	}
	PostOperations("pull", NodeName, token.Token, images)
}

func Remove(NodeName string,images []string){
	token,err := models.TokenQuery(NodeName)
	if err != nil{
		panic(err)
	}
	PostOperations("remove", NodeName, token.Token, images)
}

func Prune(NodeName string){
	token,err := models.TokenQuery(NodeName)
	if err != nil{
		panic(err)
	}
	GetOperations("prune", NodeName, token.Token)
}


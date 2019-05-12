package operations

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	req, err := http.NewRequest("POST", "https://"+NodeName+"/image/"+Method, strings.NewReader(data.Encode()))
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
	req, err := http.NewRequest("POST", "https://"+NodeName+"/image/"+Method, nil)
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

func Pull(NodeName string,token string,images []string){
	PostOperations("pull", NodeName, token, images)
}

func Remove(NodeName string,token string,images []string){
	PostOperations("remove", NodeName, token, images)
}

func Prune(NodeName string,token string){
	GetOperations("prune", NodeName, token)
}


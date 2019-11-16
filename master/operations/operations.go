package operations

import (
	"crypto/tls"
	"github.com/NJUPT-ISL/Hound/master/log"
	"github.com/NJUPT-ISL/Hound/master/models"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func PostOperations(Method string, NodeName string, token string, images []string) {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	data := url.Values{}
	for _, image := range images {
		data.Add("imageName", image)
	}
	req, err := http.NewRequest("POST", "https://"+NodeName+":8081/image/"+Method, strings.NewReader(data.Encode()))
	if err != nil {
		log.ErrPrint(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("token", token)
	_, err = client.Do(req)
	if err != nil {
		log.ErrPrint(err)
	}
}

func GetOperations(Method string, NodeName string, token string) {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", "https://"+NodeName+":8081/image/"+Method, nil)
	if err != nil {
		log.ErrPrint(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("token", token)
	_, err = client.Do(req)
	if err != nil {
		log.ErrPrint(err)
	}
}

func Pull(NodeName string, images []string) {
	token, err := models.GetToken(NodeName)
	if err != nil {
		log.ErrPrint(err)
	}
	PostOperations("pull", NodeName, token.Token, images)
}

func Remove(NodeName string, images []string) {
	token, err := models.GetToken(NodeName)
	if err != nil {
		log.ErrPrint(err)
	}
	PostOperations("remove", NodeName, token.Token, images)
}

func Prune(NodeName string) {
	token, err := models.GetToken(NodeName)
	if err != nil {
		log.ErrPrint(err)
	}
	GetOperations("prune", NodeName, token.Token)
}

func Parallelize(workers int, nodes []string, images []string, Do func(string, []string)) {
	var stop <-chan struct{}
	pieces := len(nodes)
	toProcess := make(chan string, pieces)
	for _, n := range nodes {
		toProcess <- n
	}
	close(toProcess)
	if pieces < workers {
		workers = pieces
	}
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for n := range toProcess {
				select {
				case <-stop:
					return
				default:
					Do(n, images)
				}
			}
		}()
	}
	wg.Wait()
}

func ParallelizeWithString(workers int, nodes []string, Do func(string)) {
	var stop <-chan struct{}
	pieces := len(nodes)
	toProcess := make(chan string, pieces)
	for _, n := range nodes {
		toProcess <- n
	}
	close(toProcess)
	if pieces < workers {
		workers = pieces
	}
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for n := range toProcess {
				select {
				case <-stop:
					return
				default:
					Do(n)
				}
			}
		}()
	}
	wg.Wait()
}

package lib

import (
	"bytes"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"log"
	"sync"
	"time"
)

var workers = 5

// List all images
func ListAllImages() ([]types.ImageSummary, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return images, err
}

// Pull Docker Images
func ImagePull(ImageName string) (io.ReadCloser, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.Println(err)
	}
	out, err := cli.ImagePull(ctx, ImageName, types.ImagePullOptions{All: false})
	fmt.Print("The Pull Task: " + ImageName + ", start at :")
	fmt.Println(time.Now())
	if err != nil {
		log.Println(err)
	}
	defer func() {
		outerr := out.Close()
		if err != nil {
			panic(outerr)
		}
	}()
	if _, err := new(bytes.Buffer).ReadFrom(out); err != nil {
		log.Println(err)
	}
	fmt.Print("The Pull Task: " + ImageName + ", end at :")
	fmt.Println(time.Now())
	return out, err
}

func ImagesPull(images []string) {
	var stop <-chan struct{}
	pieces := len(images)
	toProcess := make(chan string, pieces)
	for _, image := range images {
		toProcess <- image
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
			for image := range toProcess {
				select {
				case <-stop:
					return
				default:
					if _, err := ImagePull(image); err != nil {
						log.Println(err)
					}
				}
			}
		}()
	}
	wg.Wait()
	ResetWorker(5)
}

// remove images
func ImageRemove(ImageName string, Force bool) ([]types.ImageDeleteResponseItem, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	out, err := cli.ImageRemove(ctx, ImageName, types.ImageRemoveOptions{Force: Force})
	if err != nil {
		panic(err)
	}
	return out, err
}

// clean the images in the node that can not be used
func ImagesPrune() (types.ImagesPruneReport, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	out, err := cli.ImagesPrune(ctx, filters.Args{})
	if err != nil {
		panic(err)
	}
	return out, err
}

// show the docker information
func DockerInfo() (types.Info, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	out, err := cli.Info(ctx)
	return out, err
}

func ResetWorker(num int) {
	workers = num
}

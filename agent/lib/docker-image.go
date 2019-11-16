package lib

import (
	"bytes"
	"github.com/NJUPT-ISL/Hound/agent/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"sync"
)

var workers = 5

// List all images
func ListAllImages() ([]types.ImageSummary, error) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.ErrPrint(err)
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.ErrPrint(err)
	}
	return images, err
}

// Pull Docker Images
func ImagePull(ImageName string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.ErrPrint(err)
	}
	out, err := cli.ImagePull(ctx, ImageName, types.ImagePullOptions{All: false})
	log.Print("run pull task: " + ImageName)
	if err != nil {
		log.ErrPrint(err)
	}
	defer out.Close()
	if _, err := new(bytes.Buffer).ReadFrom(out); err != nil {
		log.ErrPrint(err)
	}
	log.Print("end pull task " + ImageName)
	return err
}

func ImagesPull(images []string) {
	Parallelize(workers, images, ImagePull)
}

// remove images
func ImageRemove(ImageName string, Force bool) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.ErrPrint(err)
	}
	log.Print("run remove task: " + ImageName)
	if _, err := cli.ImageRemove(ctx, ImageName, types.ImageRemoveOptions{Force: Force}); err != nil {
		log.ErrPrint(err)
	}
	return err
}

func ImagesRemove(images []string, force bool) {
	Parallelize(workers, images, func(s string) error {
		return ImageRemove(s, force)
	})
}

// clean the images in the node that can not be used
func ImagesPrune() (types.ImagesPruneReport, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.ErrPrint(err)
	}
	log.Print("run prune task.")
	out, err := cli.ImagesPrune(ctx, filters.Args{})
	if err != nil {
		log.ErrPrint(err)
	}
	return out, err
}

// show the docker information
func DockerInfo() (types.Info, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.ErrPrint(err)
	}
	out, err := cli.Info(ctx)
	return out, err
}

func Parallelize(workers int, images []string, Do func(string) error) {
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
					if err := Do(image); err != nil {
						log.ErrPrint(err)
					}
				}
			}
		}()
	}
	wg.Wait()
}

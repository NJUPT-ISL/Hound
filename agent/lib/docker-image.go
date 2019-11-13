package lib

import (
	"bytes"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"time"
)

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
	out, err := cli.ImagePull(ctx, ImageName, types.ImagePullOptions{All: false})
	fmt.Println("Start:")
	fmt.Println(time.Now())
	if err != nil {
		panic(err)
	}
	defer func() {
		outerr := out.Close()
		if err != nil {
			panic(outerr)
		}
	}()
	var buf = new(bytes.Buffer)
	_, ioerr := buf.ReadFrom(out)
	//_, ioerr := io.Copy(os.NewFile(uintptr(syscall.Stdout), "/dev/null"), out)
	if err != nil {
		panic(ioerr)
	}
	fmt.Println("End:")
	fmt.Println(time.Now())
	return out, err
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

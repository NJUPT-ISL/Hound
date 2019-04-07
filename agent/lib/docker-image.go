package lib

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"os"
)


func ListAllImages() ([]types.ImageSummary, error) {
	defer func() {
		if err := recover() ; err!= nil{
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
	return images,err
}

func ImagePull(ImageName string) (io.ReadCloser, error) {
	defer func() {
		if err := recover() ; err!= nil{
			fmt.Println(err)
		}
	}()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	out, err := cli.ImagePull(ctx, ImageName, types.ImagePullOptions{All:false})
	if err != nil {
		panic(err)
	}
	defer func() {
		outerr := out.Close()
		if err != nil {
			panic(outerr)
		}
	}()
	_, ioerr := io.Copy(os.Stdout, out)
	if err != nil{
		panic(ioerr)
	}
	return out, err
}
func ImageRemove(ImageName string, Force bool) ([]types.ImageDeleteResponseItem, error) {
	defer func() {
		if err := recover() ; err!= nil{
			fmt.Println(err)
		}
	}()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	out, err := cli.ImageRemove(ctx, ImageName, types.ImageRemoveOptions{Force:Force,})
	if err != nil {
		panic(err)
	}
	return out, err
}


func ImagesPrune() (types.ImagesPruneReport, error) {
	defer func() {
		if err := recover() ; err!= nil{
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

func DockerInfo()  (types.Info, error){
	defer func() {
		if err := recover() ; err!= nil{
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
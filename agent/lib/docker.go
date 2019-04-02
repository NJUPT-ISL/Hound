package lib

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/gin-gonic/gin"
)

type imageImpl struct {
	name string
	tag string
	ID string
	CreateTime string
}

func list_all_images() (map[string]imageImpl) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	image_list := make(map[string]imageImpl)
	for _, image := range images {
		im := imageImpl{
			name:image.ID,
		}
		image_list.append(im)
	}
	return image_list
}

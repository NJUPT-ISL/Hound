package lib

import (
	"testing"
)


func TestListAllImages(t *testing.T) {
	list, err := ListAllImages()
	if err != nil {
		panic(err)
	}
	for _,image := range list{
		t.Logf(image.ID,image.RepoDigests,image.RepoTags)
	}
}


func TestImagePull(t *testing.T) {
	_, err := ImagePull("alpine")
	if err != nil {
		panic(err)
	}
	t.Logf("ImagePullPass")
}


func TestImageRemove(t *testing.T) {
	_, err := ImageRemove("alpine", false)
	if err != nil {
		panic(err)
	}
	t.Logf("ImageRemovePass")
}


func TestImagesPrune(t *testing.T) {
	_,err := ImagesPrune()
	if err != nil {
		panic(err)
	}
}
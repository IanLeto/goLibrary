package apis_test

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	docker "github.com/docker/docker/client"
	"github.com/stretchr/testify/suite"
	"goLibrary/utils"
	"testing"
)

type DockerSuite struct {
	suite.Suite
	client *docker.Client
	ctx    context.Context
}

func (s *DockerSuite) SetupTest() {
	var (
		err error
	)
	s.ctx = context.Background()
	s.client, err = docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	s.NoError(err)
}

// 如何使用
func (s *DockerSuite) TestDockerPs() {
	images, err := s.client.ImageList(s.ctx, types.ImageListOptions{
		All:     false,
		Filters: filters.Args{},
	})
	s.NoError(err)
	utils.JustSee(images)
}

func TestDocker(t *testing.T) {
	suite.Run(t, new(DockerSuite))
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	list, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", list)
}

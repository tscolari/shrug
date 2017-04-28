package garden

import (
	"math/rand"

	uuid "github.com/satori/go.uuid"
)

type Container struct {
	Name string
}

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Containers() []Container {
	containers := []Container{}

	total := rand.Intn(100) + 1
	for i := 0; i < total; i++ {
		containers = append(containers, Container{Name: uuid.NewV4().String()})
	}

	return containers
}

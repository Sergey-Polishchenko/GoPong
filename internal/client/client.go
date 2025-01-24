package client

type Client interface{}

type client struct{}

func NewClient(host, port string) Client {
	return &client{}
}

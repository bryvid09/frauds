package domain

import "fmt"

type Client string

func NewClient(client string) (Client, error) {
	if(client == "") {
		return Client(""), fmt.Errorf("Client cannot be empty")
	}
	return Client(client), nil
}
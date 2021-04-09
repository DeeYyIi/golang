package main

import (
	"fmt"
	"strconv"
)

// New creates a new handle API client.
func New(totuus bool, num int) ClientService {
	var tmpClient = Client{true, 5}
	return &tmpClient
}

/*
Client for handle API
*/
type Client struct {
	boolean bool
	integer int
}

type Oma int

func (o Oma) Error() string {
	return strconv.Itoa(int(o))
}

// ClientService is the interface for Client methods
type ClientService interface {
	DoStuff(intti uint64, sana string) string
	DoSomeStuff(varlen ...int) int
	DoSecretStuff(mappi map[string]int) error
}

func main() {
<<<<<<< HEAD
	fmt.Println("Hello World")
=======

	var oma ClientService

	oma = New(true, 5)

	fmt.Println("Hello oma", ":", &oma, ":", oma.DoStuff(543, "jouman"))
}

func (t *Client) DoStuff(intti uint64, sana string) string {
	fmt.Println(t, ":", intti, ":", sana)

	return "ok"
}

func (t *Client) DoSomeStuff(varlen ...int) int {
	for i, j := range varlen {
		fmt.Println(t, ":", i, ":", j)
	}
	return 0
}

func (t *Client) DoSecretStuff(mappi map[string]int) error {
	for k, v := range mappi {
		fmt.Println(t, ":", k, ":", v)
	}

	return Oma(0)
>>>>>>> 0c47be3... interface testing
}

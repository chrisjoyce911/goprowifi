package goprowifi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client ...
type Client struct {

	// Use this switch to see all network communication.
	Debug bool
	// http interface can be used when testing
	myDoer Doer

	service
}

type service struct {
	client *Client
}

// Doer to make testing easer !
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// NewClient ...
func NewClient(debug bool) *Client {

	g := &Client{}

	const API = "http://10.5.5.9/gp/gpControl/"

	g.Debug = debug

	// set default http interface to use
	g.myDoer = http.DefaultClient

	return g
}

func (g *Client) request(url string) (bodyBytes []byte, statusCode int, err error) {

	fmt.Printf("[Debug] URL : %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := g.myDoer.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("[Debug] body : %s\n", string(bodyBytes))

	statusCode = resp.StatusCode
	fmt.Printf("[Debug] status : %d\n", statusCode)

	return
}

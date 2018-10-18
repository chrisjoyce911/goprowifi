package goprowifi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client ...
type Client struct {
	Protocol string
	Host     string
	Port     string
	Base     string
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

	g.Host = "10.5.5.9"
	g.Port = "80"
	g.Protocol = "http"
	g.Base = "gp/gpControl/"
	g.Debug = debug

	// set default http interface to use
	g.myDoer = http.DefaultClient

	return g
}

func (g *Client) genericRequest(url string, method string, body io.Reader) (bodyBytes []byte, statusCode int, err error) {

	API := fmt.Sprintf("http://10.5.5.9/gp/gpControl/%s", url)

	fmt.Printf("[Debug] URL : %s\n", API)
	fmt.Printf("[Debug] Reduest : %v\n", body)

	req, err := http.NewRequest(method, API, body)
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

func (g *Client) genericGET(url string) (bodyBytes []byte, statusCode int, err error) {
	return g.genericRequest(url, "GET", nil)
}

func (g *Client) genericPOST(url string, body io.Reader) (bodyBytes []byte, statusCode int, err error) {
	return g.genericRequest(url, "POST", body)
}

func (g *Client) getURL() string {
	return fmt.Sprintf("%v://%v:%v/%s", g.Protocol, g.Host, g.Port, g.Base)
}

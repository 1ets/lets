package lets

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// Type for saving header value.
type httpHeader struct {
	Name  string
	Value string
}

type basicAuth struct {
	Username string
	Password string
}

// Type for saving oy builder params.
type HttpBuilder struct {
	url      string
	client   *http.Client
	headers  []*httpHeader
	response *http.Response
	basic    basicAuth
}

// Set http client with default configuration.
func (h *HttpBuilder) Default() {
	defaultTransport := http.DefaultTransport.(*http.Transport)

	// Create new Transport that ignores self-signed SSL
	customTransport := &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		DialContext:           defaultTransport.DialContext,
		MaxIdleConns:          defaultTransport.MaxIdleConns,
		IdleConnTimeout:       defaultTransport.IdleConnTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	h.client = &http.Client{
		Timeout:   time.Duration(5) * time.Second,
		Transport: customTransport,
	}
}

// Manual set http builder.
func (h *HttpBuilder) SetClient(client *http.Client) {
	h.client = client
}

// Set End Point
func (h *HttpBuilder) SetUrl(url string) {
	// LogD("HttpBuilder: set endPoint to \"%s\"", url)

	h.url = url
}

// Set Basic Auth
func (h *HttpBuilder) SetBasicAuth(username, password string) {
	h.basic.Username = username
	h.basic.Password = password
}

// Setting up header name and value.
func (h *HttpBuilder) AddHeader(name string, value string) {
	// LogD("HttpBuilder: add header \"%s\":\"%s\"", name, value)

	for _, header := range h.headers {
		if header.Name == name {
			header.Value = value
			return
		}
	}

	h.headers = append(h.headers, &httpHeader{
		Name:  name,
		Value: value,
	})
}

// Post request.
func (h *HttpBuilder) Post(endPoint string, body interface{}) (fullUrl, response string, err error) {
	fullUrl = fmt.Sprintf("%s%s", h.url, endPoint)

	var payloadString string
	if reflect.TypeOf(body) == reflect.TypeOf([]byte(nil)) {
		payloadString = string(body.([]byte))
	} else {
		payloadString = ToJson(body)
	}

	LogI("HttpBuilder: POST \"%s\"\n%s", fullUrl, payloadString)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest(http.MethodPost, fullUrl, payload)
	if err != nil {
		return
	}

	// Header Setup
	for _, header := range h.headers {
		LogI("HttpBuilder: SetHeader: %s: %s", header.Name, header.Value)
		req.Header.Add(header.Name, header.Value)
	}

	// Basic Auth
	if h.basic.Username != "" {
		req.SetBasicAuth(h.basic.Username, h.basic.Password)
	}

	h.response, err = h.client.Do(req)
	if err != nil {
		return
	}
	defer h.response.Body.Close()

	resBody, err := io.ReadAll(h.response.Body)
	if err != nil {
		return
	}

	response = string(resBody)
	LogI("HttpBuilder: Response: %v \n%s", h.response.StatusCode, response)
	return
}

// Get request.
func (h *HttpBuilder) Get(endPoint string, body interface{}) (fullUrl, response string, err error) {
	fullUrl = fmt.Sprintf("%s%s", h.url, endPoint)

	var payloadString string
	if reflect.TypeOf(body) == reflect.TypeOf([]byte(nil)) {
		payloadString = string(body.([]byte))
	} else {
		v, _ := query.Values(body)
		payloadString = v.Encode()
	}

	LogI("HttpBuilder: GET \"%s\"\n%s", fullUrl, payloadString)

	// payload := strings.NewReader(payloadString)
	req, err := http.NewRequest(http.MethodGet, fullUrl+"?"+payloadString, &strings.Reader{})
	if err != nil {
		return
	}

	// Header Setup
	for _, header := range h.headers {
		LogI("HttpBuilder: SetHeader: %s: %s", header.Name, header.Value)
		req.Header.Add(header.Name, header.Value)
	}

	h.response, err = h.client.Do(req)
	if err != nil {
		return
	}
	defer h.response.Body.Close()

	resBody, err := io.ReadAll(h.response.Body)
	if err != nil {
		return
	}

	response = string(resBody)
	LogI("HttpBuilder: Response: %v \n%s", h.response.StatusCode, response)
	return
}

func (h *HttpBuilder) GetRequest() *http.Response {
	return h.response
}

func (h *HttpBuilder) GetResponse() *http.Response {
	return h.response
}

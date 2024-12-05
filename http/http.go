package http

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Request struct {
	Endpoint string
	Headers  map[string]string
	Query    map[string][]string
	Body     map[string]interface{}
}

func HandleRequest(buffer []byte, bufferLen int) error {
	parts := strings.Split(string(buffer[:bufferLen]), "\r\n")
	method, endpoint, httpVersion, err := getMethodEndpointAndHttpVersion(parts[0])
	if err != nil {
		panic(err) // send response
	}

	reqHeaders, err := getRequestHeaders(parts[1 : len(parts)-2])
	if err != nil {
		panic(err)
	}

	body := parts[len(parts)-1]

	fmt.Println("Method:", method)
	fmt.Println("ReqBody:", body)
	fmt.Println("Endpoint:", endpoint)
	fmt.Println("HttpVersion:", httpVersion)
	fmt.Println("RequestHeaders:", reqHeaders)

	urlA, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	req := Request{
		Endpoint: endpoint,
		Headers:  reqHeaders,
		Query:    urlA.Query(),
	}

	if method == "GET" {
		handleGetRequest(&req)
	}

	return nil
}

func getMethodEndpointAndHttpVersion(payload string) (string, string, string, error) {
	parts := strings.Split(payload, " ")
	if len(parts) != 3 {
		return "", "", "", errors.New("invalid request format")
	}
	return parts[0], parts[1], parts[2], nil
}

func getRequestHeaders(payload []string) (map[string]string, error) {
	reqHeadersMap := make(map[string]string, len(payload))
	for _, p := range payload {
		parts := strings.SplitN(p, ":", 2)
		key := strings.TrimSpace(parts[0]) // should we trim only left space
		value := strings.TrimSpace(parts[1])
		reqHeadersMap[key] = value
	}
	return reqHeadersMap, nil
}

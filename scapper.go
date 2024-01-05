package twitchgraphapi

import "github.com/valyala/fasthttp"

type Scrapper struct {
	httpClient *fasthttp.Client
	headers    *fasthttp.Request
}

func New() *Scrapper {
	return &Scrapper{
		httpClient: &fasthttp.Client{},
		headers:    initHeaders(),
	}
}

func (s *Scrapper) SetClientId(clientId string) *Scrapper {
	s.headers.Header.Set("Client-Id", clientId)
	return s
}

func initHeaders() *fasthttp.Request {
	req := fasthttp.AcquireRequest()

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")

	return req

}

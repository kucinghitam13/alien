package web

import (
	"io"
	"net/http"
)

func getDefaultRequest(method, url string, payload io.Reader) (request *http.Request) {
	request, _ = http.NewRequest(method, url, payload)
	request.Header = getDefaultHeader()
	return
}

func getDefaultHeader() http.Header {
	return http.Header{
		"Accept":                   {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"Accept-Language":          {"en-US,en;q=0.5"},
		"Connection":               {"keep-alive"},
		"Host":                     {"baak.gunadarma.ac.id"},
		"Upgrade-Insecure-Request": {"1"},
		"User-Agent":               {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0"},
	}
}

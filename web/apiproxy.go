package web

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ApiProxyHandler(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(StreamBaseAddr)
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.ServeHTTP(w, r)
}
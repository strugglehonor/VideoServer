package web

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ApiProxyHandler(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(StreamBaseAddr)
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.ServeHTTP(w, r)
}

// URL:/api handler
// func ApiHandler(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		errMsg := fmt.Sprintf("error occur reading request.Body:%v", err)
// 		SendErrorResponse(w, errMsg, http.StatusInternalServerError)
// 	}

// 	res := &ApiDef{}
// 	if err = json.Unmarshal(body, res); err != nil {
// 		errMsg := fmt.Sprintf("error occur when Unmarshal:%v", err)
// 		SendErrorResponse(w, errMsg, http.StatusInternalServerError)
// 	}
// }

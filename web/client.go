package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var httpclient *http.Client

func init() {
	httpclient = &http.Client{}
}

// api handler
func ApiHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	
	apidef := &ApiDef{}
	err = json.Unmarshal(body, apidef)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	
}

// make new request
func request(w http.ResponseWriter, r *http.Request, apidef *ApiDef) {
	switch apidef.Method {
	case http.MethodGet:
		req, _ := http.NewRequest(apidef.Method, apidef.URL, nil)
		req.Header = r.Header
		resp, err := httpclient.Do(req)
		if err != nil {
			SendErrorResponse(w, err.Error(), resp.StatusCode)
			return
		}
	case http.MethodPost:
		req, _ := http.NewRequest(apidef.Method, apidef.URL, bytes.NewBuffer([]byte(apidef.ReqBody)))
		req.Header = r.Header
		resp, err := httpclient.Do(req)
		if err != nil {
			SendErrorResponse(w, err.Error(), resp.StatusCode)
			return
		}
	case http.MethodDelete:
		req, _ := http.NewRequest(apidef.Method, apidef.URL, nil)
		req.Header = r.Header
		resp, err := httpclient.Do(req)
		if err != nil {
			SendErrorResponse(w, err.Error(), resp.StatusCode)
			return
		}
	default:
		SendErrorResponse(w, fmt.Sprintf("%s method not support", apidef.Method), http.StatusNotImplemented)
	}
}

// 
func WriteResp(w http.ResponseWriter, resp *http.Response) error {
	w.WriteHeader(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	w.Write(body)
	return nil
}
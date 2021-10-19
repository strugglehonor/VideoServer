package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// api handler
func ApiHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	
	var apidef ApiDef
	err = json.Unmarshal(body, &apidef)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	if apidef.Method == http.MethodGet {
		resp, err := http.Get(apidef.URL)
		if err != nil {
			log.Fatal(err)
		}

		if err = WriteResp(w, resp); err != nil {
			log.Fatal(err)
		}
	}

	if apidef.Method == http.MethodPost {

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
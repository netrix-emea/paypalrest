package common

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/wedancedalot/paypalrest/auth"
    "io/ioutil"
    "net/http"
    "strings"
)

func DoRequest(auth *auth.PayPalAuth, method string, url string, data interface{}) (response []byte, statusCode int, err error) {
    var req *http.Request
    var resp *http.Response
    var body []byte
    client := &http.Client{}
    var outData []byte
    if data != nil {
        outData, err = json.Marshal(data)
    }
    req, err = http.NewRequest(method, url, strings.NewReader(string(outData)))
    if err == nil {
        token, _ := auth.GetToken()
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Authorization", fmt.Sprintf("%s %s", auth.Token_type, token))
        resp, err = client.Do(req)
        if resp != nil {
            defer resp.Body.Close()
            body, err = ioutil.ReadAll(resp.Body)
        }
        if err == nil {
            statusCode = resp.StatusCode
            var generic_response map[string]interface{}
            err = json.Unmarshal(body, &generic_response)
            if name, ok := generic_response["name"]; ok == true && name == "VALIDATION_ERROR" {
                err = errors.New(string(body))
            } else {
                response = body
            }
        } else {
            err = errors.New(fmt.Sprintf("Invalid PayPal API response: %v", err))
        }
    }
    return
}


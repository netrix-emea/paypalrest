package authorizations

import (
    "fmt"
    "testing" 
    "strings"
    "net/http"
    "net/http/httptest"
    "github.com/cfsalguero/paypalrest/auth"
)

var _ts *httptest.Server
var _auth *auth.PayPalAuth

func init() {
    _ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL)
        fmt.Fprintln(w, "{}")
    }))
    // ClientId and ClientId are note needed because we are mocking the http request
    _auth, _ = auth.NewAuth("sandbox", "your_client_id", "your_client_secret")
    _auth.SetEndpoint(_ts.URL)
}

func TestGetAuthorization(t *testing.T) { 
    _, err := GetAuthorization(_auth, "this is an id")
    if err != nil {
        if err.Error() == "PayPal call status code: 404" {
            t.Log("GetAuthorization test passed. ")
        } else {
            t.Error("Error: " + err.Error())
        }
    } else  {
        t.Log("GetAuthorization test passed. ")
    }

}

func TestCaptureAuthorization(t *testing.T) { 
    _, err := CaptureAuthorization(_auth, "1", "USD", 1, false)

    if err != nil {
        if err.Error() == "PayPal call status code: 404" {
            t.Log("CaptureAuthorization test passed. ")
        } else {
            t.Error("Error: " + err.Error())
        }
    } else  {
        t.Log("CaptureAuthorization test passed. ")
    }
}    

func TestVoidAuthorization(t *testing.T) { 
    _, err := VoidAuthorization(_auth, "1")

    if err != nil {
        if pos := strings.Index(err.Error(), "status code: 404"); pos > -1 {
            t.Log("CaptureAuthorization test passed. ")
        } else {
            t.Error("Error: " + err.Error())
        }
    } else  {
        t.Log("CaptureAuthorization test passed. ")
    }
}    

func TestReauthorize(t *testing.T) { 
    _, err := Reauthorize(_auth, "1", 1.2, "USD")

    if err != nil {
        if pos := strings.Index(err.Error(), "status code: 404"); pos > -1 {
            t.Log("CaptureAuthorization test passed. ")
        } else {
            t.Error("Error: " + err.Error())
        }
    } else  {
        t.Log("CaptureAuthorization test passed. ")
    }
}    


package common

import (
    "testing"
    "fmt"
    "github.com/cfsalguero/paypalrest/auth"
)

func TestDoRequest(t *testing.T) {
    auth, _ := auth.NewAuth("sandbox", "EOJ2S-Z6OoN_le_KS1d75wsZ6y0SFdVsY9183IvxFyZp", "EClusMEUk8e9ihI7ZdVLF5cZ6y0SFdVsY9183IvxFyZp")
    url := fmt.Sprintf("%s%s%s", auth.Endpoint, "/v1/payments/payment/authorization/", "1")
    _, statusCode, err := DoRequest(auth, "GET", url, nil)
    if err != nil {
        if statusCode == 404 {
            t.Log("DoRequest test passed. ")
        } else {
            t.Error("Error: " + err.Error())
        }
    } else  {
        t.Log("DORequest test passed. ")
    }
}


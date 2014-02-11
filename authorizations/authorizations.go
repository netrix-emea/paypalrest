package authorizations

import (
    "fmt"
    "encoding/json"
    "github.com/edrans/paypalrest/auth"
    "github.com/edrans/paypalrest/common"
)

type AmountDetails struct {
    Subtotal string `json:"subtotal"`
}

type Amount struct {
    Total    string        `json:"total"`
    Currency string        `json:"currency"`
    Details  AmountDetails `json:"details"`
}

type HATEOAS_Link struct {
    Href   string `json:"href"`
    Rel    string `json:"rel"`
    Method string `json:"method"`
}

type AuthorizationResponse struct {
    Id            string         `json:"id"`
    CreateTime    string         `json:"create_time"`
    UpdateTime    string         `json:"update_time"`
    State         string         `json:"state"`
    Amount        Amount         `json:"amount"`
    ParentPayment string         `json:"parent_payment"`
    ValidUntil    string         `json:"valid_until"`
    Links         []HATEOAS_Link `json:"links"`
}

type CaptureResponse struct {
    Id             string         `json:"id"`
    CreateTime     string         `json:"create_time"`
    UpdateTime     string         `json:"update_time"`
    State          string         `json:"state"`
    Amount         CaptureAmount  `json:"amount"`
    ParentPayment  string         `json:"parent_payment"`
    IsFinalCapture bool           `json:"is_final_capture"`
    Links          []HATEOAS_Link `json:"links"`
}

type CaptureAmount struct {
    Currency string `json:"currency"`
    Total    string `json:"total"`
}

type CaptureDetails struct {
    Amount         CaptureAmount `json:"amount"`
    IsFinalCapture bool          `json:"is_final_capture"`
}

type VoidResponse struct {
    Id            string         `json:"id"`
    CreateTime    string         `json:"create_time"`
    UpdateTime    string         `json:"update_time"`
    State         string         `json:"state"`
    Amount        Amount         `json:"amount"`
    ParentPayment string         `json:"parent_payment"`
    Links         []HATEOAS_Link `json:"links"`
}

func GetAuthorization(auth *auth.PayPalAuth, id string) (response AuthorizationResponse, err error) {
    url := fmt.Sprintf("%s%s%s", auth.Endpoint, "/v1/payments/payment/authorization/", id)
    api_response, statusCode, err := common.DoRequest(auth, "GET", url, nil)
    if err == nil {
        if statusCode == 200 {
            err = json.Unmarshal(api_response, &response)
        } else {
            err = fmt.Errorf("PayPal call status code: %d", statusCode)
        }
    }
    return
}

func CaptureAuthorization(auth *auth.PayPalAuth, id string, currency string, total float64, is_final_capture bool) (response *CaptureResponse, err error) {
    amount := CaptureAmount{Currency: currency, Total: fmt.Sprintf("%0.2f", total)}
    captureDetails := &CaptureDetails{Amount: amount, IsFinalCapture: is_final_capture}
    url := fmt.Sprintf("%s%s%s/capture", auth.Endpoint, "/v1/payments/authorization/", id)
    if err == nil {
        api_response, statusCode, err := common.DoRequest(auth, "POST", url, captureDetails)
        if err == nil {
            if statusCode == 200 {
                err = json.Unmarshal(api_response, &response)
            } else {
                err = fmt.Errorf("API call status code %d, %s", statusCode, string(api_response))
            }
        } 
    } 
    return
}

func VoidAuthorization(auth *auth.PayPalAuth, id string) (response *VoidResponse, err error) {
    url := fmt.Sprintf("%s%s%s/void", auth.Endpoint, "/v1/payments/authorization/", id)
    api_response, statusCode, err := common.DoRequest(auth, "POST", url, nil)
    if err == nil {
        if statusCode == 200 {
            err = json.Unmarshal(api_response, &response)
        } else {
            err = fmt.Errorf("Api call status code: %d, %s", statusCode, string(api_response))
        }
    } 
    return
}

func Reauthorize(auth *auth.PayPalAuth, id string, amount float64, currency string) (response *VoidResponse, err error) {
    url := fmt.Sprintf("%s%s%s/reauthorize", auth.Endpoint, "/v1/payments/payment/authorization/", id)
    api_response, statusCode, err := common.DoRequest(auth, "POST", url, nil)
    if err == nil {
        if statusCode == 200 {
            err = json.Unmarshal(api_response, &response)
        } else {
            err = fmt.Errorf("API call status code: %d, %s", statusCode, string(api_response))
        }
    }
    return
}


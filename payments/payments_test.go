package payments

import (
    "testing" 
    "net/http"
    "fmt"
    "net/http/httptest"
    "github.com/cfsalguero/paypalrest/auth"
)

var _ts *httptest.Server
var _auth *auth.PayPalAuth

func init() {
    _ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "{}")
    }))
    // ClientId and ClientId are note needed because we are mocking the http request
    _auth, _ = auth.NewAuth("sandbox", "your_client_id", "your_client_secret")
    _auth.SetEndpoint(_ts.URL)
}

func TestPayWithCreditCard(t *testing.T) { 

    credit_card := CreditCard{
        Number:      "4417119669820331",
        Type:        "visa",
        ExpireMonth: "11",
        ExpireYear:  "2018",
        Cvv2:        "874",
        FirstName:   "Joe",
        LastName:    "Shopper",
        BillingAddress: Address{
            Line1:       "52 N Main ST",
            City:        "Johnstown",
            CountryCode: "US",
            PostalCode:  "43210",
            State:       "OH",
        },
    }

    transaction := TransactionRequest{
        Amount: Amount{
            Total:    "7.47",
            Currency: "USD",
            Details: AmountDetails{
                Subtotal: "7.41",
                Tax:      "0.03",
                Shipping: "0.03",
            },
        },
        Description: "a description",
    }


    _, err := PayWithCreditCard(_auth, credit_card, transaction)
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

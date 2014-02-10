package example

import (
    "fmt"
    "github.com/edrans/paypalrest/auth"
    "github.com/edrans/paypalrest/payments"
)

func example() {
    auth, err := auth.NewAuth("sandbox", "your_client_id", "your_client_secret")
    if err == nil {
        token, _ := auth.GetToken()
        fmt.Printf("\n\nAuth: %+v, Token: %+v\n\n", auth, token)
    }

    credit_card := payments.CreditCard{
        Number:      "4417119669820331",
        Type:        "visa",
        ExpireMonth: "11",
        ExpireYear:  "2018",
        Cvv2:        "874",
        FirstName:   "Joe",
        LastName:    "Shopper",
        BillingAddress: payments.Address{
            Line1:       "52 N Main ST",
            City:        "Johnstown",
            CountryCode: "US",
            PostalCode:  "43210",
            State:       "OH",
        },
    }

    transaction := payments.TransactionRequest{
        Amount: payments.Amount{
            Total:    "7.47",
            Currency: "USD",
            Details: payments.AmountDetails{
                Subtotal: "7.41",
                Tax:      "0.03",
                Shipping: "0.03",
            },
        },
        Description: "a description",
    }

    response, err := payments.PayWithCreditCard(auth, credit_card, transaction)
    if err == nil {
        fmt.Printf("\n\n%+v\n\n", response)
    } else {
        fmt.Printf("\n\nError:\n%+v\n\n", err)
    }
}


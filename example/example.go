package example

import (
    "fmt"
    "github.com/wedancedalot/paypalrest/auth"
    "github.com/wedancedalot/paypalrest/payments"
)

func payWithCC() {
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

func payWithPaypal() {
    auth, err := auth.NewAuth("sandbox", "your_client_id", "your_client_secret")
    if err == nil {
        // Save the token somewhere for later
        token, _ := auth.GetToken()
        fmt.Printf("\n\nAuth: %+v, Token: %+v\n\n", auth, token)
    }

    redirect_urls := payments.RedirectURL{
        ReturnURL: "http://example.com/return_url",
        CancelURL: "http://example.com/cancel_url",
    }

    transaction := payments.TransactionRequest{
        Amount: payments.Amount{
            Total:    "6.00",
            Currency: "USD",
            Details: payments.AmountDetails{
                Subtotal: "6.00",
                Tax:      "0.00",
                Shipping: "0.00",
            },
        },
        Description: "Test product",
        ItemList: payments.ItemList{
            Items: []payments.Item{payments.Item{
                Quantity:   "1", 
                Name:       "Big Red Hat", 
                Price:      "6.00",  
                SKU:        "sku123", 
                Currency:   "USD",
            }},
        },
    }

    response, err := payments.PayWithPaypal(auth, redirect_urls, transaction)
    if err == nil {
        fmt.Printf("\n\n%+v\n\n", response)
        // Redirect user to response.Links[1] which will use your ReturnUrl
        
        // Create new auth with the token you've saved previously
        //      auth, err := auth.NewAuthFromToken("sandbox", token)

        // Get paymentId and payerId from get parameters and execute payment
        //      payments.ExecutePaypalPayment(auth, paymentId, payerID)
    } else {
        fmt.Printf("\n\nError:\n%+v\n\n", err)
    }
} 

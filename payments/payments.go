package payments

import (
    "encoding/json"
    "fmt"
    "github.com/edrans/paypalrest/auth"
    "github.com/edrans/paypalrest/common"
    "time"
)

type Address struct {
    Line1       string `json:"line1"`
    Line2       string `json:"line2"`
    City        string `json:"city"`
    CountryCode string `json:"country_code"`
    PostalCode  string `json:"postal_code"`
    State       string `json:"state,omitempty"`
    Phone       string `json:"phone,omitempty"`
}

type ShippingAddress struct {
    RecipientName string `json:"recipient_name,omitempty"`
    Type          string `json:"type,omitempty"`
    Line1         string `json:"line1,omitempty"`
    Line2         string `json:"line2,omitempty"`
    City          string `json:"city,omitempty"`
    CountryCode   string `json:"country_code,omitempty"`
    PostalCode    string `json:"postal_code,omitempty"`
    State         string `json:"state,omitempty"`
    Phone         string `json:"phone,omitempty"`
}

type CreditCard struct {
    Id             string  `json:"id,omitempty"`
    PayerId        string  `json:"payer_id,omitempty"`
    Number         string  `json:"number"`
    Type           string  `json:"type"`
    ExpireMonth    string  `json:"expire_month"`
    ExpireYear     string  `json:"expire_year"`
    Cvv2           string  `json:"cvv2"`
    FirstName      string  `json:"first_name"`
    LastName       string  `json:"last_name"`
    BillingAddress Address `json:"billing_address"`
    State          string  `json:"state"`
    ValidUntil     string  `json:"valid_until"`
}

type CreditCardToken struct {
    CreditCardId string `json:"credit_card_id"`
    PayerId      string `json:"payer_id"`
    Last4        string `json:"last4"`
    Type         string `json:"type"`
    ExpireYear   string `json:"expire_year"`
    ExpireMonth  string `json:"expire_month"`
}

type FundingInstrument struct {
    CreditCard      CreditCard      `json:"credit_card"`
    CreditCardToken CreditCardToken `json:"credit_card_token,omitempty"`
}

type FundingInstrumentRequest struct {
    CreditCard CreditCard `json:"credit_card"`
}

type AmountDetails struct {
    Subtotal string `json:"subtotal"`
    Tax      string `json:"tax"`
    Shipping string `json:"shipping"`
}

type PayerInfo struct {
    Email           string          `json:"email,omitempty"`
    FirstName       string          `json:"first_name,omitempty"`
    LastName        string          `json:"last_name,omitempty"`
    PayerId         string          `json:"payer_id,omitempty"`
    Phone           string          `json:"phone,omitempty"`
    ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

type creditCardAndBillingAddress struct {
    CreditCard
    BillingAddress Address `json:"billing_address"`
}

type Payer struct {
    PaymentMethod      string              `json:"payment_method"`
    FundingInstruments []FundingInstrument `json:"funding_instruments"`
    PayerInfo          PayerInfo           `json:"payer_info"`
}

type PayerRequest struct {
    PaymentMethod      string                     `json:"payment_method"`
    FundingInstruments []FundingInstrumentRequest `json:"funding_instruments"`
}

type PaymetRequest struct {
    Intent string `json:"intent"`
    Payer  Payer  `json:"payer"`
}

type Item struct {
    Quantity string `json:"quantity,omitempty"`
    Name     string `json:"name,omitempty"`
    Price    string `json:"price,omitempty"`
    Currency string `json:"currency,omitempty"`
    SKU      string `json:"sku,omitempty"`
}

type ItemList struct {
    Items           []Item          `json:"items,omitempty"`
//  ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

type Amount struct {
    Total    string        `json:"total"`
    Currency string        `json:"currency"`
    Details  AmountDetails `json:"details"`
}

type RelatedResource struct {
    Sale struct {
        Amount struct {
            Currency string `json:"currency"`
            Total    string `json:"total"`
        }   `json:"amount"`
        CreateTime time.Time `json:"create_time"`
        Id         string    `json:"id"`
        Links      []struct {
            Href   string `json:"href"`
            Method string `json:"method"`
            Rel    string `json:"rel"`
        }   `json:"links"`
        ParentPayment string    `json:"parent_payment"`
        State         string    `json:"state"`
        UpdateTime    time.Time `json:"update_time"`
    }   `json:"sale"`
    Authorization struct {
        Amount struct {
            Currency string `json:"currency"`
            Details  struct {
                Subtotal string `json:"subtotal"`
            }   `json:"details"`
            Total string `json:"total"`
        }   `json:"amount"`
        CreateTime time.Time `json:"create_time"`
        Id         string    `json:"id"`
        Links      []struct {
            Href   string `json:"href"`
            Method string `json:"method"`
            Rel    string `json:"rel"`
        }   `json:"links"`
        ParentPayment string    `json:"parent_payment"`
        State         string    `json:"state"`
        UpdateTime    time.Time `json:"update_time"`
        ValidUntil    time.Time `json:"valid_until"`
    }   `json:"authorization"`
}

type Transaction struct {
    Amount           Amount        `json:"amount"`
    Description      string        `json:"description"`
    ItemList         ItemList      `json:"item_list,omitempty"`
    RelatedResources []RelatedResource `json:"related_resources"`
}

type TransactionRequest struct {
    Amount      Amount      `json:"amount"`
    Description string      `json:"description"`
    ItemList    ItemList    `json:"item_list,omitempty"`
}
type RedirectURL struct {
    ReturnURL string `json:"return_url"`
    CancelURL string `json:"cancel_url"`
}

type HATEOAS_Link struct {
    Href   string `json:"href"`
    Rel    string `json:"rel"`
    Method string `json:"method"`
}

type CreateRequest struct {
    Intent       string               `json:"intent,omitempty"`
    Payer        PayerRequest         `json:"payer,omitempty"`
    Transactions []TransactionRequest `json:"transactions,omitempty"`
    RedirectURLs RedirectURL          `json:"redirect_urls,omitempty"`
    PayerId      string               `json:"payer_id,omitempty"`
}

type CreateResponse struct {
    Intent       string         `json:"intent"`
    Payer        Payer          `json:"payer"`
    Transactions []Transaction  `json:"transactions"`
    RedirectURLs []RedirectURL  `json:"redirect_urls"`
    Id           string         `json:"id"`
    CreateTime   time.Time      `json:"create_time"`
    State        string         `json:"state"`
    UpdateTime   time.Time      `json:"update_time"`
    Links        []HATEOAS_Link `json:"links"`
}

func PayWithCreditCard(auth *auth.PayPalAuth, credit_card CreditCard, transaction TransactionRequest) (response CreateResponse, err error) {
    r := CreateRequest{
        Intent: "sale",
        Payer: PayerRequest{
            PaymentMethod:      "credit_card",
            FundingInstruments: []FundingInstrumentRequest{{CreditCard: credit_card}},
        },
        Transactions: []TransactionRequest{transaction},
    }

    return create(auth, "/v1/payments/payment", r)
}

func AuthorizeWithCreditCard(auth *auth.PayPalAuth, credit_card CreditCard, transaction TransactionRequest) (response CreateResponse, err error) {
    r := CreateRequest{
        Intent: "authorize",
        Payer: PayerRequest{
            PaymentMethod:      "credit_card",
            FundingInstruments: []FundingInstrumentRequest{{CreditCard: credit_card}},
        },
        Transactions: []TransactionRequest{transaction},
    }

    return create(auth, "/v1/payments/payment", r)
}

func PayWithPaypal(auth *auth.PayPalAuth, redirect_url RedirectURL, transaction TransactionRequest) (response CreateResponse, err error) {
    r := CreateRequest{
        Intent: "sale",
        Payer: PayerRequest{
            PaymentMethod: "paypal",
        },
        Transactions: []TransactionRequest{transaction},
        RedirectURLs: redirect_url,
    }

    return create(auth, "/v1/payments/payment", r)
}

func ExecutePaypalPayment(auth *auth.PayPalAuth, paymentId, payerId string) (response CreateResponse, err error) {
    r := CreateRequest{
        PayerId: payerId,
    }

    return create(auth, "/v1/payments/payment/" + paymentId + "/execute/", r)
}

func create(auth *auth.PayPalAuth, url string, r CreateRequest) (response CreateResponse, err error) {
    var api_response []byte
    var statusCode int
    
    url = fmt.Sprintf("%s%s", auth.Endpoint, url)

    api_response, statusCode, err = common.DoRequest(auth, "POST", url, r)
    if err == nil {
        if statusCode == 201 || statusCode == 200 {
            err = json.Unmarshal(api_response, &response)
        } else {
            err = fmt.Errorf("API call status code: %d. Response: %s", statusCode, string(api_response))
        }
    }
    return
}

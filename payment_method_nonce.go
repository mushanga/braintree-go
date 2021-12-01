package braintree

type PaymentMethodNonce struct {
	Type             string                     `xml:"type"`
	Nonce            string                     `xml:"nonce"`
	Details          *PaymentMethodNonceDetails `xml:"details"`
	ThreeDSecureInfo *ThreeDSecureInfo          `xml:"three-d-secure-info"`
}

type PaymentMethodNonceDetails struct {
	CardType        string            `xml:"card-type"`
	Last2           string            `xml:"last-two"`
	BIN             string            `xml:"bin"`
	ExpirationMonth string            `xml:"expiration-month"`
	ExpirationYear  string            `xml:"expiration-year"`
	PayerInfo       map[string]string `xml:"payer-info"`
}

package braintree

type Verification struct {
	XMLName         string `xml:"verification"`
	CVVResponseCode string `xml:"cvv-response-code,omitempty"`
	CreatedAt       string `xml:"created-at,omitempty"`
}

type Verifications struct {
	Verification []*Verification `xml:"verification"`
}

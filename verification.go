package braintree

import (
	"encoding/xml"
)

type Verification struct {
	XMLName         xml.Name
	CVVResponseCode string `xml:"cvv-response-code,omitempty"`
}

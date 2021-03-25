package types

//Money - represents a monetary amount
//in minimum units (cents, kopecks, diramas, etc.).
type Money int64

//PAN (Payment Card Number) - cards number.
type PAN string

//Currency
type Currency string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // use Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//PaymentSource - represents information about the payment source.
type PaymentSource struct {
	Type     string //'card'
	Number   string // "5058 xxxx xxxx 8888"
	Balance  Money  // balance in diramas
}

package card

import (
	"bank-first/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {

	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: 20_000_00,
			Active:  false,
		},
		{
			PAN:     "5058 xxxx xxxx 5555",
			Balance: -15_000_00,
			Active:  true,
		},
	}

	payment := PaymentSources(cards)
	for _, outputPrint := range payment {
		fmt.Println(outputPrint.Number)
	}
	//Output:
	//5058 xxxx xxxx 8888
}

//Another method of example test.
/*
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "2323 xxxx xxxx 2422",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "2221 xxxx xxxx 6566"
			Active:  false,
		},
		{
			PAN:     "3343 xxxx xxxx 9383",
			Balance: -10_000_00,
			Active:  false,
		},
	}

	payment := []types.PaymentSource(PaymentSources(cards))

	for _, outputPrint := range payment {
		fmt.Println(outputPrint.Number)
	}
	//Output: //2323 xxxx xxxx 2422
}
*/

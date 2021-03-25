package card

import (
	"bank/pkg/bank/types"
)

//Max - a function that returns the maximum payment.
func PaymentSources(cards []types.Card) []types.PaymentSource {

	outputCard := []types.PaymentSource{}
	for _, operation := range cards {
		if operation.Balance < 0 {
			continue
		}
		if !operation.Active {
			continue
		}
		outputCard = append(outputCard, types.PaymentSource{
			Type:    "card",
			Number:  string(operation.PAN),
			Balance: operation.Balance,
		})
	}
	return outputCard
}

//Another method of function.
/*
func PaymentSources(cards []types.Card) []types.PaymentSource{

	var outputSlice []types.PaymentSource
	var output types.PaymentSource

	for _, card := range cards {
		if (card.Balance>0){
			if (card.Active==true){
				output.Type=card.Name
				output.Number=string(card.PAN)
				output.Balance=card.Balance
				outputSlice=append(outputSlice,output)
				}
			}
	}
	return outputSlice
}
*/

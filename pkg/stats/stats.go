package stats

import (
	"github.com/FaranushKarimov/bank/pkg/types"

)

//Avg рассчитывает среднюю сумму платежа~
func Avg(payments []types.Payment) types.Money {
	paySum := types.Money(0)
	payCount := types.Money(0)
	//payments := []types.Payment{}

	for i := 0; i < len(payments); i++ {
		paySum += payments[i].Amount
		payCount++
	}
	return paySum / payCount
}

//TotalInCategory  returned total sum in one category
func TotalInCategory(payments []types.Payment, category types.Category) (total types.Money) {

	for _, v := range payments {
		if category == v.Category {
			total += v.Amount
		}
	}
	return
}

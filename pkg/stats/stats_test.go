package stats

import (
	"testing"
	"github.com/FaranushKarimov/bank/v2/pkg/types"
	"reflect"
)


func TestCategoriesAvg_multiple(t *testing.T) {
	payments := []types.Payment{
	  {
		ID:       1,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       2,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       3,
		Category: "restaurant",
		Amount:   10_00,
	  },
	  {
		ID:       2,
		Category: "restaurant",
		Amount:   100_00,
	  },
	}
	  expected := map[types.Category]types.Money{
	  "cafe":100_00,
	  "restaurant":55_00,
	}
  
  
	res := CategoriesAvg(payments)
  
  
	if !reflect.DeepEqual(expected, res) {
	  t.Errorf(" got > %v want > %v", res, expected)
	} 
  
  }


  func TestPeriodsDynamic_positive(t *testing.T) {
	first := map[types.Category]types.Money{
	  "cafe": 20,
	  "auto": 14,
	}
	second := map[types.Category]types.Money{
	  "cafe": 35,
	  "auto": 17,
	}
	want := map[types.Category]types.Money{
	  "cafe": 15,
	  "auto": 3,
	}
  
	got := PeriodsDynamic(first, second)
  
	if !reflect.DeepEqual(want, got) {
	  t.Errorf(" got > %v \n want > %v", got, want)
	} 
  
  }
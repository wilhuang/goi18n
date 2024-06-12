package test

import (
	"fmt"
	"goi18n/lang"
	"testing"
)

func TestEnum(t *testing.T) {
	startVal, endVal := 4, 20558
	entrieEnum := lang.NewEnum([]lang.Code{lang.ENTRIE_4}, &lang.EnumOption[int]{
		Count:     endVal - startVal + 1,
		Values:    []int{startVal},
		ValueStep: 1,
	})
	for k := startVal; k <= endVal; k++ {
		keyName := fmt.Sprintf("ENTRIE_%d", k)
		if entrieEnum.ToCode(k) != lang.ToCode(keyName) {
			t.Fatalf(keyName + " EnumTest Error")
		}
	}
}

package test

import (
	"fmt"
	"goi18n/lang"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {
	nt := time.Now()
	langsTmp := lang.GetLocaleSort()
	t.Log(langsTmp)
	langs := make([]string, len(langsTmp))
	copy(langs, langsTmp[:])
	sort.Slice(langs, func(i, j int) bool {
		return rand.Intn(2) == 0
	})
	lang.SetLocaleSort(langs...)
	langsNew := lang.GetLocaleSort()
	for i, langstr := range langsNew {
		if langstr != langs[i] {
			t.Fatalf("Langs Error")
		}
	}
	codes := []lang.Code{
		lang.ENTRIE_SPECIAL,
		lang.ENTRIE_NORMAL,
		lang.ENTRIE_ORDER,
		lang.ENTRIE_4,
	}
	for i := lang.ENTRIE_4; i < lang.ENTRIE_20558; i++ {
		codes = append(codes, i)
	}
	for _, code := range codes {
		transAll := code.TransAll()
		for i, trans := range transAll {
			if code.Trans(langs[i]) != trans {
				t.Fatalf("Lang:%s[%d] TransAll() != Trans()", langs[i], code)
			}
		}
	}
	for _, locale := range langs {
		t.Log("Locale:", locale)
		lang.SetDefaultLocale(locale)
		if lang.GetDefaultLocale() != locale {
			t.Fatalf("SetDefaultLocale Error")
		}
		for _, code := range codes {
			trans := code.Trans(locale)
			if code.Default() != trans {
				t.Fatalf("Lang:%s[%d] Default() != Trans()", locale, code)
			}
			codeStr := code.String()
			if codeStr == "ERR_CODE_UNKNOW" || strings.HasPrefix(codeStr, "Code[") {
				t.Fatalf("Lang:%s[%d] String() Error", locale, code)
			}
			if lang.CodeString(codeStr) <= 0 {
				t.Fatalf("Lang:%s[%d] CodeString() Error", locale, code)
			}
			if lang.ToLocaleString(codeStr, locale) != trans {
				t.Fatalf("Lang:%s[%d] ToLocaleString() != Trans()", locale, code)
			}
			if code <= lang.ENTRIE_4 {
				t.Log(trans)
			}
		}
		if locale == "xx" {
			fmt.Println("zh-aa", locale)
		}
		t.Log(lang.ENTRIE_ORDER.Trans(locale, "a", 2.0, 3))
		t.Log(lang.ENTRIE_ORDER.Trans(locale, []string{"a", "b"}, 2.0, 3))
		t.Log(lang.ENTRIE_ORDER.Trans(locale, []float64{0.1, 0.2}, 2.0, 3))
		t.Log(lang.ENTRIE_ORDER.Trans(locale, []lang.Code{lang.ENTRIE_4, lang.ENTRIE_5, lang.ENTRIE_6}, 2.0, 3))

		t.Log("\n")
	}
	t.Log("Take Time:", time.Since(nt))
}

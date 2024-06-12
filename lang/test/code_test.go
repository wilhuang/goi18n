package test

import (
	"goi18n/lang"
	"strings"
	"testing"
	"time"
)

func TestCode(t *testing.T) {
	nt := time.Now()
	langs := lang.GetLocaleIds()
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
	// 常规方法测试
	for id, locale := range langs {
		t.Log("Locale:", locale)
		for _, code := range codes {
			trans := code.Trans(locale)
			if code.TransById(id) != trans {
				t.Fatalf("Lang:%s[%d] TranById() != Trans()", locale, code)
			}
			codeStr := code.String()
			if codeStr == "ERR_CODE_UNKNOW" || strings.HasPrefix(codeStr, "Code[") {
				t.Fatalf("Lang:%s[%d] String() Error", locale, code)
			}
			if lang.ToCode(codeStr) <= 0 {
				t.Fatalf("Lang:%s[%d] CodeString() Error", locale, code)
			}
			if lang.ToCode(codeStr).Trans(locale) != trans {
				t.Fatalf("Lang:%s[%d] ToLocaleString() != Trans()", locale, code)
			}
		}
	}
	// 默认方法与一些特例测试
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
			if code <= lang.ENTRIE_4 {
				t.Log(trans)
			}
		}
		t.Log(lang.ENTRIE_ORDER.Trans(locale, "a", 2.0, 3))
		t.Log(lang.ENTRIE_ORDER.Trans(locale, []string{"a", "b"}, 2.0, 3))
		t.Log(lang.ENTRIE_ORDER.Trans(locale, []float64{0.1, 0.2}, 2.0, 3))
		t.Log(lang.ENTRIE_ORDER.Trans(locale, []lang.Code{lang.ENTRIE_4, lang.ENTRIE_5, lang.ENTRIE_6}, 2.0, 3))

		t.Log("\n")
	}
	t.Log("Take Time:", time.Since(nt))
}

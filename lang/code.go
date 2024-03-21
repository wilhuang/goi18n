package lang

import (
	"fmt"
)

const _KEY = "k"

var (
	DefaultLocale    string = _KEY
	_DefaultLangCode uint16 = 0
)

// 支持的语言
var _Code_supported = map[string]uint16{
	_KEY: 0,
}

// IsLocaleSupport 语言是否支持
func IsLocaleSupport(locale string) bool {
	_, ok := _Code_supported[locale]
	return ok
}

func SetDefaultLocale(locale string) bool {
	langCode, ok := _Code_supported[locale]
	if ok {
		DefaultLocale = locale
		_DefaultLangCode = langCode
	}
	return ok
}

// 国际化的对象结构
type Code uint32

// _transOne 读取翻译
func (i Code) _transOne(locale uint16) string {
	if i <= _start {
		// TODO 后续添加内置错误码
		return "ERR_CODE_UNKNOW"
	}

	if i < _end {
		return _arr[locale][i-_start-1]
	}

	return fmt.Sprintf("Code[%d](%d)", locale, i)
}

// String 获取string类型的key值
func (i Code) String() string {
	return i._transOne(0)
}

func (i Code) _trans(locale uint16, args ...interface{}) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			if typ, ok := arg.(Code); ok {
				com = append(com, typ._transOne(locale))
			} else {
				com = append(com, arg) // arg as string scalar
			}
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}

// Default 获取默认语言的翻译
//   - args   Code类型或fmt支持的类型
func (i Code) Default(args ...interface{}) string {
	return i._trans(_DefaultLangCode, args...)
}

// Trans 获取指定语言的翻译
//   - locale 指定的语言版本
//   - args   Code类型或fmt支持的类型
func (i Code) Trans(locale string, args ...interface{}) string {
	langCode, ok := _Code_supported[locale]
	if !ok {
		langCode = _DefaultLangCode
	}
	return i._trans(langCode, args...)
}

// TransAll 获取全部翻译
func (i Code) TransAll(args ...interface{}) [langSize]string {
	var str [langSize]string
	for _, langCode := range LangCodes {
		str[langCode] = i._trans(langCode)
	}
	return str
}

// Error 获取指定语言错误类型的翻译
//   - locale 指定的语言版本
//   - args   Code类型或fmt支持的类型
func (i Code) Error(locale string, args ...interface{}) error {
	langCode, ok := _Code_supported[locale]
	if !ok {
		langCode = _DefaultLangCode
	}
	return NewError(i._trans(langCode, args...))
}

type CodeError struct {
	s string
}

func (e *CodeError) Error() string {
	return e.s
}

func NewError(s string) error {
	return &CodeError{s}
}

package lang

// ErrorById 获取错误类型的指定语言翻译
//   - id  指定的语言版本
//   - args Code类型或fmt支持的类型
func (c Code) ErrorById(id int, args ...any) error {
	return NewError(c.TransById(id, args...))
}

// ErrorDefault 获取错误类型的默认翻译
//   - args Code类型或fmt支持的类型
func (c Code) ErrorDefault(args ...any) error {
	return NewError(c.Default(args...))
}

// Error 获取错误类型的指定语言翻译
//   - l  指定的语言版本
//   - args Code类型或fmt支持的类型
func (c Code) Error(l string, args ...any) error {
	return NewError(c.Trans(l, args...))
}

// ErrorAll 获取错误类型的全部翻译
//   - args Code类型或fmt支持的类型
func (c Code) ErrorAll(args ...any) [I18N_LEN]error {
	var errs [I18N_LEN]error
	for id := range _LANGS {
		errs[id] = NewError(c.TransById(id, args...))
	}
	return errs
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

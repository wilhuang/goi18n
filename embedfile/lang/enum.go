package lang

type Enum[T Number] struct {
	valToCode map[T]Code
	codeToVal map[Code]T
	defCode   Code
	defVal    T
	sort      []T
}

type EnumOption[T Number] struct {
	Count               int  // 生成的枚举值个数，默认为0，表示不生效
	DefaultCode         Code // 非正确枚举值时的返回，默认为ERR_CODE_UNKNOW
	DefaultVal          T    // 非正确枚举对象时的返回，默认为0
	CodeStep            Code // 枚举对象步长，正数，等于0时，则默认为1
	IsCodeStepNegative  bool // 是否CodeStep为负数，默认为false
	Values              []T  // 枚举值，按照顺序排列
	ValueStep           T    // 枚举值步长，正数，等于0时，则默认为1
	IsValueStepNegative bool // 是否ValueStep为负数，默认为false
}

// NewEnum 生成一个lang的枚举管理对象
// 取codes、values、Count三者的最大长度，未达最大长度的将根据对应步长进行补齐，
//   - codes: 国际化对象，按照顺序排列
func NewEnum[T Number](codes []Code, option *EnumOption[T]) *Enum[T] {
	if option == nil {
		option = &EnumOption[T]{
			DefaultCode: ERR_CODE_UNKNOW,
			DefaultVal:  0,
			CodeStep:    1,
			ValueStep:   1,
		}
	} else {
		if option.CodeStep == 0 {
			option.CodeStep = 1
		}
		if option.ValueStep == 0 {
			option.ValueStep = 1
		}
	}

	if len(codes) == 0 {
		return nil
	}
	values := option.Values
	if len(values) == 0 {
		values = append(values, 0)
	}

	lenCode := len(codes)
	lenValue := len(values)
	option.Count = Max(option.Count, lenCode, lenValue)
	if option.Count == 0 {
		return nil
	}

	if option.Count > lenCode {
		codes = append(codes, make([]Code, option.Count-lenCode)...)
		for i := lenCode; i < option.Count; i++ {
			if option.IsCodeStepNegative {
				codes[i] = codes[i-1] - option.CodeStep
			} else {
				codes[i] = codes[i-1] + option.CodeStep
			}
		}
	}

	for _, code := range codes {
		if code <= _start || code >= _end {
			return nil
		}
	}

	if option.Count > lenValue {
		values = append(values, make([]T, option.Count-lenValue)...)
		for i := lenValue; i < option.Count; i++ {
			if option.IsValueStepNegative {
				values[i] = values[i-1] - option.ValueStep
			} else {
				values[i] = values[i-1] + option.ValueStep
			}
		}
	}
	m := make(map[T]Code, option.Count)
	for i, v := range values {
		m[v] = codes[i]
	}

	return &Enum[T]{
		valToCode: m,
		defCode:   option.DefaultCode,
		defVal:    option.DefaultVal,
		sort:      values,
	}
}

// ToCode 根据枚举值生成对应的Code
//   - v: 枚举值
func (e *Enum[T]) ToCode(v T) Code {
	if e == nil {
		return ERR_CODE_UNKNOW
	}
	if code, ok := e.valToCode[v]; ok {
		return code
	}
	return e.defCode
}

// ToValue 根据枚举对象生成对应的枚举值
//   - code: 枚举对象
func (e *Enum[T]) ToValue(code Code) T {
	if e == nil {
		return 0
	}
	if v, ok := e.codeToVal[code]; ok {
		return v
	}
	return e.defVal
}

// ListValue 返回枚举值列表
func (e *Enum[T]) ListValue() []T {
	if e == nil {
		return nil
	}
	return e.sort
}

// ListCode 返回枚举对象列表
func (e *Enum[T]) ListCode() []Code {
	if e == nil {
		return nil
	}
	codes := make([]Code, len(e.sort))
	for i, v := range e.sort {
		codes[i] = e.valToCode[v]
	}
	return codes
}

// ListKV 返回枚举值和枚举对象的对应关系
func (e *Enum[T]) ListKV() map[T]Code {
	if e == nil {
		return nil
	}
	return e.valToCode
}

// ListVK 返回枚举对象和枚举值的对应关系
func (e *Enum[T]) ListVK() map[Code]T {
	if e == nil {
		return nil
	}
	m := make(map[Code]T, len(e.valToCode))
	for k, v := range e.valToCode {
		m[v] = k
	}
	return m
}

// Add 添加一个枚举值和枚举对象的对应关系
func (e *Enum[T]) Add(value T, code Code) {
	if e == nil {
		return
	}
	e.valToCode[value] = code
	e.codeToVal[code] = value
	e.sort = append(e.sort, value)
}

// Remove 删除一个枚举值和枚举对象的对应关系
func (e *Enum[T]) Remove(value T) {
	if e == nil {
		return
	}
	code := e.valToCode[value]
	delete(e.valToCode, value)
	delete(e.codeToVal, code)
	for i, v := range e.sort {
		if v == value {
			e.sort = append(e.sort[:i], e.sort[i+1:]...)
			break
		}
	}
}

// RemoveCode 删除一个枚举值和枚举对象的对应关系
func (e *Enum[T]) RemoveCode(code Code) {
	if e == nil {
		return
	}
	value := e.codeToVal[code]
	delete(e.valToCode, value)
	delete(e.codeToVal, code)
	for i, v := range e.sort {
		if v == value {
			e.sort = append(e.sort[:i], e.sort[i+1:]...)
			break
		}
	}
}

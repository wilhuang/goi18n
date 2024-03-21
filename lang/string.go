package lang

const ERR_CODE_UNKNOW Code = 0 + iota

var root *node

// 使用 radix 树（或称为字典树）来存储字符串前缀
type node struct {
	children map[rune]*node
	code     Code
}

func addString(s string, code Code) {
	curr := root
	for _, ch := range s {
		if child, ok := curr.children[ch]; ok {
			curr = child
		} else {
			newNode := &node{
				children: make(map[rune]*node),
			}
			curr.children[ch] = newNode
			curr = newNode
		}
	}
	curr.code = code
}

// CodeString 通过string类型的key值获取Code对象
func CodeString(s string) Code {
	curr := root
	for _, ch := range s {
		if child, ok := curr.children[ch]; ok {
			curr = child
		} else {
			return ERR_CODE_UNKNOW
		}
	}
	return curr.code
}

// ToLocaleString 通过string类型的key值获取翻译
func ToLocaleString(s, locale string) string {
	if codeKey := CodeString(s); codeKey != ERR_CODE_UNKNOW {
		return codeKey.Trans(locale)
	}
	return s
}

// ToLocaleString 通过string类型的key值获取默认翻译
func ToLocaleDefault(s string) string {
	if codeKey := CodeString(s); codeKey != ERR_CODE_UNKNOW {
		return codeKey.Default()
	}
	return s
}

// ToLocaleStringAll 通过string类型的key值获取全部翻译
func ToLocaleStringAll(s string) [langSize]string {
	if codeKey := CodeString(s); codeKey != ERR_CODE_UNKNOW {
		return codeKey.TransAll()
	}
	var str [langSize]string
	for i := langSize; i > 0; i-- {
		str[i] = s
	}
	return str
}

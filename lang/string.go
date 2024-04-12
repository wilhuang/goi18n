package lang

const ERR_CODE_UNKNOW Code = 0 + iota

// 使用 radix 树（或称为字典树）来存储字符串前缀
type doubleLinkNode struct {
	parent *doubleLinkNode
	child  map[rune]*doubleLinkNode
	code   Code
}

func newDoubleLinkNode() *doubleLinkNode {
	return &doubleLinkNode{
		child: make(map[rune]*doubleLinkNode),
	}
}

func (r *doubleLinkNode) Add(s string, code Code) {
	curr := r
	for _, ch := range s {
		if curr.child != nil {
			if child, ok := curr.child[ch]; ok {
				curr = child
				continue
			}
		}
		newNode := &doubleLinkNode{
			parent: curr,
			child:  make(map[rune]*doubleLinkNode),
		}
		curr.child[ch] = newNode
		curr = newNode
	}
	curr.code = code
}

func (r *doubleLinkNode) End() {
	traverseDelLeaf(r)
	r.Gen()
}

var root *node

type node struct {
	child map[rune]*node
	code  Code
}

func (r *doubleLinkNode) Gen() {
	curr := r
	root = &node{
		child: make(map[rune]*node, len(curr.child)),
		code:  curr.code,
	}
	traverse(r, root)
}

func traverse(n *doubleLinkNode, nc *node) {
	if n == nil {
		return
	}
	// 遍历子节点
	for ch, child := range n.child {
		newNode := &node{
			code: child.code,
		}
		if len(child.child) > 0 {
			newNode.child = make(map[rune]*node, len(child.child))
			traverse(child, newNode)
		}
		nc.child[ch] = newNode
	}
}

func traverseDelLeaf(n *doubleLinkNode) {
	if n == nil {
		return
	}
	if len(n.child) == 0 && n.parent.code == 0 && len(n.parent.child) == 1 {
		n.parent.code = n.code
		n.parent.child = nil

		traverseDelLeaf(n.parent.parent)
		// 从叶子节点开始裁剪
		return
	}

	// 遍历子节点
	for _, child := range n.child {
		traverseDelLeaf(child)
	}
}

// CodeString 通过string类型的key值获取Code对象
func CodeString(s string) Code {
	curr := root
	for _, ch := range s {
		child, ok := curr.child[ch]
		if !ok {
			return ERR_CODE_UNKNOW
		}
		if len(child.child) == 0 {
			return child.code
		}
		curr = child
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

package main

import (
	"flag"
	"fmt"
	"goi18n/utils"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	FLAG_DIR  string
	FLAG_OUT  string
	FLAG_LANG string
	FLAG_SORT int
)

const LANG_PREFIX = "I18N_"

func init() {
	// -dir 指定读取路径
	flag.StringVar(&FLAG_DIR, "dir", "./lang/csv", `CSV目录, 示例: -dir ./lang/csv`)
	// -dir 指定读取路径
	flag.StringVar(&FLAG_OUT, "out", "./lang", `输出目录, 示例: -out ./lang`)
	// -dir 指定读取路径
	flag.StringVar(&FLAG_LANG, "lang", "k", `默认语言, 示例: -lang zh_CN`)
	// -sort 指定排序类型
	flag.IntVar(&FLAG_SORT, "sort", 0, `排序类型, 示例: -sort
0: 不排序
1: Ascll码升序(单表)
2: Ascll码降序(单表)
3: Ascll码升序(全局)
4: Ascll码降序(全局)`)
}

// langkey eg:zh_CN
//
// code,key,langkey1,langkey2,...
// code1,key1,keylang1,keylang2,...
// code2,key2,keylang1,keylang2,...

func main() {
	fmt.Println("goi18n RUN")
	fileList := GetFileList(FLAG_DIR)

	kvMap := make(map[string]map[string]string)
	codeMap := make(map[string]uint32)
	langSupport := make(map[string]struct{})
	var keySort []string

	for _, fileName := range fileList {
		fmt.Println("开始载入文件:" + fileName)
		info := utils.ReadCSV(fileName, kvMap, codeMap)
		if info == nil || len(info.Keys) == 0 {
			continue
		}
		for _, langstr := range info.Langs {
			langSupport[langstr] = struct{}{}
		}
		keySort = append(keySort, info.Keys...)
	}
	var langs []string
	for langstr := range langSupport {
		langs = append(langs, langstr)
	}
	if len(langs) == 0 {
		return
	}
	sort.Slice(langs, func(i, j int) bool {
		return langs[i] > langs[j]
	})

	if _, ok := langSupport[FLAG_LANG]; !ok {
		FLAG_LANG = langs[0]
	}

	fmt.Println("默认语言:", FLAG_LANG)

	build_code_go(langs, keySort, kvMap)
	build_key_go(langs, keySort, kvMap)
	fmt.Println("goi18n END")
	fmt.Println("Program completed. Press any key to exit.")

	var input string
	fmt.Scanln(&input) // 等待用户输入，然后程序会结束
}

func build_code_go(langs, keySort []string, kvMap map[string]map[string]string) {
	var langsBuilder, initBuilder, keyEnumBuilder, keyMapBuilder strings.Builder

	for _, key := range keySort {
		keyMapBuilder.WriteString(fmt.Sprintf("\n	addString(\"%s\", %s)", key, key))
		keyEnumBuilder.WriteString("\n	" + key)
		for _, langstr := range langs {
			desc := strings.ReplaceAll(kvMap[key][langstr], "\n", "")
			keyEnumBuilder.WriteString(" // " + desc)
		}
	}
	initBuilder.WriteString("SetDefaultLocale(" + LANG_PREFIX + strings.ToUpper(formatLocale(FLAG_LANG)) + ")")
	for i, locale := range langs {
		varName := LANG_PREFIX + strings.ToUpper(formatLocale(locale))
		langsBuilder.WriteString(fmt.Sprintf("\n	%s = %s", varName, strconv.Quote(locale)))
		initBuilder.WriteString(fmt.Sprintf("\n	Langs[%d] = %s", i, varName))
	}

	enumFile := fmt.Sprintf(`package lang

const (
	arrLen   = _end - _start - 1
	langSize = %d
)

var (
	Langs     [langSize]string
	LangCodes [langSize]uint16
	_arr      [langSize + 1][arrLen]string
)

const (%s
)

func init() {
	%s
	for i := uint16(0); i < langSize; i++ {
		LangCodes[i] = i + 1
		_Code_supported[Langs[i]] = i + 1
	}
}

const (
	_start Code = 1000 + iota%s
	_end
)
`, len(langs), langsBuilder.String(), initBuilder.String(), keyEnumBuilder.String())

	mapFile := fmt.Sprintf(`package lang

func init() {%s
}
	`, keyMapBuilder.String())

	err := output(FLAG_OUT, "enum.go", enumFile)
	if err != nil {
		panic(err)
	}

	err = output(FLAG_OUT, "map.go", mapFile)
	if err != nil {
		panic(err)
	}
}

func formatLocale(locale string) string {
	re := regexp.MustCompile("_+")
	locale = strings.ReplaceAll(locale, "-", "_")
	locale = strings.ReplaceAll(locale, " ", "_")

	locale = re.ReplaceAllString(locale, "_")
	// 移除首尾下划线
	return strings.Trim(locale, "_")
}

func build_key_go(langs, keySort []string, kvMap map[string]map[string]string) {
	builders := make([]strings.Builder, len(langs)+1)
	for _, key := range keySort {
		builders[len(langs)].WriteString("\n		" + strconv.Quote(key) + ",")
		for i, langstr := range langs {
			builders[i].WriteString("\n		" + strconv.Quote(kvMap[key][langstr]) + ",")
		}
	}

	// 合并连续的下划线为一个

	build_xx_go := func(langCode int, locale string, s string) {
		err := output(FLAG_OUT, "lang_"+formatLocale(locale)+".go", fmt.Sprintf(`package lang

func init() {
	_arr[%d] = [arrLen]string{%s
	}
}
`, langCode, s))
		if err != nil {
			panic(err)
		}
	}

	build_xx_go(0, "key", builders[len(langs)].String())
	for i, langstr := range langs {
		build_xx_go(i+1, langstr, builders[i].String())
	}
}

func GetFileList(dir string) []string {
	var fileList []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if !info.IsDir() {
			if strings.HasSuffix(strings.ToLower(path), "csv") {
				fileList = append(fileList, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

	if len(fileList) == 0 {
		fmt.Println("There is no CSV file in the directory")
	}
	return fileList
}

func output(dir, filename string, str string) error {
	f, err := os.OpenFile(filepath.Join(dir, filename), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}

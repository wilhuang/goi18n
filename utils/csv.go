package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type FileInfo struct {
	FileName string
	MinCode  uint32
	MaxCode  uint32
	Keys     []string
	Langs    []string
}

const COL_KEY = 1
const COL_CODE = 0
const COL_MAX = 2
const CODE_START uint32 = 1000

func ReadCSV(fileName string, kvMap map[string]map[string]string, codeMap map[string]uint32) (info *FileInfo) {
	// 打开CSV文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个GBK编码的reader
	reader := transform.NewReader(file, simplifiedchinese.GBK.NewDecoder())

	// 创建一个CSV Reader
	csvReader := csv.NewReader(reader)
	// 读取CSV文件中的内容

	row, err := csvReader.Read()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(row) < COL_MAX+1 {
		panic(fileName + " header is nil")
	}
	// 打印CSV文件中的内容
	langIdxMap := make(map[int]string)
	// kvMap := make(map[string]map[string]string)
	info = &FileInfo{
		FileName: fileName,
	}

	if len(row[COL_CODE]) > 0 {
		arr := strings.Split(row[0], "-")
		if len(arr) == 2 {
			min, err1 := strconv.ParseUint(arr[0], 10, 33)
			max, err2 := strconv.ParseUint(arr[1], 10, 33)
			if err1 == nil && err2 == nil && min < max {
				info.MinCode = uint32(min)
				info.MaxCode = uint32(max)
			}
		}
	}

	for j := len(row) - 1; j > COL_MAX-1; j-- {
		if len(row[j]) > 0 {
			langIdxMap[j] = row[j]
			info.Langs = append(info.Langs, row[j])
		}
	}

	var keys []string
	for {
		row, err = csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading CSV:", err)
			continue
		}
		if len(row) == 0 || len(row[COL_KEY]) == 0 {
			continue
		}
		key := row[COL_KEY]

		if _, ok := kvMap[key]; ok {
			panic(fileName + " has the same key:" + key)
		}

		kvMap[key] = make(map[string]string)
		keys = append(keys, key)

		for j := len(row) - 1; j > COL_MAX-1; j-- {
			if langStr, ok := langIdxMap[j]; ok {
				kvMap[key][langStr] = row[j]
			}
		}
	}

	info.Keys = keys
	return info
}

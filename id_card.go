package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 定义各省市的代码
var provinceCodes = map[string]string{
	"11": "北京",
	"12": "天津",
	"31": "上海",
	"50": "重庆",
	"15": "内蒙古",
	"65": "新疆",
	"54": "西藏",
	"64": "宁夏",
	"45": "广西",
	"81": "香港",
	"82": "澳门",
	"23": "黑龙江",
	"22": "吉林",
	"21": "辽宁",
	"13": "河北",
	"14": "山西",
	"63": "青海",
	"37": "山东",
	"41": "河南",
	"32": "江苏",
	"34": "安徽",
	"33": "浙江",
	"35": "福建",
	"36": "江西",
	"43": "湖南",
	"42": "湖北",
	"44": "广东",
	"46": "海南",
	"62": "甘肃",
	"61": "陕西",
	"51": "四川",
	"52": "贵州",
	"53": "云南",
	"71": "台湾",
}

// isChineseIDValid 验证是否为有效的中国身份证
func isChineseIDValid(id string) bool {
	if len(id) != 18 {
		return false
	}

	// 前17位应该是数字
	for i := 0; i < 17; i++ {
		if !isDigit(id[i]) {
			return false
		}
	}

	// 第18位可以是数字或者'X'
	if !isDigit(id[17]) && id[17] != 'X' && id[17] != 'x' {
		return false
	}

	// 验证前6位地区是否合法
	if !isValidAreaCode(id[:6]) {
		return false
	}

	// 验证日期部分是否合法
	birthDate := id[6:14]
	if !isValidDate(birthDate) {
		return false
	}

	// 验证校验码
	return verifyChecksum(id)
}

// isDigit 检查一个字符是否为数字
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// verifyChecksum 验证身份证校验码
func verifyChecksum(id string) bool {
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checksumChars := "10X98765432"

	sum := 0
	for i := 0; i < 17; i++ {
		num, _ := strconv.Atoi(string(id[i]))
		sum += num * weights[i]
	}

	mod := sum % 11
	expectedChecksum := checksumChars[mod]

	return strings.ToUpper(string(id[17])) == string(expectedChecksum)
}

// isValidDate 验证日期是否合法
func isValidDate(date string) bool {
	if len(date) != 8 {
		return false
	}

	year, err := strconv.Atoi(date[:4])
	if err != nil {
		return false
	}

	month, err := strconv.Atoi(date[4:6])
	if err != nil {
		return false
	}

	day, err := strconv.Atoi(date[6:])
	if err != nil {
		return false
	}

	// 检查日期是否有效
	location := time.Now().Location()
	d := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	_, err = time.ParseInLocation("2006-01-02", d, location)
	return err == nil
}

// isValidAreaCode 验证地区代码是否合法
func isValidAreaCode(areaCode string) bool {
	if len(areaCode) != 6 {
		return false
	}
	provinceCode := areaCode[:2]
	return provinceCodes[provinceCode] != ""
}
func IdCard() {
	tests := []string{
		"12345678901234567X", // 无效 ID（日期无效）
		"11010519491231002X", // 有效 ID
		"abcdefghijkmlnopqr", // 非法字符
		"12345678901234",     // 长度不足
		"11010519990101123X", // 有效 ID
		"992503199612100319", // 有效 ID
	}

	for _, test := range tests {
		fmt.Printf("ID: %s, Is Valid Chinese ID: %v\n", test, isChineseIDValid(test))
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/Eyevinn/mp4ff/mp4"
	"os"
)

func main() {
	//printUserMeteData()
	//printAllMeteData()
	//printJson()

	HWOD3()
}

//打印用户元数据信息
func printUserMeteData() {
	// 打开 MP4 文件
	file, err := os.Open("/Users/a1234/Downloads/merged.mp4")
	if err != nil {
		fmt.Printf("打开文件错误: %v\n", err)
		return
	}
	defer file.Close()

	// 解析 MP4 文件
	parsedMp4, err := mp4.DecodeFile(file)
	if err != nil {
		fmt.Printf("解析 MP4 文件错误: %v\n", err)
		return
	}

	// 打印 MP4 文件的标签
	printMP4Tags(parsedMp4)
}
func printMP4Tags(mp4file *mp4.File) {
	for _, box := range mp4file.Moov.Children {
		boxType := box.Type()
		switch boxType {
		case "udta":
			processUdtaBox(box.(*mp4.UdtaBox))
		default:
			fmt.Printf("跳过类型为: %s 的盒子\n", boxType)
		}
	}
}

func processUdtaBox(udtaBox *mp4.UdtaBox) {
	fmt.Println("处理用户数据盒 (udta):")
	for _, child := range udtaBox.Children {
		childType := child.Type()
		if childType == "meta" {
			processMetaBox(child.(*mp4.MetaBox))
		}
	}
}

func processMetaBox(metaBox *mp4.MetaBox) {
	fmt.Println("处理元数据盒 (meta):")
	for _, child := range metaBox.Children {
		childType := child.Type()
		if childType == "ilst" {
			processIlstBox(child)
		}
	}
}

func processIlstBox(box mp4.Box) {
	fmt.Println("处理项目列表盒 (ilst):")
	ilstBox := box.(*mp4.IlstBox)
	for _, child := range ilstBox.Children {
		tag := child.Type()
		processIlstChild(tag, child)
	}
}

func processIlstChild(tag string, box mp4.Box) {
	fmt.Printf("     - 处理 ilst 子项，键：%s\n", tag)
	childContainer, ok := box.(mp4.ContainerBox)
	if !ok {
		fmt.Printf("       - 错误: 无法将盒子转换为 ContainerBox，键：%s\n", tag)
		return
	}
	for _, subChild := range childContainer.GetChildren() {
		subChildType := subChild.Type()
		if subChildType == "data" {
			dataBox, ok := subChild.(*mp4.DataBox)
			if ok {
				// 根据数据类型适当地解析数据
				dataValue := parseDataBox(dataBox)
				fmt.Printf("       - 键：%s, 值：%s\n", tag, dataValue)
				fmt.Printf("       - 完整数据信息:\n")
				fmt.Printf("         - 数据长度：%d\n", len(dataBox.Data))
				fmt.Printf("         - 原始数据（十六进制）：%x\n", dataBox.Data)
				// 如果有其他相关属性，也可以打印出来
			} else {
				fmt.Printf("       - 错误: 无法将子项转换为 DataBox，类型：data\n")
			}
		}
	}
}

func parseDataBox(dataBox *mp4.DataBox) string {
	return string(dataBox.Data)
}

// 打印所有元数据
func printAllMeteData() {
	// 打开 MP4 文件
	file, err := os.Open("/Users/a1234/Downloads/merged.mp4")
	if err != nil {
		fmt.Printf("打开文件错误: %v\n", err)
		return
	}
	defer file.Close()

	// 解析 MP4 文件
	parsedMp4, err := mp4.DecodeFile(file)
	if err != nil {
		fmt.Printf("解析 MP4 文件错误: %v\n", err)
		return
	}

	// 打印 MP4 文件的所有元数据
	printBox(parsedMp4.Moov, 0)
}

// 打印盒子及其子盒子的信息
func printBox(box mp4.Box, level int) {
	indent := getIndent(level)
	boxType := box.Type()
	fmt.Printf("%s盒子类型: %s\n", indent, boxType)

	// 如果是 ContainerBox，递归打印子盒子
	if container, ok := box.(mp4.ContainerBox); ok {
		for _, child := range container.GetChildren() {
			printBox(child, level+1)
		}
	} else {
		printLeafBox(box, level)
	}
}

// 打印叶子盒子的详细信息
func printLeafBox(box mp4.Box, level int) {
	indent := getIndent(level)
	switch leafBox := box.(type) {
	case *mp4.DataBox:
		fmt.Printf("%s数据盒 (data): 数据长度 = %d, 数据内容 = %s\n", indent, len(leafBox.Data), string(leafBox.Data))
	default:
		fmt.Printf("%s未知类型盒子的详细信息无法解析\n", indent)
	}
}

// 生成缩进字符串
func getIndent(level int) string {
	return fmt.Sprintf("%s", fmt.Sprintf("%*s", level*2, ""))
}

// BoxData 用于存储盒子的基本信息和子盒子
type BoxData struct {
	Type     string     `json:"type"`
	Size     uint64     `json:"size"`
	Data     string     `json:"data,omitempty"`
	Children []*BoxData `json:"children,omitempty"`
}

var boxDataMap = make(map[string]BoxData)

func printJson() {
	// 打开 MP4 文件
	file, err := os.Open("/Users/a1234/Downloads/merged.mp4")
	if err != nil {
		fmt.Printf("打开文件错误: %v\n", err)
		return
	}
	defer file.Close()

	// 解析 MP4 文件
	parsedMp4, err := mp4.DecodeFile(file)
	if err != nil {
		fmt.Printf("解析 MP4 文件错误: %v\n", err)
		return
	}

	// 获取 MP4 文件的层次结构
	rootBoxData := handleBox(parsedMp4)

	// 将层次结构转换为 JSON
	jsonData, err := json.MarshalIndent(rootBoxData, "", "  ")
	if err != nil {
		fmt.Printf("JSON 序列化错误: %v\n", err)
		return
	}

	// 打印 JSON 数据
	fmt.Println(string(jsonData))
	//descBox := boxDataMap["desc"]
	//bytes, _ := json.Marshal(descBox)
	//fmt.Printf(string(bytes))
	//for _, box := range descBox.Children {
	//	if box.Type == "data" {
	//		fmt.Printf(box.Data)
	//	}
	//}
	//bytes2, _ := json.Marshal(boxDataMap)
	//fmt.Printf(string(bytes2))
	fmt.Printf(GetDataByType("desc", rootBoxData))
}

func GetDataByType(boxType string, rootBox *BoxData) string {
	result := make(map[string]string)

	var findFun func(box *BoxData)
	findFun = func(box *BoxData) {
		if box.Type == boxType {
			for _, child := range box.Children {
				result[child.Type] = child.Data
			}
			return
		}

		for _, child := range box.Children {
			findFun(child)
		}
	}

	findFun(rootBox)
	return result["data"]
}

func handleBox(mp4file *mp4.File) *BoxData {
	boxData := &BoxData{}
	for _, box := range mp4file.Moov.Children {
		boxType := box.Type()
		switch boxType {
		case "udta":
			boxData = getBoxData(box.(*mp4.UdtaBox))
		default:
			fmt.Printf("跳过类型为: %s 的盒子\n", boxType)
		}
	}
	return boxData
}

// 获取盒子及其子盒子的详细信息
func getBoxData(box mp4.Box) *BoxData {
	boxData := &BoxData{
		Type: box.Type(),
		Size: box.Size(),
	}

	// 如果是 ContainerBox，递归获取子盒子的信息
	if container, ok := box.(mp4.ContainerBox); ok {
		for _, child := range container.GetChildren() {
			childData := getBoxData(child)
			boxData.Children = append(boxData.Children, childData)
		}
	} else {
		boxData.Data = getBoxDataContent(box)
	}
	return boxData
}

// 获取叶子盒子的内容
func getBoxDataContent(box mp4.Box) string {
	switch leafBox := box.(type) {
	case *mp4.DataBox:
		return fmt.Sprintf("%s", leafBox.Data)
	default:
		return ""
	}
}

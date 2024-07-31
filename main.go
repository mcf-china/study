package main

import (
	"fmt"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

func main() {
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

//
//package main
//
//import (
//	"bytes"
//	"encoding/binary"
//	"fmt"
//	"io"
//	"os"
//	"path/filepath"
//)
//
//// BoxHeader 信息头
//type BoxHeader struct {
//	Size       uint32
//	FourccType [4]byte
//	Size64     uint64
//}
//
//func main() {
//	file, err := os.Open("/Users/a1234/Downloads/merged.mp4")
//	if err != nil {
//		panic(err)
//	}
//	duration, err := GetMP4Duration(file)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(filepath.Base("/Users/a1234/Downloads/merged.mp4"), duration)
//	var info = make([]byte, 0x100)
//	file.ReadAt(info, 0)
//	fmt.Println(info)
//}
//
//// GetMP4Duration 获取视频时长，以秒计
//func GetMP4Duration(reader io.ReaderAt) (lengthOfTime uint32, err error) {
//	var info = make([]byte, 0x10)
//	var boxHeader BoxHeader
//	var offset int64 = 0
//	// 获取moov结构偏移
//	for {
//		_, err = reader.ReadAt(info, offset)
//		if err != nil {
//			return
//		}
//		boxHeader = getHeaderBoxInfo(info)
//		fourccType := getFourccType(boxHeader)
//		if fourccType == "moov" {
//			break
//		}
//		// 有一部分mp4 mdat尺寸过大需要特殊处理
//		if fourccType == "mdat" {
//			if boxHeader.Size == 1 {
//				offset += int64(boxHeader.Size64)
//				continue
//			}
//		}
//		offset += int64(boxHeader.Size)
//	}
//	// 获取moov结构开头一部分
//	moovStartBytes := make([]byte, 0x100)
//	_, err = reader.ReadAt(moovStartBytes, offset)
//	if err != nil {
//		return
//	}
//	// 定义timeScale与Duration偏移
//	timeScaleOffset := 0x1C
//	durationOffest := 0x20
//	timeScale := binary.BigEndian.Uint32(moovStartBytes[timeScaleOffset : timeScaleOffset+4])
//	Duration := binary.BigEndian.Uint32(moovStartBytes[durationOffest : durationOffest+4])
//	lengthOfTime = Duration / timeScale
//	return
//}
//
//// getHeaderBoxInfo 获取头信息
//func getHeaderBoxInfo(data []byte) (boxHeader BoxHeader) {
//	buf := bytes.NewBuffer(data)
//	binary.Read(buf, binary.BigEndian, &boxHeader)
//	return
//}
//
//// getFourccType 获取信息头类型
//func getFourccType(boxHeader BoxHeader) (fourccType string) {
//	fourccType = string(boxHeader.FourccType[:])
//	return
//}

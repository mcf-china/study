package main

import (
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	text := "下单【2024-07-22 12:09:33】\n老板排队到了【2024-07-22 12:09:33】\n老板进房订单开始【2024-07-22 12:09:33】\n老板主动结束【2024-07-22 12:09:36】\n订单结算，原始金额(未计算分成):500，结算子订单数量:1【2024-07-22 12:12:36】"
	println(text)
	replace := strings.ReplaceAll(text, "\n", "<br/>")
	println(replace)
}

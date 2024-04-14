package macts

import (
	"testing"
)

func TestFgcolor(t *testing.T) {
	strBlack := "测试文字颜色: 黑色 Black"
	strRed := "测试文字颜色: 红色 Red"
	strGreen := "测试文字颜色: 绿色 Green"
	strYellow := "测试文字颜色: 黄色 Yellow"
	strBlue := "测试文字颜色: 蓝色 Blue"
	strMagenta := "测试文字颜色: 品红 Magenta"
	strCyan := "测试文字颜色: 青色 Cyan"
	strWhite := "测试文字颜色: 白色 White"

	black := FgColor.Black(strBlack)
	red := FgColor.Red(strRed)
	green := FgColor.Green(strGreen)
	yellow := FgColor.Yellow(strYellow)
	blue := FgColor.Blue(strBlue)
	magenta := FgColor.Magenta(strMagenta)
	cyan := FgColor.Cyan(strCyan)
	white := FgColor.White(strWhite)

	if black != "\u001B[30m测试文字颜色: 黑色 Black\u001B[0m" {
		t.Error("黑色 Black 测试失败")
	} else {
		t.Log(black)
	}
	if red != "\u001B[31m测试文字颜色: 红色 Red\u001B[0m" {
		t.Error("红色 Red 测试失败")
	} else {
		t.Log(red)
	}
	if green != "\u001B[32m测试文字颜色: 绿色 Green\u001B[0m" {
		t.Error("绿色 Green 测试失败")
	} else {
		t.Log(green)
	}
	if yellow != "\u001B[33m测试文字颜色: 黄色 Yellow\u001B[0m" {
		t.Error("黄色 Yellow 测试失败")
	} else {
		t.Log(yellow)
	}
	if blue != "\u001B[34m测试文字颜色: 蓝色 Blue\u001B[0m" {
		t.Error("蓝色 Blue 测试失败")
	} else {
		t.Log(blue)
	}
	if magenta != "\u001B[35m测试文字颜色: 品红 Magenta\u001B[0m" {
		t.Error("品红 Magenta 测试失败")
	} else {
		t.Log(magenta)
	}
	if cyan != "\u001B[36m测试文字颜色: 青色 Cyan\u001B[0m" {
		t.Error("青色 Cyan 测试失败")
	} else {
		t.Log(cyan)
	}
	if white != "\u001B[37m测试文字颜色: 白色 White\u001B[0m" {
		t.Error("白色 White 测试失败")
	} else {
		t.Log(white)
	}
}

func TestBgcolor(t *testing.T) {
	strBlack := "测试背景颜色: 黑色 Black"
	strRed := "测试背景颜色: 红色 Red"
	strGreen := "测试背景颜色: 绿色 Green"
	strYellow := "测试背景颜色: 黄色 Yellow"
	strBlue := "测试背景颜色: 蓝色 Blue"
	strMagenta := "测试背景颜色: 品红 Magenta"
	strCyan := "测试背景颜色: 青色 Cyan"
	strWhite := "测试背景颜色: 白色 White"

	black := BgColor.Black(strBlack)
	red := BgColor.Red(strRed)
	green := BgColor.Green(strGreen)
	yellow := BgColor.Yellow(strYellow)
	blue := BgColor.Blue(strBlue)
	magenta := BgColor.Magenta(strMagenta)
	cyan := BgColor.Cyan(strCyan)
	white := BgColor.White(strWhite)

	if black != "\u001B[40m测试背景颜色: 黑色 Black\u001B[0m" {
		t.Error("黑色 Black 测试失败")
	} else {
		t.Log(black)
	}
	if red != "\u001B[41m测试背景颜色: 红色 Red\u001B[0m" {
		t.Error("红色 Red 测试失败")
	} else {
		t.Log(red)
	}
	if green != "\u001B[42m测试背景颜色: 绿色 Green\u001B[0m" {
		t.Error("绿色 Green 测试失败")
	} else {
		t.Log(green)
	}
	if yellow != "\u001B[43m测试背景颜色: 黄色 Yellow\u001B[0m" {
		t.Error("黄色 Yellow 测试失败")
	} else {
		t.Log(yellow)
	}
	if blue != "\u001B[44m测试背景颜色: 蓝色 Blue\u001B[0m" {
		t.Error("蓝色 Blue 测试失败")
	} else {
		t.Log(blue)
	}
	if magenta != "\u001B[45m测试背景颜色: 品红 Magenta\u001B[0m" {
		t.Error("品红 Magenta 测试失败")
	} else {
		t.Log(magenta)
	}
	if cyan != "\u001B[46m测试背景颜色: 青色 Cyan\u001B[0m" {
		t.Error("青色 Cyan 测试失败")
	} else {
		t.Log(cyan)
	}
	if white != "\u001B[47m测试背景颜色: 白色 White\u001B[0m" {
		t.Error("白色 White 测试失败")
	} else {
		t.Log(white)
	}
}
func TestBgColorBright(t *testing.T) {

	strBlack := "测试背景高亮颜色: 黑色 Black"
	strRed := "测试背景高亮颜色: 红色 Red"
	strGreen := "测试背景高亮颜色: 绿色 Green"
	strYellow := "测试背景高亮颜色: 黄色 Yellow"
	strBlue := "测试背景高亮颜色: 蓝色 Blue"
	strMagenta := "测试背景高亮颜色: 品红 Magenta"
	strCyan := "测试背景高亮颜色: 青色 Cyan"
	strWhite := "测试背景高亮颜色: 白色 White"

	black := BgColorBright.Black(strBlack)
	red := BgColorBright.Red(strRed)
	green := BgColorBright.Green(strGreen)
	yellow := BgColorBright.Yellow(strYellow)
	blue := BgColorBright.Blue(strBlue)
	magenta := BgColorBright.Magenta(strMagenta)
	cyan := BgColorBright.Cyan(strCyan)
	white := BgColorBright.White(strWhite)

	if black != "\u001B[100m测试背景高亮颜色: 黑色 Black\u001B[0m" {
		t.Error("黑色 Black 测试失败")
	} else {
		t.Log(black)
	}
	if red != "\u001B[101m测试背景高亮颜色: 红色 Red\u001B[0m" {
		t.Error("红色 Red 测试失败")
	} else {
		t.Log(red)
	}
	if green != "\u001B[102m测试背景高亮颜色: 绿色 Green\u001B[0m" {
		t.Error("绿色 Green 测试失败")
	} else {
		t.Log(green)
	}
	if yellow != "\u001B[103m测试背景高亮颜色: 黄色 Yellow\u001B[0m" {
		t.Error("黄色 Yellow 测试失败")
	} else {
		t.Log(yellow)
	}
	if blue != "\u001B[104m测试背景高亮颜色: 蓝色 Blue\u001B[0m" {
		t.Error("蓝色 Blue 测试失败")
	} else {
		t.Log(blue)
	}
	if magenta != "\u001B[105m测试背景高亮颜色: 品红 Magenta\u001B[0m" {
		t.Error("品红 Magenta 测试失败")
	} else {
		t.Log(magenta)
	}
	if cyan != "\u001B[106m测试背景高亮颜色: 青色 Cyan\u001B[0m" {
		t.Error("青色 Cyan 测试失败")
	} else {
		t.Log(cyan)
	}
	if white != "\u001B[107m测试背景高亮颜色: 白色 White\u001B[0m" {
		t.Error("白色 White 测试失败")
	} else {
		t.Log(white)
	}
}

func TestStyle(t *testing.T) {

	strBold := "测试字体样式: 加粗 Bold"
	strDownplay := "测试字体样式: 淡化 Downplay"
	strItalic := "测试字体样式: 斜体 Italic"
	strUnderline := "测试字体样式: 下划线 Underline"
	strReverse := "测试字体样式: 反转 Reverse"

	Bold := Style.Bold(strBold)
	Downplay := Style.Downplay(strDownplay)
	Italic := Style.Italic(strItalic)
	Underline := Style.Underline(strUnderline)
	Reverse := Style.Reverse(strReverse)

	if Bold != "\u001B[1m测试字体样式: 加粗 Bold\u001B[0m" {
		t.Error("加粗 Bold 测试失败")
	} else {
		t.Log(Bold)
	}
	if Downplay != "\u001B[2m测试字体样式: 淡化 Downplay\u001B[0m" {
		t.Error("淡化 Downplay 测试失败")
	} else {
		t.Log(Downplay)
	}
	if Italic != "\u001B[3m测试字体样式: 斜体 Italic\u001B[0m" {
		t.Error("斜体 Italic 测试失败")
	} else {
		t.Log(Italic)
	}
	if Underline != "\u001B[4m测试字体样式: 下划线 Underline\u001B[0m" {
		t.Error("下划线 Underline 测试失败")
	} else {
		t.Log(Underline)
	}
	if Reverse != "\u001B[7m测试字体样式: 反转 Reverse\u001B[0m" {
		t.Error("反转 Reverse 测试失败")
	} else {
		t.Log(Reverse)
	}

}
func TestRender(t *testing.T) {

}

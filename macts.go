package macts

import (
	"fmt"
)

// fmt.Println("\u001B[0;31;1m" + "颜色测试" + "\u001B[0m") // 35
// fmt.Println("\u001B[1;31;1m" + "颜色测试" + "\u001B[0m") // 35
// fmt.Println("\u001B[2;31;1m" + "颜色测试" + "\u001B[0m") // 35
// fmt.Println("\u001B[3;31;1m" + "颜色测试" + "\u001B[0m") // 35
// fmt.Println("\u001B[4;31;1m" + "颜色测试" + "\u001B[0m") // 35
// fmt.Println("\u001B[5;31;1m" + "颜色测试" + "\u001B[0m") // 35

var (

	/*
		FgColor 字体颜色修改函数
			- FgColor.Black() 黑色
			- FgColor.Red() 红色
			- FgColor.Green() 绿色
			- FgColor.Yellow() 黄色
			- FgColor.Blue() 蓝色
			- FgColor.Magenta() 品红
			- FgColor.Cyan() 青色
			- FgColor.White() 白色
	*/
	FgColor struct {
		Black   func(string) string
		Red     func(string) string
		Green   func(string) string
		Yellow  func(string) string
		Blue    func(string) string
		Magenta func(string) string
		Cyan    func(string) string
		White   func(string) string
	}

	/*
		BgColor 字体背景修改函数
			- BgColor.Black() 黑色
			- BgColor.Red() 红色
			- BgColor.Green() 绿色
			- BgColor.Yellow() 黄色
			- BgColor.Blue() 蓝色
			- BgColor.Magenta() 品红
			- BgColor.Cyan() 青色
			- BgColor.White() 白色
	*/
	BgColor struct {
		Black   func(string) string
		Red     func(string) string
		Green   func(string) string
		Yellow  func(string) string
		Blue    func(string) string
		Magenta func(string) string
		Cyan    func(string) string
		White   func(string) string
	}

	/*
		BgColorBright 高亮背景颜色

			- BgColorBright.Black() 黑色高亮/灰色
			- BgColorBright.Red() 红色高亮
			- BgColorBright.Green() 绿色高亮
			- BgColorBright.Yellow() 黄色高亮
			- BgColorBright.Blue() 蓝色高亮
			- BgColorBright.Magenta() 品红高亮
			- BgColorBright.Cyan() 青色高亮
			- BgColorBright.White() 白色高亮
	*/
	BgColorBright struct {
		Black   func(string) string
		Red     func(string) string
		Green   func(string) string
		Yellow  func(string) string
		Blue    func(string) string
		Magenta func(string) string
		Cyan    func(string) string
		White   func(string) string
	}

	/*
		Style 字体样式调整
			- Style.Bold() 加粗
			- Style.Downplay() 淡化
			- Style.Italic() 倾斜
			- Style.Underline() 下划线
			- Style.Reverse() 反转
	*/
	Style struct {
		Bold      func(string) string
		Downplay  func(string) string
		Italic    func(string) string
		Underline func(string) string
		Reverse   func(string) string
	}
)

func init() {
	//  字体颜色调整
	FgColor.Black = func(s string) string {
		return fmt.Sprintf("\u001B[30m%s\u001B[0m", s)
	}
	FgColor.Red = func(s string) string {
		return fmt.Sprintf("\u001B[31m%s\u001B[0m", s)
	}
	FgColor.Green = func(s string) string {
		return fmt.Sprintf("\u001B[32m%s\u001B[0m", s)
	}
	FgColor.Yellow = func(s string) string {
		return fmt.Sprintf("\u001B[33m%s\u001B[0m", s)
	}
	FgColor.Blue = func(s string) string {
		return fmt.Sprintf("\u001B[34m%s\u001B[0m", s)
	}
	FgColor.Magenta = func(s string) string {
		return fmt.Sprintf("\u001B[35m%s\u001B[0m", s)
	}
	FgColor.Cyan = func(s string) string {
		return fmt.Sprintf("\u001B[36m%s\u001B[0m", s)
	}
	FgColor.White = func(s string) string {
		return fmt.Sprintf("\u001B[37m%s\u001B[0m", s)
	}

	// 字体背景调整
	BgColor.Black = func(s string) string {
		return fmt.Sprintf("\u001B[40m%s\u001B[0m", s)
	}
	BgColor.Red = func(s string) string {
		return fmt.Sprintf("\u001B[41m%s\u001B[0m", s)
	}
	BgColor.Green = func(s string) string {
		return fmt.Sprintf("\u001B[42m%s\u001B[0m", s)
	}
	BgColor.Yellow = func(s string) string {
		return fmt.Sprintf("\u001B[43m%s\u001B[0m", s)
	}
	BgColor.Blue = func(s string) string {
		return fmt.Sprintf("\u001B[44m%s\u001B[0m", s)
	}
	BgColor.Magenta = func(s string) string {
		return fmt.Sprintf("\u001B[45m%s\u001B[0m", s)
	}
	BgColor.Cyan = func(s string) string {
		return fmt.Sprintf("\u001B[46m%s\u001B[0m", s)
	}
	BgColor.White = func(s string) string {
		return fmt.Sprintf("\u001B[47m%s\u001B[0m", s)
	}

	// 字体高亮背景
	BgColorBright.Black = func(s string) string {
		return fmt.Sprintf("\u001B[100m%s\u001B[0m", s)
	}
	BgColorBright.Red = func(s string) string {
		return fmt.Sprintf("\u001B[101m%s\u001B[0m", s)
	}
	BgColorBright.Green = func(s string) string {
		return fmt.Sprintf("\u001B[102m%s\u001B[0m", s)
	}
	BgColorBright.Yellow = func(s string) string {
		return fmt.Sprintf("\u001B[103m%s\u001B[0m", s)
	}
	BgColorBright.Blue = func(s string) string {
		return fmt.Sprintf("\u001B[104m%s\u001B[0m", s)
	}
	BgColorBright.Magenta = func(s string) string {
		return fmt.Sprintf("\u001B[105m%s\u001B[0m", s)
	}
	BgColorBright.Cyan = func(s string) string {
		return fmt.Sprintf("\u001B[106m%s\u001B[0m", s)
	}
	BgColorBright.White = func(s string) string {
		return fmt.Sprintf("\u001B[107m%s\u001B[0m", s)
	}

	// 字体样式设置

	Style.Bold = func(s string) string {
		return fmt.Sprintf("\u001B[1m%s\u001B[0m", s)
	}
	Style.Downplay = func(s string) string {
		return fmt.Sprintf("\u001B[2m%s\u001B[0m", s)
	}
	Style.Italic = func(s string) string {
		return fmt.Sprintf("\u001B[3m%s\u001B[0m", s)
	}
	Style.Underline = func(s string) string {
		return fmt.Sprintf("\u001B[4m%s\u001B[0m", s)
	}
	Style.Reverse = func(s string) string {
		return fmt.Sprintf("\u001B[7m%s\u001B[0m", s)
	}
}

/*
Render 统一的渲染器，可实现所有的渲染效果

# 参数：FontStyle(string) 字体样式
  - "Bold" 加粗
  - "Downplay" 减淡
  - "Italic" 倾斜
  - "Underline" 下划线
  - "Reverse" 颜色反转

# 参数：ForeGroundColor/BackGroundColor(string) 字体样式/背景颜色
  - "Black" 黑色
  - "Red" 红色
  - "Green" 绿色
  - "Yellow" 黄色
  - "Blue" 蓝色
  - "Magenta" 品红
  - "Cyan" 青色
  - "White" 白色

# 参数：BackGroundBright(int) 背景高亮
  - 0 默认 不高亮
  - 1 高亮背景
*/
func Render(s string, FontStyle string, ForeGroundColor string, BackGroundColor string, BackGroundBright int) string {

	st, fgc, bgc := 0, 0, 0

	// fmt.Println(FontStyle, ForeGroundColor, BackGroundColor)
	// 样式
	if FontStyle == "Bold" {
		st = 1
	} else if FontStyle == "Downplay" {
		st = 2
	} else if FontStyle == "Italic" {
		st = 3
	} else if FontStyle == "Underline" {
		st = 4
	} else if FontStyle == "Reverse" {
		st = 7
	} else {
		st = 0
	}

	// 字体颜色
	if ForeGroundColor == "Black" {
		fgc = 30
	} else if ForeGroundColor == "Red" {
		fgc = 31
	} else if ForeGroundColor == "Green" {
		fgc = 32
	} else if ForeGroundColor == "Yellow" {
		fgc = 33
	} else if ForeGroundColor == "Blue" {
		fgc = 34
	} else if ForeGroundColor == "Magenta" {
		fgc = 35
	} else if ForeGroundColor == "Cyan" {
		fgc = 36
	} else if ForeGroundColor == "White" {
		fgc = 37
	} else {
		fgc = 20 // 默认值 这个值下能正常输出 但无法确定是否会有影响
	}

	// 背景色
	if BackGroundColor == "Black" {
		bgc = 40
	} else if BackGroundColor == "Red" {
		bgc = 41
	} else if BackGroundColor == "Green" {
		bgc = 42
	} else if BackGroundColor == "Yellow" {
		bgc = 43
	} else if BackGroundColor == "Blue" {
		bgc = 44
	} else if BackGroundColor == "Magenta" {
		bgc = 45
	} else if BackGroundColor == "Cyan" {
		bgc = 46
	} else if BackGroundColor == "White" {
		bgc = 47
	} else {
		bgc = 20 // 默认值 这个值下能正常输出 但无法确定是否会有影响 加粗情况下会变成80
	}

	// 高亮背景色
	if BackGroundBright != 0 {
		bgc += 60
	}
	// fmt.Println(fmt.Sprintf("[%d;%d;%dm", st, fgc, bgc))

	return fmt.Sprintf("\u001B[%d;%d;%dm%s\u001B[0m", st, fgc, bgc, s)

}

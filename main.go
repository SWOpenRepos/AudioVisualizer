package main

import (
	"fyne.io/fyne/v2/app"

	"SWOpen/AudioVisualizer/internal/window"
)

func main() {
	// 创建应用
	a := app.NewWithID("wiki.simulink.open.AudioVisualizer")
	// 初始化主窗口
	window.InitWindow(a)
}

package window

import (
	"github.com/webview/webview_go"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"SWOpen/AudioVisualizer/internal/server"
)

var (
	w_main fyne.Window
)

func InitWindow(a fyne.App) {
	w_main = a.NewWindow("音频驱动电控悬架 - CarSim.Net")
	w_main.Resize(fyne.NewSize(460, 180))

	startBtn1 := widget.NewButtonWithIcon("使用WebView打开", theme.MediaPlayIcon(), func() {
		w := webview.New(false)
		defer w.Destroy()
		w.SetTitle("音频驱动电控悬架 - CarSim.Net")
		w.SetSize(460, 720, webview.HintNone)
		w.Navigate("http://localhost:1516/")
		w.Run()
	})
	startBtn1.Disable()
	startBtn2 := widget.NewButtonWithIcon("使用浏览器打开", theme.MediaPlayIcon(), func() {
		err := openURL("http://localhost:1516/")
		if err != nil {
			dialog.ShowError(err, w_main)
		}
	})
	startBtn2.Disable()
	webSiteBtn := widget.NewButton("CarSim.Net官网", func() {
		err := openURL("https://carsim.net/")
		if err != nil {
			dialog.ShowError(err, w_main)
		}
	})
	c_v1_h1 := container.NewHBox(
		startBtn1,
		widget.NewSeparator(),
		startBtn2,
	)
	c_v1_h2 := container.NewHBox(
		webSiteBtn,
	)
	c_v1 := container.New(layout.NewHBoxLayout(),
		c_v1_h1,
		layout.NewSpacer(),
		c_v1_h2,
	)

	logEntry := widget.NewEntry()
	status := server.Start()
	if status == 1 {
		logEntry.SetText("后台服务启动失败")
	} else if status == 2 {
		logEntry.SetText("请将control.js和control.wasm文件放到程序目录下，重新启动程序")
	} else {
		logEntry.SetText("后台服务启动成功")
		startBtn1.Enable()
		startBtn2.Enable()
	}
	c_v2_v1 := container.NewScroll(logEntry)
	c_v2_v1.SetMinSize(fyne.NewSize(0, 100))
	c_v2 := container.NewVBox(
		c_v2_v1,
	)
	mainContent := container.New(layout.NewVBoxLayout(),
		container.NewVBox(
			c_v1,
			widget.NewSeparator(),
			c_v2,
		),
	)

	w_main.SetContent(mainContent)
	w_main.Show()
	a.Run()
}

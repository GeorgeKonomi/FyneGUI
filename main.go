package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(1000, 700)
	window.SetWindowTitle("Hello Widgets Example")

	box := widgets.NewQGroupBox(nil)

	box2 := widgets.NewQGroupBox(nil)

	lineEdit1 := widgets.NewQLineEdit(box)
	lineEdit1.SetPlaceholderText("Main Panel...")

	textEdit1 := widgets.NewQTextEdit(box2)
	textEdit1.SetMinimumWidth(2 * box2.Width())
	textEdit1.SetMinimumHeight(box2.Height() - 100)
	textEdit1.AdjustSize()
	textEdit1.SetTextBackgroundColor(gui.NewQColor3(100, 25, 100, 200))
	textEdit1.SetTextColor(gui.NewQColor3(100, 200, 100, 200))
	textEdit1.SetPlaceholderText("Write text here...")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	widget.Layout().AddWidget(box)
	widget.Layout().AddWidget(box2)
	createButtons(widget, textEdit1)

	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

// func helper() {

// 	view := widgets.NewQGraphicsView(nil)
// 	scene := widgets.NewQGraphicsScene3(0, 0, 300, 300, view)
// 	view.SetScene(scene)

// 	font := gui.NewQFont2("Meiryo", 20, 2, false)
// 	text := scene.AddText("Hello World", font)
// 	text.SetPos2((scene.Width()-text.BoundingRect().Width())/2, (scene.Height()-text.BoundingRect().Height())/2)

// 	pen := gui.NewQPen3(gui.NewQColor3(255, 0, 0, 255))
// 	scene.AddLine2(20, scene.Height()*0.75, scene.Width()-20, scene.Height()*0.75, pen)

// 	view.Pointer()
// }

func createButtons(widget *widgets.QWidget, text *widgets.QTextEdit) {
	// create a line edit
	// with a custom placeholder text
	// and add it to the central widgets layout
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	// create a button
	// connect the clicked signal
	// and add it to the central widgets layout
	buttonModemPanelTest := widgets.NewQPushButton2("Run modem panel navigation test", nil)
	buttonModemPanelTest.ConnectClicked(func(bool) {
		/////
		str, err := ExecuteCmd("modemPanelTest")

		if err != nil {
			fmt.Println("exoume error")
		} else {
			fmt.Println(str + " den exoume error")
		}
		/////
		text.SetText(str)
		widgets.QMessageBox_Information(nil, "OK", str, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(buttonModemPanelTest)

	buttonEmulationTest := widgets.NewQPushButton2("Run emulation test", nil)
	buttonEmulationTest.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(buttonEmulationTest)

	buttonZerotierConnectivityTest := widgets.NewQPushButton2("Run zerotier connectivity test", nil)
	buttonZerotierConnectivityTest.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(buttonZerotierConnectivityTest)

}

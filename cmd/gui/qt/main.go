package main

import (
	"os"

	//"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	//1. Создайте новое приложение
	app := widgets.NewQApplication(len(os.Args), os.Args)

	//2. Создайте главное окно
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(600, 300)
	window.SetWindowTitle("Hello Deploy Example")

	//3. Создайте макет ⇒ https://doc.qt.io/qt-5/qlayout.html
	layout := widgets.NewQHBoxLayout()

	//4. Создайте виджет и назначьте макет для виджета
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	//5-1. Создайте строку ввода
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Напиши здесь что-нибудь!")

	//5-2. Создайте кнопку
	button := widgets.NewQPushButton2("Кликай сюда", nil)
	button.ConnectClicked(func(checked bool) {
		//6. Назначьте действие на кнопку, которое должно
		//   выполняться при её нажатии
		widgets.QMessageBox_Information(nil, "Заголовок информ. окна", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	//7.  Назначьте виджеты ввода макету
	layout.AddWidget(input, 0, 0)  // ⇒ http://doc.qt.io/qt-5/qboxlayout.html#addWidget
	layout.AddWidget(button, 0, 0) // ⇒ https://doc.qt.io/qt-5/qboxlayout.html#addWidget

	//8. Главный виджет должен быть назначен главному окну
	window.SetCentralWidget(widget)

	//9. Выведите главное окно
	window.Show()

	//10. Чтобы приложение могло динамически реагировать на пользовательский ввод/события,
	//    оно должно работать в бесконечном цикле.
	//widgets.QApplication_Exec()

	app.Exec()

	//widget := widgets.NewQWidget(nil, 0)
	//widget.SetLayout(widgets.NewQVBoxLayout())
	////widget.Layout().QLayoutItem.SetAlignment(core.Qt__AlignCenter)
	//window.SetCentralWidget(widget)
	//for _, d := range [][]string{
	//	{"https://github.com/therecipe/qt/wiki/Setting-the-Application-Icon#windows", "Docs"},
	//	{"https://doc.qt.io/qt-5/windows-deployment.html", "Qt docs"},
	//	{"https://www.iconfinder.com/icons/52510/application_icon", "Icon credits"},
	//} {
	//	label := widgets.NewQLabel2(fmt.Sprintf("<a href=\"%v\">%v</a>", d[0], d[1]), nil, 0)
	//	label.SetOpenExternalLinks(true)
	//	widget.Layout().AddWidget(label)
	//}

}

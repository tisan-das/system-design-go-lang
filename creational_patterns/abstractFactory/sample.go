package abstractFactory

type IWidgetFactory interface {
	CreateWindow() IWindow
	CreateButton() IButton
	CreateMenu() IMenu
}

type WinWidgetFactory struct{}

func NewWinWidgetFactory() IWidgetFactory {
	return &WinWidgetFactory{}
}

func (factory *WinWidgetFactory) CreateWindow() IWindow {
	return &WinWindow{WindowsWindow: "windows window 1"}
}
func (factory *WinWidgetFactory) CreateButton() IButton {
	return &WinButton{WindowsButton: "windows button 1"}
}
func (factory *WinWidgetFactory) CreateMenu() IMenu {
	return &WinMenu{WindowsMenu: "windows menu 1"}
}

type MacWidgetFactory struct{}

func NewMacWidgetFactory() IWidgetFactory {
	return &MacWidgetFactory{}
}

func (factory *MacWidgetFactory) CreateWindow() IWindow {
	return &MacWindow{MacWindow: "macintosh window 1"}
}
func (factory *MacWidgetFactory) CreateButton() IButton {
	return &MacButton{MacButton: "macintosh button 1"}
}
func (factory *MacWidgetFactory) CreateMenu() IMenu {
	return &MacMenu{MacMenu: "macintosh menu 1"}
}

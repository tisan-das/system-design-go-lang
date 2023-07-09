package abstractFactory

type IWindow interface {
	GetWindow() IWindow
}

type IButton interface {
	GetButton() IButton
}

type IMenu interface {
	GetMenu() IMenu
}

type WinWindow struct {
	WindowsWindow string
}

func (window *WinWindow) GetWindow() IWindow {
	return &WinWindow{WindowsWindow: "windows window"}
}

type MacWindow struct {
	MacWindow string
}

func (window *MacWindow) GetWindow() IWindow {
	return &MacWindow{MacWindow: "mac window"}
}

type WinButton struct {
	WindowsButton string
}

func (button *WinButton) GetButton() IButton {
	return &WinButton{WindowsButton: "windows Button"}
}

type MacButton struct {
	MacButton string
}

func (button *MacButton) GetButton() IButton {
	return &MacButton{MacButton: "mac button"}
}

type WinMenu struct {
	WindowsMenu string
}

func (menu *WinMenu) GetMenu() IMenu {
	return &WinMenu{WindowsMenu: "windows menu"}
}

type MacMenu struct {
	MacMenu string
}

func (menu *MacMenu) GetMenu() IMenu {
	return &MacMenu{MacMenu: "mac menu"}
}

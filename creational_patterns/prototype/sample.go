package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Prototype interface {
	Clone() Prototype
	Print()
}

type Window struct {
	WindowName string
	OpenedFile *File
}

func NewWindowPrototype() Prototype {
	return &Window{WindowName: "window1", OpenedFile: &File{Name: "file1"}}
}

func (window *Window) Clone() Prototype {
	return window.shallowCopy()
	// return window.DeepCopy()
}

// ---- ShallowCopy
func (window *Window) shallowCopy() Prototype {
	fmt.Println("---> Going for a shallow copy!")
	return &Window{
		WindowName: window.WindowName,
		OpenedFile: window.OpenedFile,
	}
}

// --- Deep copy: Copy through serialization
func (window *Window) DeepCopy() Prototype {
	fmt.Println("---> Going for a deep copy!")
	buf := bytes.Buffer{}

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(window)
	fmt.Println("Error occurred while encoding: ", err)

	var newWindow Window
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&newWindow)
	fmt.Println("Error occurred while decoding: ", err)
	return &newWindow
}

func (window *Window) Print() {
	fmt.Printf("Memory loc: window:%p windowName:%p, openedFile:%p\n",
		window, window.WindowName, window.OpenedFile)
	fmt.Printf("Value: window:%+v windowName:%+v, openedFile:%+v\n",
		window, window.WindowName, window.OpenedFile)
}

type File struct {
	Name string
}

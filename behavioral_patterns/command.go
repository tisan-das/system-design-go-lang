package behavioral_patterns

import "fmt"

type TvDevice struct {
	state bool
}

// Invoker
type Button struct {
	toggleState Command
}

func (button *Button) press() {
	button.toggleState.Execute()
}

// Command interface
type Command interface {
	Execute()
}

// TODO: how to have constructor??
type RemoteController struct {
	tvDevice *TvDevice
}

func (remoteCon *RemoteController) SetDevice(tvDevice *TvDevice) {
	remoteCon.tvDevice = tvDevice
}

func (remoteCon *RemoteController) Execute() {
	remoteCon.tvDevice.state = !remoteCon.tvDevice.state
	fmt.Println("TV state flipped to: ", remoteCon.tvDevice.state)
}

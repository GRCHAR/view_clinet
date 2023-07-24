package operate

import (
	"github.com/go-vgo/robotgo"
)

const (
	// Version get the gohook version
	Version = "v0.31.2.113, Sierra Nevada!"

	// HookEnabled honk enable status
	HookEnabled  = 1 // iota
	HookDisabled = 2

	KeyDown = 3
	KeyHold = 4
	KeyUp   = 5

	MouseUp    = 6
	MouseHold  = 7
	MouseDown  = 8
	MouseMove  = 9
	MouseDrag  = 10
	MouseWheel = 11

	FakeEvent = 12

	// Keychar could be v
	CharUndefined = 0xFFFF
	WheelUp       = -1
	WheelDown     = 1
)

var controller = newLocalOperateController()

func init() {

	go operateThread()
}

type localOperateController struct {
	localOperateChan chan operateStep
}

func newLocalOperateController() *localOperateController {
	return &localOperateController{
		localOperateChan: make(chan operateStep, 100000),
	}
}

func operateThread() {
	setMouseSleep(10)
	for {
		select {
		case step := <-operateController.operateChan:
			switch step.Event.Kind {
			case MouseMove:
				moveMouse(step.Event.X, step.Event.Y)
			case MouseDrag:

			}
		}
	}
}

func setMouseSleep(sleepTime int) {
	robotgo.MouseSleep = sleepTime
}

func moveMouse(x int16, y int16) {
	robotgo.MoveSmooth(int(x), int(y))
}

func dragMouse(x int16, y int16) {
	robotgo.DragSmooth(int(x), int(y))
}

func wheelMouse() {
	robotgo.ScrollSmooth(100)
}

func clickDownMouse(button int, clicks int) {
	switch button {
	case 1:
		if clicks >= 2 {
			robotgo.Click("left", true)
		} else {
			robotgo.Click("left", false)
		}
	case 2:
		if clicks >= 2 {
			robotgo.Click("right", true)
		} else {
			robotgo.Click("right", false)
		}
	}

}

func clickUpMouse(button int) {

}

func scrollMouse() {
	robotgo.ScrollSmooth(100)
}

func clickKeyboard() {

}

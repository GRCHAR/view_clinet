package operate

import (
	"changeme/src/logger"
	hook "github.com/robotn/gohook"
	"net"
)

func init() {
	//go add()
	go low()
}

func InitListenOperate() {

}

func add() {
	logger.GetLogger().Info("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		logger.GetLogger().Info("ctrl-shift-q")

	})

	logger.GetLogger().Info("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		logger.GetLogger().Info("w")
	})
	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()
	for ev := range evChan {
		//fmt.Println("hook: ", ev)

		step := operateStep{Event: ev, ClientKey: getMacAddr()}
		operateController.sendOperateChan <- step
	}

}

func create() {

}

func getMacAddr() string {
	macAddr := ""
	interfaces, err := net.Interfaces()
	if err != nil {
		logger.GetLogger().Error(err)
		return ""
	}
	for _, inter := range interfaces {
		macAddr += inter.HardwareAddr.String()
	}
	return macAddr
}

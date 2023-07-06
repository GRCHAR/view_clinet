package operate

import (
	"bytes"
	"changeme/src/logger"
	"encoding/json"
	hook "github.com/robotn/gohook"
	"golang.org/x/net/ipv6"
	"net"
	"sync"
)

type operateStep struct {
	ClientKey string
	Event     hook.Event
}

type operateControl struct {
	operateChan     chan operateStep
	sendOperateChan chan operateStep
	conn            net.Conn
}

var operateController = operateControl{
	operateChan:     make(chan operateStep, 10000),
	sendOperateChan: make(chan operateStep, 10000),
	conn:            nil,
}

func init() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go operateController.connectServer(&wg)
}

func (control *operateControl) connectServer(wg *sync.WaitGroup) {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		logger.GetLogger().Error(err)
		return
	}
	control.conn = conn
	control.conn.Write([]byte("test"))
	wg.Done()

	go func() {
		for {
			buffer := bytes.Buffer{}
			n, err := buffer.ReadFrom(control.conn)
			if err != nil {
				logger.GetLogger().Error(err)
				return
			}
			result := buffer.Bytes()
			logger.GetLogger().Info(string(buffer.Bytes()[:n]))
			data := operateStep{}
			err = json.Unmarshal(result, &data)
			if err != nil {
				logger.GetLogger().Error(err)
			}

			control.operateChan <- data
		}
	}()
	control.sendDataToRemoteClient(wg)
}

func (control *operateControl) sendDataToRemoteClient(wg *sync.WaitGroup) {
	wg.Wait()
	for {
		select {
		case step := <-control.sendOperateChan:
			selfStep := step
			logger.GetLogger().Info(step.ClientKey, step.Event.String())
			data, err := json.Marshal(&selfStep)
			logger.GetLogger().Info(string(data))
			localStep := operateStep{}
			json.Unmarshal(data, &localStep)

			if err != nil {
				logger.GetLogger().Error(err)
				continue
			}
			//data = append(data, byte(0xff))
			_, err = control.conn.Write(data)
			if err != nil {
				logger.GetLogger().Error(err)
				continue
			}
		}
	}
}

func (control *operateControl) p2pIPV6(ip6 string) {

	dial, err := net.Dial("ip6", ip6)
	if err != nil {
		return
	}
	conn := ipv6.NewConn(dial)
	for {
		select {
		case step := <-control.operateChan:
			message := &ipv6.Message{}
			stepBuffer, _ := json.Marshal(step)
			message.Buffers = append([][]byte{}, stepBuffer)
			conn.SendMsg(message, 0)
		}
	}
}

func listenLocalOperate() {

}

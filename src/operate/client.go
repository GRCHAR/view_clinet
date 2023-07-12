package operate

import (
	"changeme/src/logger"
	"encoding/json"
	hook "github.com/robotn/gohook"
	"golang.org/x/net/ipv6"
	"net"
	"sync"
	"time"
)

type operateStep struct {
	ClientKey string
	Event     hook.Event
}

type operateControl struct {
	operateChan     chan operateStep
	sendOperateChan chan operateStep
	conn            net.UDPConn
	tcpConn         net.Conn
}

var operateController = operateControl{
	operateChan:     make(chan operateStep, 10000),
	sendOperateChan: make(chan operateStep, 10000),
	conn:            net.UDPConn{},
}

func init() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go operateController.connectServer(&wg)
}

func (control *operateControl) connectServer(wg *sync.WaitGroup) {
	//tcpConn, err := net.Dial("tcp", "127.0.0.1:8002")
	//if err != nil {
	//	logger.GetLogger().Error(err)
	//	return
	//}
	//control.tcpConn = tcpConn
	//_, err = control.tcpConn.Write([]byte("test"))
	//if err != nil {
	//	logger.GetLogger().Info()
	//	return
	//}
	//124.70.94.103
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8001")
	if err != nil {
		logger.GetLogger().Error("服务连接失败:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		logger.GetLogger().Error("服务连接失败:", err)
		return
	}
	control.conn = *conn
	control.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	control.conn.Write([]byte("test"))
	wg.Done()

	go func() {

		for {
			buf := make([]byte, 1024)
			n, _, err := control.conn.ReadFromUDP(buf)

			if err != nil {
				logger.GetLogger().Error(err)
				return
			}
			logger.GetLogger().Info(string(buf[:n]))
			data := operateStep{}
			err = json.Unmarshal(buf[:n], &data)
			if err != nil {
				logger.GetLogger().Error(err)
			}

			control.operateChan <- data
		}
		//logger.GetLogger().Info(string(result))

	}()
	go control.sendDataToRemoteClient(wg)
}

func (control *operateControl) sendDataToRemoteClient(wg *sync.WaitGroup) {
	wg.Wait()
	for {
		select {
		case step := <-control.sendOperateChan:
			selfStep := step
			//logger.GetLogger().Info(step.ClientKey, step.Event.String())
			data, err := json.Marshal(&selfStep)
			//logger.GetLogger().Info(string(data))
			localStep := operateStep{}
			json.Unmarshal(data, &localStep)

			if err != nil {
				logger.GetLogger().Error(err)
				continue
			}
			//data = append(data, byte(0xff))
			logger.GetLogger().Info(control.conn.LocalAddr().String(), len(data))
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

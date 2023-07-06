package task

import "context"

type cmdTask struct {
	taskId      string
	taskName    string
	runFunction func(ctx context.Context) error
	cancelFunc  context.CancelFunc
	allRunTime  int64
	timeChan    chan bool
	status      string
	taskType    string
}

func (ctask *cmdTask) changeStatus(status string) {
	ctask.status = status
}

func (ctask *cmdTask) start() {
	ctask.changeStatus("running")
}

func (ctask *cmdTask) stop() {
	ctask.changeStatus("stop")
}

package task

import (
	"changeme/src/logger"
	"context"
	"github.com/google/uuid"
	"time"
)

type taskManager struct {
	taskMap          map[string]cmdTask
	taskChan         chan string
	maxTaskCount     int
	runningTaskCount int
}

var manager = taskManager{
	taskMap:          map[string]cmdTask{},
	taskChan:         make(chan string, 1000),
	maxTaskCount:     100,
	runningTaskCount: 0,
}

func GetTaskManager() *taskManager {
	return &manager
}

func init() {
	go GetTaskManager().runTaskFunc()
}

func (manager *taskManager) runTaskFunc() {
	for {
		select {
		case nextTaskId := <-manager.taskChan:
			logger.GetLogger().Info("任务ID：", nextTaskId)
			nextTask := manager.taskMap[nextTaskId]
			ctx, cancel := context.WithCancel(context.Background())
			nextTask.cancelFunc = cancel
			manager.taskMap[nextTaskId] = nextTask
			go nextTask.runFunction(ctx)
			select {
			case <-ctx.Done():

			}
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

func (manager *taskManager) reckonTime(taskId string) {
	for {

		time.Sleep(time.Second * 1)

		task := manager.taskMap[taskId]
		if <-task.timeChan {
			break
		}
		task.allRunTime = time.Now().Unix()
		manager.taskMap[taskId] = task
	}

}

func (manager *taskManager) CreateTask(taskFunc func(ctx context.Context) error, name string, taskType string) {

	task := cmdTask{
		taskId:      uuid.New().String(),
		taskName:    name,
		runFunction: taskFunc,
		allRunTime:  0,
		status:      "stop",
		taskType:    taskType,
	}
	manager.taskMap[task.taskId] = task
	logger.GetLogger().Info("创建任务:", manager.taskMap[task.taskId])
	manager.startTask(task.taskId)
}

func (manager *taskManager) startTask(taskId string) {
	task := manager.taskMap[taskId]

	if task.status == "running" {
		return
	}
	logger.GetLogger().Info("开始任务:", task.taskId)
	task.changeStatus("running")
	manager.taskChan <- task.taskId

}

func (manager *taskManager) stopTask(taskId string) {
	task := manager.taskMap[taskId]
	if task.status == "stop" {
		return
	}
	task.changeStatus("stop")
	task.cancelFunc()
	task.timeChan <- true
	manager.taskMap[taskId] = task
}

func (manager *taskManager) deleteTask(taskId string) {
	task := manager.taskMap[taskId]
	if task.status == "running" {
		manager.stopTask(taskId)
	}
	delete(manager.taskMap, taskId)
}

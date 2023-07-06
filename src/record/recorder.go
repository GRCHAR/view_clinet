package record

import (
	"changeme/src/ffmpeg"
	"changeme/src/task"
	"context"
)

type recorder struct {
}

var (
	fp      = ffmpeg.New()
	manager = task.GetTaskManager()
)

func NewRecorder() recorder {
	return *new(recorder)
}

func (r *recorder) RecordScreen() {

	manager.CreateTask(getRecordFunc(), "testname", "screen")
}

func getRecordFunc() func(context.Context) error {
	return func(ctx context.Context) error {
		return fp.RecordScreen(ctx)
	}
}

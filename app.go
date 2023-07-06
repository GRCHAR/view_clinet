package main

import (
	"changeme/src/logger"
	"changeme/src/record"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

var (
	recorder = record.NewRecorder()
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {

}

func (a *App) ScreenRecord() {
	logger.GetLogger().Info("start  screen record")
	recorder.RecordScreen()
	logger.GetLogger().Info("stop   screen record")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) JsLog(message string) {
	logger.GetLogger().Info("js log:", message)
}

package ffmpeg

import (
	"bytes"
	"changeme/src/logger"
	"context"
	"os/exec"
	"time"
)

type ffmpeg struct {
	screenNumber string
	fps          string
	videoCode    string
	preset       string
	outPath      string
	ffmpegPath   string
	osType       string
}

var log = logger.GetLogger()

func init() {

}

func New() *ffmpeg {
	return new(ffmpeg)
}

func (ff *ffmpeg) RecordScreen(ctx context.Context) error {
	ff.fps = "30"
	ff.screenNumber = "1"
	ff.videoCode = "libx264"
	ff.outPath = "./"
	cmd := exec.CommandContext(ctx, "ffmpeg", "-f", "avfoundation", "-capture_cursor", "1", "-pixel_format", "uyvy422", "-i", ff.screenNumber, "-r",
		ff.fps, "-c:v", ff.videoCode, "-preset", "ultrafast", ff.outPath+time.Now().Format("060102150405")+".mp4")
	outBuffer := bytes.Buffer{}
	errBuffer := bytes.Buffer{}
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	logger.GetLogger().Info("执行命令：", cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Info("recordScreen命令运行错误!\n", err)
		log.Info("输出:\n", errBuffer.String())
		return err
	}
	return nil
}

func (ff *ffmpeg) screenShoot() {

}

func (ff *ffmpeg) getVideoDevicesInfo(ctx context.Context) {
	cmd := &exec.Cmd{}
	switch ff.osType {
	case "mac":
		cmd = exec.CommandContext(ctx, "ffmpeg", "-f", "avfoundation", "-list_devices", "true", "-i", "''")
	case "windows":
		cmd = exec.CommandContext(ctx, "ffmpeg")
	}
	outBuffer := bytes.Buffer{}
	errBuffer := bytes.Buffer{}
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	logger.GetLogger().Info("执行命令：", cmd.String())
	err := cmd.Run()
	if err != nil {
		logger.GetLogger().Error(err)
	}
	logger.GetLogger().Info(outBuffer.String())
}

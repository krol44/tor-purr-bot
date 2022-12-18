package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strconv"
)

type Struct struct {
	IsDev bool

	ChatIdChannelLog int64

	DirBot        string
	DirDB         string
	DownloadLimit int

	WelcomeFileId string

	BotDebug      bool
	BotToken      string
	TgPathLocal   string
	TgApiEndpoint string

	AllowVideoFormats []string

	ActiveRange     int
	ActiveRangeUser int
}

var config Struct

func init() {
	dev := os.Getenv("dev") == "true"
	dl, _ := strconv.Atoi(os.Getenv("download_limit"))
	chatIdChannelLog, _ := strconv.ParseInt(os.Getenv("chat_id_channel_log"), 10, 64)
	botDebug := os.Getenv("bot_debug") == "true"

	config = Struct{
		dev,
		chatIdChannelLog,
		os.Getenv("dir_bot"),
		os.Getenv("dir_db"),
		dl,
		os.Getenv("welcome_video_id"),
		botDebug,
		os.Getenv("bot_token"),
		os.Getenv("tg_path_local"),
		"http://" + os.Getenv("tg_api_endpoint") + "/bot%s/%s",
		[]string{".avi", ".mkv", ".mp4", "MP4", "m4v", "flv", "TS", "ts", ".mov", "wmv", "webm"},
		3,
		1,
	}

	logSetup()
}

func logSetup() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	})
	if l, err := log.ParseLevel("debug"); err == nil {
		log.SetLevel(l)
		log.SetReportCaller(l == log.DebugLevel)
		log.SetOutput(os.Stdout)
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

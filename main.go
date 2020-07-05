package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

// replaced by ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type appContext struct {
	client *twitch.Client

	channel string
}

var Context appContext

func main() {
	if len(os.Args) != 2 {
		fmt.Println("tclog", version)
		fmt.Println("Usage: tclog [channel]")
		os.Exit(1)
	}

	Context.client = twitch.NewAnonymousClient()

	Context.client.OnConnect(func() {
		logMessage("CONNECT", []byte(""))
	})

	Context.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnWhisperMessage(func(message twitch.WhisperMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnClearChatMessage(func(message twitch.ClearChatMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnClearMessage(func(message twitch.ClearMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnRoomStateMessage(func(message twitch.RoomStateMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnUserStateMessage(func(message twitch.UserStateMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnGlobalUserStateMessage(func(message twitch.GlobalUserStateMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnNoticeMessage(func(message twitch.NoticeMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnUserJoinMessage(func(message twitch.UserJoinMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnUserPartMessage(func(message twitch.UserPartMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnReconnectMessage(func(message twitch.ReconnectMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.OnUnsetMessage(func(message twitch.RawMessage) {
		msg, _ := json.Marshal(message)
		logMessage(message.RawType, msg)
	})

	Context.client.Join(os.Args[1])

	err := Context.client.Connect()
	if err != nil {
		panic(err)
	}
}

func logMessage(messageType string, message []byte) {
	currentTime := time.Now()

	logMsg := `{"Time":"` + currentTime.Local().String() + `","MessageType":"` + messageType + `","Message":` + string(message[:]) + `}`
	fmt.Println(logMsg)

	writeLog("tclog_"+os.Args[1]+"_"+currentTime.Format("2006-01-02")+".log", logMsg)
}

func writeLog(fileName string, message string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintln(f, message)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

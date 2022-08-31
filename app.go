package main

import (
	"context"
	"fmt"
	// "time"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var MainApp *App

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	MainApp = &App{}
	return MainApp
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// go func() {
	// 	for {
	// 		msg := fmt.Sprintf("TS:%d",time.Now().Unix())
	// 		fmt.Println(msg)
	// 		runtime.EventsEmit(a.ctx, "MESSAGE_TEST", msg)
	// 		time.Sleep(time.Second)
	// 	}
	// }()
}

func (a *App) Greet(outputFilePath string) string {
	fmt.Println(outputFilePath)
	if len(outputFilePath) == 0 {
		return outputFilePath
	}
	err := NewDownload(outputFilePath)
	if err != nil {
		return err.Error()
	}
	return "完成"
}

func (a *App) Broadcast(channal string, msg string) {
	runtime.EventsEmit(a.ctx, channal, msg)
}
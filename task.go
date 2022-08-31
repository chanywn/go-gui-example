package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"io/ioutil"
	"math"
	"strings"
	"github.com/monkeyWie/gopeed-core/pkg/base"
	"github.com/monkeyWie/gopeed-core/pkg/download"
)

var (
	outputFilePath string
 	unitArr = []string{"B", "KB", "MB", "GB", "TB", "EB"}
	TaskCount int
	FinishTaskCount int
	RootDir string
)

const progressWidth = 20

type Task struct {
	Url string
	Dir string
	Dirs []string
	Name string
	Error string
}

func NewTak(url string) *Task {
	t := &Task {
		Url:url,
	}
	path := strings.SplitN(url, "/", -1)
	t.Dirs = path[3:len(path)-1]
	t.Dir = strings.Join(t.Dirs,"/")
	t.Name = path[len(path)-1]
	return t
}

func (t *Task)Downloader() {
	fmt.Println("start",t.Name)
	finallyCh := make(chan error, 10)
	err := download.Boot().
		URL(t.Url).
		Listener(func(event *download.Event) {
			if FinishTaskCount == TaskCount {
				finallyCh <- event.Err
				return
			}
			switch event.Key {
			case download.EventKeyProgress:
				fmt.Println("err")
				printProgress(event.Task, FinishTaskCount+1, TaskCount)
			case download.EventKeyFinally:
				printProgress(event.Task, FinishTaskCount+1, TaskCount)
				FinishTaskCount++
				finallyCh <- event.Err
			}
		}).
		Create(&base.Options{
			Name:t.Name,
			Path: fmt.Sprintf("%s/download",RootDir),
			Connections: 10,
		})
	if err != nil {
		t.Error = err.Error()
		fmt.Println(err)
		return
	}
	<-finallyCh

	fmt.Println("end",t.Name)
}

func printProgress(task *download.TaskInfo, f int, a int) {
	rate := float64(task.Progress.Downloaded) / float64(task.Res.TotalSize)
	// completeWidth := int(progressWidth * rate)
	speed := ByteFmt(task.Progress.Speed)
	totalSize := ByteFmt(task.Res.TotalSize)
	// fmt.Printf("] %.1f%%    %s/s    %s", rate*100, speed, totalSize)
	msg := fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t",f,a,int(rate*100),speed,totalSize)
	MainApp.Broadcast("MESSAGE_TEST", msg)
}

func ByteFmt(size int64) string {
	if size == 0 {
		return "0"
	}
	fs := float64(size)
	p := int(math.Log(fs) / math.Log(1024))
	val := fs / math.Pow(1024, float64(p))
	_, frac := math.Modf(val)
	if frac > 0 {
		return fmt.Sprintf("%.1f%s", math.Floor(val*10)/10, unitArr[p])
	} else {
		return fmt.Sprintf("%d%s", int(val), unitArr[p])
	}
}

func GetCurrPath() string {
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))
    ret := path[:index]
    return ret
}

func NewDownload(outputFilePath string) error {
	content, err := ioutil.ReadFile(outputFilePath)
    if err != nil {
        return err
    }
	files := strings.SplitN(string(content), "\n", -1)
	tasks := []*Task{}
	for _, v := range files {
		if len(v) == 0 {
			continue
		}
		v = strings.TrimRight(v, "\r")
		tasks = append(tasks, NewTak(v))
	}
	// RootDir = "/Users/aomeng/projects/myproject"
	if _, err := os.Stat(fmt.Sprintf("%s/download",RootDir)); err != nil {
		if err := os.MkdirAll(fmt.Sprintf("%s/download",RootDir), 0711);err != nil {
			return err
		}
	}
	max := len(tasks)
	for i, t := range tasks {
		msg := fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t",i+1,max,0,"0 MB/s","")
		MainApp.Broadcast("MESSAGE_TEST", msg)
		t.Downloader()
		msg = fmt.Sprintf("%d\t%d\t%d\t%s\t%s\t",i+1,max,100,"0 MB/s","")
		MainApp.Broadcast("MESSAGE_TEST", msg)
		fmt.Println(t.Name ,"______")
	}
	return nil
}



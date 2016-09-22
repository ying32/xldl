// XLSdkTest project main.go
// 迅雷下载sdk 翻译by: ying32  qq:1444386932
// api说明见：http://open.xunlei.com/wiki/api_doc.html#1
package main

import (
	"fmt"
	"os"

	"time"

	"github.com/ying32/xldl"

	"os/exec"
	"path/filepath"
)

var (
	exitLoop bool
)

func main() {

	file, _ := exec.LookPath(os.Args[0])
	path := filepath.Dir(file)

	println(path)

	if !xldl.InitXLEngine() {
		fmt.Println("初始引擎失败！")
		return
	}
	defer xldl.UnInitXLEngine()
	dloader := xldl.NewXLDownloader(path)

	task := dloader.AddTask("http://sw.bos.baidu.com/sw-search-sp/software/19de58890ffb8/QQ_8.6.18804.0_setup.exe", "QQ_8.6.18804.0_setup.exe")
	fmt.Println("任务启动：", task.Start())

	go func() {
		var exit string
		for {
			fmt.Scan(&exit)
			if exit == "q" {
				exitLoop = true
				break
			}
		}
	}()

	for !exitLoop {
		info, ret := task.Info()
		if ret {
			fmt.Printf("TotalSize=%d, Percent=%f, Speed=%d\n", info.TotalSize, info.Percent, info.Speed)
			switch info.Stat {
			case xldl.TSC_COMPLETE:
				fmt.Println("下载完成")
				task.Delete()
				return
			case xldl.TSC_ERROR:
				fmt.Println("下载错误。")
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
	task.Stop()
	task.Delete()
	fmt.Println("删除：", task.DeleteTempFile())
	dloader.RemoveAll()
}

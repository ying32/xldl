package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ying32/xldl"
)

var (
	dloader xldl.XLDownloader
)

func homeHandler(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("hello ying!"))
}

func main() {
	/*	http.Handle("/", http.FileServer(http.Dir(SvrConf.Path)))
		err = http.ListenAndServe(fmt.Sprintf(":%d", SvrConf.Port), nil)
	*/
	if !xldl.InitXLEngine() {
		panic("初始下载引擎失败。")
	}
	defer xldl.UnInitXLEngine()
	file, _ := exec.LookPath(os.Args[0])
	path := filepath.Dir(file)

	dloader := xldl.NewXLDownloader(path)
	defer dloader.RemoveAll()

	http.HandleFunc("/", homeHandler)
	fmt.Println("服务启动:")
	err := http.ListenAndServe(":9991", nil)
	if err != nil {
		fmt.Println("服务启动失败。")
	}

}

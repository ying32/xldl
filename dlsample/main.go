package main

import (
	"fmt"

	"net/http"
)

func homeHandler(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("hello ying!"))
}

func main() {
	/*	http.Handle("/", http.FileServer(http.Dir(SvrConf.Path)))
		err = http.ListenAndServe(fmt.Sprintf(":%d", SvrConf.Port), nil)
	*/
	http.HandleFunc("/", homeHandler)
	fmt.Println("服务启动:")
	err := http.ListenAndServe(":9991", nil)
	if err != nil {
		fmt.Println("服务启动失败。")
	}

}

// xldl project xldl.go
// 迅雷下载sdk 翻译by: ying32  qq:1444386932
// api说明见：http://open.xunlei.com/wiki/api_doc.html#1
package xldl

import (
	"syscall"
)

type XLTask struct {
	hander   uintptr
	Url      string
	FileName string
	SavePath string
	param    *DownTaskParam
}

type XLDownloader struct {
	Tasks    map[string]*XLTask
	SavePath string
}

func NewXLDownloader(savePath string) *XLDownloader {
	dloader := &XLDownloader{}
	dloader.SavePath = savePath
	dloader.Tasks = make(map[string]*XLTask)
	return dloader
}

func UnInitXLEngine() bool {
	return XL_UnInit()
}

func InitXLEngine() bool {
	return XL_Init()
}

/* XLTask */

// 启动一个任务
func (self *XLTask) Start() bool {
	if self.hander == 0 {
		return false
	}
	return XL_StartTask(self.hander)
}

// 停止一个任务
func (self *XLTask) Stop() bool {
	if self.hander == 0 {
		return false
	}
	return XL_StopTask(self.hander)
}

// 删除一个任务
func (self *XLTask) Delete() bool {
	if self.hander == 0 {
		return false
	}
	ret := XL_DeleteTask(self.hander)
	self.hander = 0
	return ret
}

func (self *XLTask) Info() (*DownTaskInfo, bool) {
	if self.hander == 0 {
		return nil, false
	}
	return XL_QueryTaskInfoEx(self.hander)
}

func (self *XLTask) DeleteTempFile() bool {
	if self.param == nil {
		return false
	}
	return XL_DelTempFile(self.param)
}

/* XLDownloader */

func (self *XLDownloader) AddTask(wstrUrl, wstrFileName string) *XLTask {
	if v, ok := self.Tasks[wstrUrl]; ok {
		return v
	}
	param := new(DownTaskParam)
	param.SetDefault()

	copy(param.szTaskUrl[:], syscall.StringToUTF16(wstrUrl))
	copy(param.szFilename[:], syscall.StringToUTF16(wstrFileName))
	copy(param.szSavePath[:], syscall.StringToUTF16(self.SavePath))

	xltask := &XLTask{0, wstrUrl, wstrFileName, self.SavePath, param}
	xltask.hander = XL_CreateTask(param)
	self.Tasks[wstrUrl] = xltask
	return xltask
}

func (self *XLDownloader) Remove(task *XLTask) {
	if _, ok := self.Tasks[task.Url]; ok {
		self.Tasks[task.Url] = nil
		delete(self.Tasks, task.Url)
	}
	task.Delete()
}

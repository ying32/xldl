// xldl project xldl.go
// 迅雷下载sdk 翻译by: ying32  qq:1444386932
// api说明见：http://open.xunlei.com/wiki/api_doc.html#1
package xldl

type XLTask struct {
	hander uintptr
	url    string
}

type XLDownloader struct {
	Tasks map[string]*XLTask
}

func NewXLDownloader() *XLDownloader {
	dloader := &XLDownloader{}
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
	self.hander = 0;
	return ret
}

func (self *XLTask) Info() (*DownTaskInfo, bool) {
	if self.hander == 0 {
		return nil, false
	}
	return XL_QueryTaskInfoEx(self.hander)
}

/* XLDownloader */

func (self *XLDownloader) AddTask(wstrUrl, wstrFileName, wstrSavePath string) *XLTask {
	if v, ok := self.Tasks[wstrUrl]; ok {
		return v
	}
	param := new(DownTaskParam)
	param.SetDefault()
	xltask := &XLTask{XL_CreateTask(param.GetParam(wstrUrl, wstrFileName, wstrSavePath)), wstrUrl}
	self.Tasks[wstrUrl] = xltask
	return xltask
}

func (self *XLDownloader) Remove(task *XLTask) {
	if _, ok := self.Tasks[task.url]; ok {
		self.Tasks[task.url] = nil
	}
	task.Delete()
}

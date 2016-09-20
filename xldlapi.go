// 迅雷下载sdk 翻译by: ying32  qq:1444386932
// api说明见：http://open.xunlei.com/wiki/api_doc.html#1
package xldl

import (
	"syscall"
	"unsafe"
)

var (
	xldldll          = syscall.NewLazyDLL("xldl.dll")
	xl_Init          = xldldll.NewProc("XL_Init")
	xl_UnInit        = xldldll.NewProc("XL_UnInit")
	xl_CreateTask    = xldldll.NewProc("XL_CreateTask")
	xl_DeleteTask    = xldldll.NewProc("XL_DeleteTask")
	xl_StartTask     = xldldll.NewProc("XL_StartTask")
	xl_StopTask      = xldldll.NewProc("XL_StopTask")
	xl_ForceStopTask = xldldll.NewProc("XL_ForceStopTask")

	//	xl_QueryTaskInfo          = xldldll.NewProc("XL_QueryTaskInfo")

	xl_QueryTaskInfoEx = xldldll.NewProc("XL_QueryTaskInfoEx")

	//	xl_DelTempFile            = xldldll.NewProc("XL_DelTempFile")

	xl_SetSpeedLimit       = xldldll.NewProc("XL_SetSpeedLimit")
	xl_SetUploadSpeedLimit = xldldll.NewProc("XL_SetUploadSpeedLimit")

	//	xl_SetProxy               = xldldll.NewProc("XL_SetProxy")

	xl_SetUserAgent = xldldll.NewProc("XL_SetUserAgent")

//	xl_ParseThunderPrivateUrl = xldldll.NewProc("XL_ParseThunderPrivateUrl")
//	xl_GetFileSizeWithUrl     = xldldll.NewProc("XL_GetFileSizeWithUrl")
//	xl_SetFileIdAndSize       = xldldll.NewProc("XL_SetFileIdAndSize")
//	xl_SetAdditionInfo        = xldldll.NewProc("XL_SetAdditionInfo")

//	xl_CreateTaskByURL        = xldldll.NewProc("XL_CreateTaskByURL")
//	xl_CreateTaskByThunder    = xldldll.NewProc("XL_CreateTaskByThunder")
//	xl_CreateBTTaskByThunder  = xldldll.NewProc("XL_CreateBTTaskByThunder")
)

const (
	MAX_PATH = 260
)

type DOWN_TASK_STATUS int32

const (
	NOITEM = 0 + iota
	TSC_ERROR
	TSC_PAUSE
	TSC_DOWNLOAD
	TSC_COMPLETE
	TSC_STARTPENDING
	TSC_STOPPENDING
)

type TASK_ERROR_TYPE int32

const (
	TASK_ERROR_UNKNOWN                = 0x00 // 未知错误
	TASK_ERROR_DISK_CREATE            = 0x01 // 创建文件失败
	TASK_ERROR_DISK_WRITE             = 0x02 // 写文件失败
	TASK_ERROR_DISK_READ              = 0x03 // 读文件失败
	TASK_ERROR_DISK_RENAME            = 0x04 // 重命名失败
	TASK_ERROR_DISK_PIECEHASH         = 0x05 // 文件片校验失败
	TASK_ERROR_DISK_FILEHASH          = 0x06 // 文件全文校验失败
	TASK_ERROR_DISK_DELETE            = 0x07 // 删除文件失败失败
	TASK_ERROR_DOWN_INVALID           = 0x10 // 无效的DOWN地址
	TASK_ERROR_PROXY_AUTH_TYPE_UNKOWN = 0x20 // 代理类型未知
	TASK_ERROR_PROXY_AUTH_TYPE_FAILED = 0x21 // 代理认证失败
	TASK_ERROR_HTTPMGR_NOT_IP         = 0x30 // http下载中无ip可用
	TASK_ERROR_TIMEOUT                = 0x40 // 任务超时
	TASK_ERROR_CANCEL                 = 0x41 // 任务取消
	TASK_ERROR_TP_CRASHED             = 0x42 // MINITP崩溃
	TASK_ERROR_ID_INVALID             = 0x43 // TaskId 非法
)

// 他按4字节对齐的，然迅雷要求按1字节对齐，so下面两个有问题了，要换种方式了
type DownTaskInfo struct {
	Stat           DOWN_TASK_STATUS
	FailCode       TASK_ERROR_TYPE
	Filename       [MAX_PATH]uint16
	Reserved0      [MAX_PATH]uint16
	TotalSize      int64   // 该任务总大小(字节)
	TotalDownload  int64   // 下载有效字节数(可能存在回退的情况)
	Percent        float32 // 下载进度
	reserved0      int32
	SrcTotal       int32 // 总资源数
	SrcUsing       int32 // 可用资源数
	reserved1      int32
	reserved2      int32
	reserved3      int32
	reserved4      int32
	reserved5      int64
	DonationP2P    int64 // p2p贡献字节数
	reserved6      int64
	DonationOrgin  int64 // 原始资源共享字节数
	DonationP2S    int64 // 镜像资源共享字节数
	reserved7      int64
	reserved8      int64
	Speed          int32   // 即时速度(字节/秒)
	SpeedP2S       int32   // 即时速度(字节/秒)
	SpeedP2P       int32   // 即时速度(字节/秒)
	IsOriginUsable bool    // 原始资源是否有效
	HashPercent    float32 // 现不提供该值
	IsCreatingFile bool    // 是否正在创建文件
	reserved       [64]uint32
}

func (c *DownTaskInfo) SetDefault() {
	c.Stat = TSC_PAUSE
	c.FailCode = TASK_ERROR_UNKNOWN
	c.Percent = 0.0
	c.IsOriginUsable = false
	c.HashPercent = 0
}

type DownTaskParam struct {
	nReserved         int32
	szTaskUrl         [2084]uint16     // 任务URL
	szRefUrl          [2084]uint16     // 引用页
	szCookies         [4096]uint16     // 浏览器cookie
	szFilename        [MAX_PATH]uint16 // 下载保存文件名.
	szReserved0       [MAX_PATH]uint16
	szSavePath        [MAX_PATH]uint16 // 文件保存目录
	hReserved         int32
	bReserved         int32
	szReserved1       [64]uint16
	szReserved2       [64]uint16
	IsOnlyOriginal    int32 // 是否只从原始地址下载
	nReserved1        uint32
	DisableAutoRename int32 // 禁止智能命名
	IsResume          int32 // 是否用续传
	reserved          [2048]uint32
}

func (c *DownTaskParam) SetDefault() {
	c.nReserved1 = 5
	c.bReserved = 0
	c.DisableAutoRename = 0
	c.IsOnlyOriginal = 0
	c.IsResume = 1
}

func (c *DownTaskParam) GetParam(wstrUrl, wstrFileName, wstrSavePath string) *DownTaskParam {
	ubytes := syscall.StringToUTF16(wstrUrl)
	for i := 0; i < len(ubytes); i++ {
		c.szTaskUrl[i] = ubytes[i]
	}
	ubytes = syscall.StringToUTF16(wstrFileName)
	for i := 0; i < len(ubytes); i++ {
		c.szFilename[i] = ubytes[i]
	}
	ubytes = syscall.StringToUTF16(wstrSavePath)
	for i := 0; i < len(ubytes); i++ {
		c.szSavePath[i] = ubytes[i]
	}
	return c
}

func XL_Init() bool {
	ret, _, _ := xl_Init.Call()
	return ret != 0
}

func XL_UnInit() bool {
	ret, _, _ := xl_UnInit.Call()
	return ret != 0
}

func XL_CreateTask(param *DownTaskParam) uintptr {
	ret, _, _ := xl_CreateTask.Call(uintptr(unsafe.Pointer(param)))
	return ret
}

func XL_DeleteTask(hTask uintptr) bool {
	ret, _, _ := xl_DeleteTask.Call(hTask)
	return ret != 0
}

func XL_StartTask(hTask uintptr) bool {
	ret, _, _ := xl_StartTask.Call(hTask)
	return ret != 0
}

func XL_StopTask(hTask uintptr) bool {
	ret, _, _ := xl_StopTask.Call(hTask)
	return ret != 0
}

func XL_ForceStopTask(hTask uintptr) bool {
	ret, _, _ := xl_ForceStopTask.Call(hTask)
	return ret != 0
}

func XL_SetSpeedLimit(nKBps int32) {
	xl_SetSpeedLimit.Call(uintptr(nKBps))
}

func XL_SetUploadSpeedLimit(nTcpKBps, nOtherKBps uint32) {
	xl_SetUploadSpeedLimit.Call(uintptr(nTcpKBps), uintptr(nOtherKBps))
}

func XL_SetUserAgent(pszUserAgent string) {
	xl_SetUserAgent.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(pszUserAgent))))
}

func XL_QueryTaskInfoEx(hTask uintptr) (*DownTaskInfo, bool) {
	info := &DownTaskInfo{}
	info.SetDefault()
	ret, _, _ := xl_QueryTaskInfoEx.Call(hTask, uintptr(unsafe.Pointer(info)))
	return info, ret != 0
}

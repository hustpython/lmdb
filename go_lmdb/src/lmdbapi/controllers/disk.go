package controllers

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"lmdbapi/models"
	"strings"
)

type DiskController struct {
	ErrorController
}

//PS C:\Users\mxq> Get-WmiObject Win32_logicaldisk

//DeviceID     : C:
//DriveType    : 3
//ProviderName :
//FreeSpace    : 55401881600
//Size         : 161061269504
//VolumeName   : OS
//
//DeviceID     : D:
//DriveType    : 3
//ProviderName :
//FreeSpace    : 4375560192
//Size         : 79820746752
//VolumeName   : Data
//
//DeviceID     : E:
//DriveType    : 3 磁盘 2 U盘
//ProviderName :
//FreeSpace    : 494309482496
//Size         : 1000202039296
//VolumeName   : 数据
// https://learn.microsoft.com/zh-cn/windows/win32/cimwin32prov/win32-logicaldisk

type logicalDisk struct {
	DeviceID  string
	DriveType uint32
	FreeSpace uint64
	Size      uint64
}

type logicalDiskRsp struct {
	logicalDisk
	DriveTypeStr string
	FreeSpaceStr string
	SizeStr      string
	VideoSize    uint64
	VideoSizeStr string
	VideoCount   uint32
}

var driveTypeMap = map[uint32]string{
	3: "硬盘",
	2: "U盘",
}

func (d *DiskController) GetDiskInfo() {
	var infos []logicalDisk
	err := wmi.Query("Select DeviceID, DriveType, Size,FreeSpace from Win32_LogicalDisk", &infos)
	if err != nil {
		d.setInternalError(err.Error())
		return
	}
	mRes := models.GetMovieSizeAndCount()
	var resRsp []*logicalDiskRsp
	for i := range infos {
		var res logicalDiskRsp
		res.logicalDisk = infos[i]
		res.DeviceID = strings.Trim(infos[i].DeviceID, ":")
		res.FreeSpaceStr = byteSizeTransform(infos[i].FreeSpace)
		res.SizeStr = byteSizeTransform(infos[i].Size)
		res.DriveTypeStr = res.DeviceID + "-" + driveTypeMap[infos[i].DriveType]
		if mRes[res.DeviceID] != nil {
			res.VideoSize = mRes[res.DeviceID].Size
			res.VideoCount = mRes[res.DeviceID].Count
			res.VideoSizeStr = byteSizeTransform(mRes[res.DeviceID].Size)
		}
		resRsp = append(resRsp, &res)
	}
	d.Data["json"] = resRsp
	d.ServeJSON()
}

// 以1024作为基数
func byteSizeTransform(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %c",
		float64(b)/float64(div), "KMGTPE"[exp])
}

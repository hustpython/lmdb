package controllers

import (
	"fmt"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	//fmt.Println(DiskUsage().All/1024/1024/1024)
	s := getLogicalDisk()
	for _, k := range s {
		fmt.Println(k.DeviceID, ByteCountIEC(k.Size), ByteCountIEC(k.FreeSpace))
	}
}

package health

import (
	"github.com/yoyofx/yoyogo/abstractions/platform/systeminfo"
	"os"
)

type DiskHealthIndicator struct{}

func NewDiskHealthIndicator() *DiskHealthIndicator {
	return &DiskHealthIndicator{}
}

func (u *DiskHealthIndicator) Health() ComponentStatus {
	wd, _ := os.Getwd()
	diskStatus := systeminfo.DiskUsage(wd)
	status := Up("diskHealth")
	if diskStatus.Free < 5242880 {
		status.SetStatus("down")
	}

	return status.WithDetail("disk", diskStatus)
}

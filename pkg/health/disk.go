package health

import (
	"github.com/yoyofx/yoyogo/abstractions/health"
	"github.com/yoyofx/yoyogo/abstractions/platform/systeminfo"
	"os"
)

type DiskHealthIndicator struct{}

func (u DiskHealthIndicator) Health() health.ComponentStatus {
	wd, _ := os.Getwd()
	diskStatus := systeminfo.DiskUsage(wd)
	status := health.Up("diskHealth")
	if diskStatus.Free < 5242880 {
		status.SetStatus("down")
	}

	return status.WithDetail("disk", diskStatus)
}

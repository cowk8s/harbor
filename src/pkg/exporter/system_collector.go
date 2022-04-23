package exporter

const (
	systemInfoCollectorName = "SystemInfoCollector"
	sysInfoURL = "/api/v2.0/systeminfo"
)

var (
	harborSysInfo = 
)

func NewSystemInfoCollector(hbrCli *HarborClient)
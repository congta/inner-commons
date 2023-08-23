package env

import "github.com/duke-git/lancet/v2/netutil"

const (
	UnknownRegion = "-"
	R_CN          = "CN"
	R_SG          = "SG"
	R_US          = "US"

	UnknownIDC = "-"
	DC_ZJK     = "zjk" // ALI ZJK
	DC_SEA     = "sea" // Vultr Seattle
	DC_SG      = "sg"  // Vultr Singapore
)

var (
	psm     = ""
	region  = UnknownRegion
	idc     = UnknownIDC
	cluster = "default"
)

func SetRegion(s string) {
	region = s
}

func Region() string {
	return region
}

func SetCluster(s string) {
	cluster = s
}

func Cluster() string {
	return cluster
}

func SetPSM(s string) {
	psm = s
}

func PSM() string {
	return psm
}

func SetIDC(s string) {
	idc = s
}

func IDC() string {
	return idc
}

func Env() string {
	return "prod"
}

func Stage() string {
	return "single_dc"
}

func IsProduct() bool {
	return true
}

func IsPreview() bool {
	return false
}

func IsTest() bool {
	return false
}

// InCE in container engine
func InCE() bool {
	return true
}

func HostIP() string {
	return HostIPV4()
}

func HostIPV4() string {
	ips := netutil.GetIps()
	if len(ips) == 0 {
		return ""
	}
	return ips[0]
}

func GetIDCFromHost(ip string) string {
	// lookup by ip
	return UnknownIDC
}

func GetRegionFromIDC(idc string) string {
	// lookup by idc
	return UnknownRegion
}

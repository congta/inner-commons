package env

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
	return UnknownRegion
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

func IsProduct() bool {
	return true
}

func IsTest() bool {
	return false
}

// InCE in container engine
func InCE() bool {
	return true
}

func HostIP() string {
	return ""
}

func GetIDCFromHost(ip string) string {
	// lookup by ip
	return UnknownIDC
}

func GetRegionFromIDC(idc string) string {
	// lookup by idc
	return UnknownRegion
}

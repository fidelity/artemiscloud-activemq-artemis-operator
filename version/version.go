package version

var (
	Version = "1.0.0"
	// PriorVersion - prior version
	PriorVersion = "0.20.1"
)

const (
	// LatestVersion product version supported
	LatestVersion        = "2.20.0"
	CompactLatestVersion = "2200"
	// LastMicroVersion product version supported
	LastMicroVersion = "2.18.0"
	// LastMinorVersion product version supported
	LastMinorVersion = "2.15.0"
)

// SupportedVersions - product versions this operator supports
var SupportedVersions = []string{LatestVersion, LastMicroVersion, LastMinorVersion}
var SupportedMicroVersions = []string{LatestVersion, LastMicroVersion}
var OperandVersionFromOperatorVersion map[string]string = map[string]string{
	"0.18.0": "2.15.0",
	"0.19.0": "2.16.0",
	"0.20.0": "2.18.0",
	"0.20.1": "2.18.0",
	"1.0.0":  "2.20.0",
}
var FullVersionFromMinorVersion map[string]string = map[string]string{
	"150": "2.15.0",
	"160": "2.16.0",
	"180": "2.18.0",
	"200": "2.20.0",
}

var CompactFullVersionFromMinorVersion map[string]string = map[string]string{
	"150": "2150",
	"160": "2160",
	"180": "2180",
	"200": "2200",
}

var CompactVersionFromVersion map[string]string = map[string]string{
	"2.15.0": "2150",
	"2.16.0": "2160",
	"2.18.0": "2180",
	"2.20.0": "2200",
}

var FullVersionFromCompactVersion map[string]string = map[string]string{
	"2150": "2.15.0",
	"2160": "2.16.0",
	"2180": "2.18.0",
	"2200": "2.20.0",
}

var MinorVersionFromFullVersion map[string]string = map[string]string{
	"2.15.0": "150",
	"2.16.0": "160",
	"2.18.0": "180",
	"2.20.0": "200",
}

//The yacfg profile to use for a given full version of broker
var YacfgProfileVersionFromFullVersion map[string]string = map[string]string{
	"2.15.0": "2.15.0",
	"2.16.0": "2.16.0",
	"2.18.0": "2.18.0",
	"2.20.0": "2.18.0",
}

var YacfgProfileName string = "artemis"

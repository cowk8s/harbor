package chart

import "time"

// VersionDetails keeps the detailed data info of the chart version
type VersionDetails struct {
	Values map[string]interface{} `json:"values"`
	Files  map[string]string      `json:"files"`
}

// SecurityReport keeps the info related with security
// e.g.: digital signature, vulnerability scanning etc.
type SecurityReport struct {
}

// DigitalSignature used to indicate if the chart has been signed
type DigitalSignature struct {
	Signed     bool   `json:"signed"`
	Provenance string `json:"prov_file"`
}

// Info keeps the information of the chart
type Info struct {
	Name          string    `json:"name"`
	TotalVersions uint32    `json:"total_versions"`
	LatestVersion string    `json:"latest_version"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	Icon          string    `json:"icon"`
	Home          string    `json:"home"`
	Deprecated    bool      `json:"deprecated"`
}

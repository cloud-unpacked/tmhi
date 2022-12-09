package cmd

type loginResp struct {
	Auth struct {
		Expiration int64  `json:"expiration"`
		Token      string `json:"token"`
	} `json:"auth"`
}

type device struct {
	Firmware     string `json:"softwareVersion"`
	FriendlyName string `json:"friendlyName"`
	MAC          string `json:"macID"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Serial       string `json:"serial"`
}

type deviceResp struct {
	Device device `json:"device"`
}

type signalMetrics struct {
	Bands []string `json:"bands"`
	Bars  float64  `json:"bars"`
	ENBID int      `json:"eNBID"`
	RSRQ  int      `json:"rsrq"`
	SINR  int      `json:"sinr"`
}

type signal struct {
	Signal4G signalMetrics `json:"4g"`
	Signal5G signalMetrics `json:"5g"`
	Generic  struct {
		Roaming bool `json:"roaming"`
	}
}

type signalResp struct {
	Signal signal `json:"signal"`
}

type timeResp struct {
	Uptime int `json:"upTime"`
}

type gatewayResp struct {
	Device device   `json:"device"`
	Signal signal   `json:"signal"`
	Time   timeResp `json:"time"`
}

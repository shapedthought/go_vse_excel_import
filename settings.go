package main

type ServerSpec struct {
	Cores    int32 `json:"cores"`
	RAM      int32 `json:"ram"`
	Capacity int32 `json:"capacity"`
}

type VbrSettings struct {
	NumVMwithPerVM    int `json:"numVMwithPerVM"`
	NumVMWithperJob   int `json:"numVmWithperJob"`
	VbrConcurrentJobs int `json:"vbrConcurrentJobs"`
	ConJobsForCores   int `json:"conJobsForCores"`
	ConJobsForMem     int `json:"conJobsForMem"`
	CoresFor25ConJobs int `json:"coresFor25ConJobs"`
	MemFor25ConJobs   int `json:"memFor25ConJobs"`
	MemPerConJobs     int `json:"memPerConJobs"`
}

type ProxySettings struct {
	IngestPerCPUCoreFull int `json:"ingestPerCpuCoreFull"`
	IngestPerCPUCoreInc  int `json:"ingestPerCpuCoreInc"`
	ProxyTaskConsumesMem int `json:"proxyTaskConsumesMem"`
}

type RepoSettings struct {
	DailyCrm          int    `json:"dailyCrm"`
	WeeklyCrm         int    `json:"weeklyCrm"`
	MonthlyCrm        int    `json:"monthlyCrm"`
	YearlyCrm         int    `json:"yearlyCrm"`
	RepoTaskConMemory int    `json:"repoTaskConMemory"`
	TaskCoreRatio     int    `json:"taskCoreRatio"`
	UseRPC            string `json:"useRpc"`
}

type EmSettings struct {
	EmUseAPIMemAdd        int `json:"emUseApiMemAdd"`
	EmUseAPICoreAdd       int `json:"emUseApiCoreAdd"`
	EmUseMultiVbrMemAdd   int `json:"emUseMultiVbrMemAdd"`
	EmUseMultiVbrCoresAdd int `json:"emUseMultiVbrCoresAdd"`
	EmUseSelfMemAdd       int `json:"emUseSelfMemAdd"`
	EmUseSelfCoresAdd     int `json:"emUseSelfCoresAdd"`
}

type Settings struct {
	ServerMin struct {
		VbrServer    ServerSpec `json:"vbrServer"`
		SqlServer    ServerSpec `json:"sqlServer"`
		VProxyServer ServerSpec `json:"vProxyServer"`
		RepoServer   ServerSpec `json:"repoServer"`
		EmServer     ServerSpec `json:"emServer"`
	} `json:"serverMin"`
	VbrSettings   VbrSettings   `json:"vbrSettings"`
	ProxySettings ProxySettings `json:"proxySettings"`
	RepoSettings  RepoSettings  `json:"repoSettings"`
	EmSettings    EmSettings    `json:"emSettings"`
}

type Workload struct {
	WorkloadActive  string  `json:"workloadActive"`
	WorkLoadName    string  `json:"workLoadName"`
	BackupType      string  `json:"backupType"`
	Site            string  `json:"site"`
	WorkLoadCap     float64 `json:"workLoadCap"`
	GrowthPercent   int     `json:"growthPercent"`
	ScopeYears      int     `json:"scopeYears"`
	BackupWindow    int     `json:"backupWindow"`
	ChangeRate      int     `json:"changeRate"`
	Reduction       int     `json:"reduction"`
	VMQty           int     `json:"vmQty"`
	VMVmdkRatio     int     `json:"vmVmdkRatio"`
	UsePerVM        string  `json:"usePerVM"`
	UseReFs         string  `json:"useReFs"`
	RpsBu           int     `json:"rpsBu"`
	BuWeekly        int     `json:"buWeekly"`
	BuMonthly       int     `json:"buMonthly"`
	BuYearly        int     `json:"buYearly"`
	CopySite        string  `json:"copySite"`
	RpsBuCopy       int     `json:"rpsBuCopy"`
	BuCopyWeekly    int     `json:"buCopyWeekly"`
	BuCopyMonthly   int     `json:"buCopyMonthly"`
	BuCopyYearly    int     `json:"buCopyYearly"`
	CloudMove       int     `json:"cloudMove"`
	CloudEnabled    bool    `json:"cloudEnabled"`
	ProcessCapacity float64 `json:"processCapacity"`
	BandWidthInc    float64 `json:"bandWidthInc"`
	VmdkQty         int     `json:"vmdkQty"`
}

type Site struct {
	SiteName      string `json:"siteName"`
	WanSpeed      int    `json:"wanSpeed"`
	NetworkSpeed  int    `json:"networkSpeed"`
	InternetSpeed int    `json:"internetSpeed"`
}

type VseInput struct {
	Workload []Workload `json:"workload"`
	Settings Settings   `json:"settings"`
	Sites    []Site     `json:"sites"`
}

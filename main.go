package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {

	settingsJson := `{
		"serverMin": {
			"vbrServer": {
				"cores": 4,
				"ram": 8
			},
			"sqlServer": {
				"cores": 2,
				"ram": 4
			},
			"vProxyServer": {
				"cores": 4,
				"ram": 8
			},
			"repoServer": {
				"cores": 20,
				"ram": 128,
				"capacity": 400
			},
			"emServer": {
				"cores": 2,
				"ram": 4
			}
		},
		"vbrSettings": {
			"numVMwithPerVM": 70,
			"numVmWithperJob": 30,
			"vbrConcurrentJobs": 10,
			"conJobsForCores": 25,
			"conJobsForMem": 25,
			"coresFor25ConJobs": 2,
			"memFor25ConJobs": 4,
			"memPerConJobs": 512
		},
		"proxySettings": {
			"ingestPerCpuCoreFull": 100,
			"ingestPerCpuCoreInc": 25,
			"proxyTaskConsumesMem": 2
		},
		"repoSettings": {
			"dailyCrm": 1,
			"weeklyCrm": 3,
			"monthlyCrm": 5,
			"yearlyCrm": 15,
			"repoTaskConMemory": 4,
			"TaskCoreRatio": 3,
			"UseRPC": "yes"
		},
		"emSettings": {
			"emUseApiMemAdd": 2,
			"emUseApiCoreAdd": 1,
			"emUseMultiVbrMemAdd": 2,
			"emUseMultiVbrCoresAdd": 1,
			"emUseSelfMemAdd": 4,
			"emUseSelfCoresAdd": 2
		}
	}`

	var settings Settings

	json.Unmarshal([]byte(settingsJson), &settings)

	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, _ := f.GetRows("Workload")

	// range over the rows
	// If the row
	var wls []Workload
	var s []string // sites

	for r, row := range rows {
		if r != 0 {
			var wl Workload
			for i, cell := range row {
				switch i {
				case 0:
					wl.Site = cell
					s = append(s, cell)
				case 1:
					wl.CopySite = cell
					s = append(s, cell)
				case 2:
					wl.WorkLoadName = cell
				case 3:
					wl.BackupType = cell
				case 4:
					wl.VMQty, _ = strconv.Atoi(cell)
				case 5:
					wl.VmdkQty, _ = strconv.Atoi(cell)
				case 6:
					wl.WorkLoadCap, _ = strconv.ParseFloat(cell, 64)
				case 7:
					wl.ChangeRate, _ = strconv.Atoi(cell)
				case 8:
					wl.GrowthPercent, _ = strconv.Atoi(cell)
				case 9:
					wl.BackupWindow, _ = strconv.Atoi(cell)
				case 10:
					wl.ScopeYears, _ = strconv.Atoi(cell)
				case 11:
					wl.Reduction, _ = strconv.Atoi(cell)
				case 12:
					wl.UseReFs = cell
				case 13:
					var check = "perVM"
					if cell == "no" {
						check = "perJob"
					}
					wl.UsePerVM = check
				case 14:
					var check = false
					if cell == "yes" {
						check = true
					}
					wl.CloudEnabled = check
				case 15:
					wl.CloudMove, _ = strconv.Atoi(cell)
				case 16:
					wl.RpsBu, _ = strconv.Atoi(cell)
				case 17:
					wl.BuWeekly, _ = strconv.Atoi(cell)
				case 18:
					wl.BuMonthly, _ = strconv.Atoi(cell)
				case 19:
					wl.BuYearly, _ = strconv.Atoi(cell)
				case 20:
					wl.RpsBuCopy, _ = strconv.Atoi(cell)
				case 21:
					wl.BuCopyWeekly, _ = strconv.Atoi(cell)
				case 22:
					wl.BuCopyMonthly, _ = strconv.Atoi(cell)
				case 23:
					wl.BuCopyYearly, _ = strconv.Atoi(cell)
				}
			}
			wls = append(wls, wl)
		}
	}

	// remove duplicates from sites
	ds := removeDuplicateValues(s)

	var sites []Site

	// create an slice of sites
	for _, item := range ds {
		site := Site{
			SiteName:      item,
			WanSpeed:      1000,
			NetworkSpeed:  1000,
			InternetSpeed: 1000,
		}
		sites = append(sites, site)
	}

	vseInput := VseInput{
		Workload: wls,
		Settings: settings,
		Sites:    sites,
	}

	file, _ := json.MarshalIndent(vseInput, "", " ")

	err = ioutil.WriteFile("vse_input_file.txt", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func removeDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

package debian

import (
	"encoding/json"

	"github.com/TuM0xA-S/bugTracker/load"
	"github.com/TuM0xA-S/bugTracker/types"
)

//Load loads bugs from debian source
func Load() ([]types.BugData, error) {
	data, err := load.Fetch("https://security-tracker.debian.org/tracker/data/json")
	if err != nil {
		return nil, err
	}
	var jd map[string]interface{}
	if err := json.Unmarshal(data, &jd); err != nil {
		return nil, err
	}
	bd := map[string]types.BugData{}
	for pkg := range jd {
		cveMap := jd[pkg].(map[string]interface{})
		for cve := range cveMap{
			v := bd[cve]
			v.CVE = cve
			v.Data = cveMap[cve].(map[string]interface{})
			v.Packages = append(v.Packages, pkg)
			v.Source = "debian"
			bd[cve] = v
		}
	}	
	
	var res []types.BugData
	for _, v := range bd {
		res = append(res, v)
	}

	return res, nil
}

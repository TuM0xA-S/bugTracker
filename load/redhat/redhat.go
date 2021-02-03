package redhat

import (
	"encoding/json"

	"github.com/TuM0xA-S/bugTracker/load"
	"github.com/TuM0xA-S/bugTracker/types"
)

//Load redhat source
func Load() ([]types.BugData, error) {
	data, err := load.Fetch("https://access.redhat.com/labs/securitydataapi/cve.json?per_page=1000000000")
	if err != nil {
		return nil, err
	}

	var jd []map[string]interface{}
	if err := json.Unmarshal(data, &jd); err != nil {
		return nil, err
	}

	var res []types.BugData
	for _, v := range jd {
		bug := types.BugData{}
		bug.Source = "redhat"
		bug.CVE = v["CVE"].(string)
		for _, pkg := range v["affected_packages"].([]interface{}) {
			bug.Packages = append(bug.Packages, pkg.(string))
		}
		bug.Data = v
		res = append(res, bug)
	}

	return res, nil
}

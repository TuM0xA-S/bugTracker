package ubuntu

import (
	"bytes"

	"strings"

	"github.com/TuM0xA-S/bugTracker/load"
	"github.com/TuM0xA-S/bugTracker/types"
	"golang.org/x/net/html"
)

func getCVEList(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, at := range n.Attr {
			if at.Key == "class" && at.Val == "ls-blob" {
				text := n.FirstChild.Data
				if strings.HasPrefix(text, "CVE-") {
					links = append(links, n.FirstChild.Data)
				}
			}
		}
	}
	for cur := n.FirstChild; cur != nil; cur = cur.NextSibling {
		links = getCVEList(links, cur)
	}
	return links
}

func Load() (res []types.BugData, err error) {
	page, err := load.Fetch("https://git.launchpad.net/ubuntu-cve-tracker/tree/active")
	if err != nil {
		return nil, err
	}
	root, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		return nil, err
	}
	list := getCVEList(nil, root)

	for _, cve := range list {
		bug := types.BugData{}
		bug.CVE = cve
		bug.Source = "ubuntu"
		bug.Data = "https://git.launchpad.net/ubuntu-cve-tracker/plain/active/" + cve
		res = append(res, bug)
	}

	page, err = load.Fetch("https://git.launchpad.net/ubuntu-cve-tracker/tree/retired")
	if err != nil {
		return nil, err
	}
	root, err = html.Parse(bytes.NewReader(page))
	if err != nil {
		return nil, err
	}
	list = getCVEList(nil, root)

	for _, cve := range list {
		bug := types.BugData{}
		bug.CVE = cve
		bug.Source = "ubuntu"
		bug.Data = "https://git.launchpad.net/ubuntu-cve-tracker/plain/retired/" + cve
		res = append(res, bug)
	}

	return
}

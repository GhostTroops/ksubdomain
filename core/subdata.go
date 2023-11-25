package core

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed data/subnext.txt
var subnext string

//go:embed data/subdomain.txt
var subdomain string
var DefaultDomainList []string

func GetDefaultSubdomainData() []string {
	if 0 < len(DefaultDomainList) {
		return DefaultDomainList
	}
	reader := bufio.NewScanner(strings.NewReader(subdomain))
	reader.Split(bufio.ScanLines)
	var ret []string
	for reader.Scan() {
		ret = append(ret, reader.Text())
	}
	return ret
}
func GetDefaultSubNextData() []string {
	if 0 < len(DefaultDomainList) {
		return DefaultDomainList
	}
	reader := bufio.NewScanner(strings.NewReader(subnext))
	reader.Split(bufio.ScanLines)
	var ret []string
	for reader.Scan() {
		ret = append(ret, reader.Text())
	}
	return ret
}

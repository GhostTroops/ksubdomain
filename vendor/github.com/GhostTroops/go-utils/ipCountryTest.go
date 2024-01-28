package go_utils

import (
	"embed"
	"net"
	"strings"
	"sync"
)

//go:embed ipdb/*
var IpDbs embed.FS
var (
	CtIps = map[string][]*net.IPNet{}
	ipLc  sync.Mutex
)

func GetEmbedPay(fs embed.FS, s string) []byte {
	if data, err := fs.ReadFile(s); nil == err {
		return data
	}
	return nil
}
func init() {
	if fs, err := IpDbs.ReadDir("ipdb"); nil == err {
		for _, x := range fs {
			if strings.HasSuffix(x.Name(), ".txt") {
				if data := GetEmbedPay(IpDbs, "ipdb/"+x.Name()); nil != data {
					a := strings.Split(strings.TrimSpace(string(data)), "\n")
					k1 := strings.Split(x.Name(), ".")[0]
					var a1 = CtIps[k1]
					for _, x := range a {
						if _, n1, err := net.ParseCIDR(x); nil == err {
							a1 = append(a1, n1)
						}
					}
					CtIps[k1] = a1
				}
			}
		}
	}
}

func CheckIpIsCountry(ip, ctName string) bool {
	ipLc.Lock()
	defer ipLc.Unlock()
	ip1 := net.ParseIP(ip)
	if a, ok := CtIps[strings.ToUpper(ctName)]; ok {
		for _, x := range a {
			if x.Contains(ip1) {
				return true
			}
		}
	}
	return false
}

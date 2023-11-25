package runner

import (
	util "github.com/hktalent/go-utils"
	"strings"
)

func (r *runner) handleResult() {
	go util.DoRunning()
	defer util.CloseLogBigDb()
	var szSkp = "0.0.0.1"
	//log.Println("r.options.Writer len:", len(r.options.Writer))
	for result := range r.recver {
		if -1 < strings.Index(result.Subdomain, szSkp) {
			continue
		}
		for _, x := range result.Answers {
			if x == szSkp {
				return
			}
		}

		var m1 = map[string]interface{}{"ip": result.Answers, "subdomain": result.Subdomain, "tags": "subdomain"}
		//KvCc.KvCc.Put(result.Subdomain, []byte("1"))

		go util.PushLog(&m1)
		for _, out := range r.options.Writer {
			_ = out.WriteDomainResult(result)
		}
		r.printStatus()
	}
}

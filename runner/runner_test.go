package runner

import (
	"context"
	util "github.com/hktalent/go-utils"
	"github.com/hktalent/ksubdomain/core/dns"
	"github.com/hktalent/ksubdomain/core/gologger"
	"github.com/hktalent/ksubdomain/core/options"
	"github.com/hktalent/ksubdomain/runner/outputter"
	"github.com/hktalent/ksubdomain/runner/outputter/output"
	"github.com/hktalent/ksubdomain/runner/processbar"
	"testing"
)

func TestRunner(t *testing.T) {
	process := processbar.ScreenProcess{}
	screenPrinter, _ := output.NewScreenOutput(false)
	domains := []string{"stu.baidu.com", "haokan.baidu.com"}
	_, ns, err := dns.LookupNS("baidu.com", "1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
	domainChanel := make(chan string)
	util.DefaultPool.Submit(func() {
		for _, d := range domains {
			domainChanel <- d
		}
		close(domainChanel)
	})
	opt := &options.Options{
		Rate:        options.Band2Rate("1m"),
		Domain:      domainChanel,
		DomainTotal: 2,
		Resolvers:   options.GetResolvers(""),
		Silent:      false,
		TimeOut:     10,
		Retry:       3,
		Method:      VerifyType,
		DnsType:     "a",
		Writer: []outputter.Output{
			screenPrinter,
		},
		ProcessBar: &process,
		EtherInfo:  options.GetDeviceConfig(),
		SpecialResolvers: map[string][]string{
			"baidu.com": ns,
		},
	}
	opt.Check()
	r, err := New(opt)
	if err != nil {
		gologger.Fatalf(err.Error())
	}
	ctx := context.Background()
	r.RunEnumeration(ctx)
	r.Close()
}

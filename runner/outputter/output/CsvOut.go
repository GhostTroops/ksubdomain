package output

import (
	"encoding/csv"
	"github.com/hktalent/ksubdomain/runner/result"
	"os"
	"strings"
)

type CsvOutImp struct {
	FileOutPut
	IsCsv bool
}

func NewCsvOutImp(filename string, onlyDomain, IsCsv bool) (*CsvOutImp, error) {
	output, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return nil, err
	}
	f := new(CsvOutImp)
	f.output = output
	f.onlyDomain = onlyDomain
	f.IsCsv = IsCsv
	return f, err
}
func (f *CsvOutImp) WriteDomainResult(domain result.Result) error {
	var a []string
	if f.onlyDomain {
		a = append(a, domain.Subdomain)
	} else {
		a = append(a, domain.Subdomain)
		a = append(a, strings.Join(domain.Answers, ","))
	}
	buf := csv.NewWriter(f.output)
	err := buf.Write(a)
	buf.Flush()
	return err
}
func (f *CsvOutImp) Close() {
	f.output.Close()
}

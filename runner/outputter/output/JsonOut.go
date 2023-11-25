package output

import (
	"bufio"
	util "github.com/hktalent/go-utils"
	"github.com/hktalent/ksubdomain/runner/result"
	"os"
)

type JsonOutImp struct {
	FileOutPut
}

func NewJsonOutImp(filename string, onlyDomain bool) (*JsonOutImp, error) {
	output, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return nil, err
	}
	f := new(JsonOutImp)
	f.output = output
	f.onlyDomain = onlyDomain
	return f, err
}
func (f *JsonOutImp) WriteDomainResult(domain result.Result) error {
	buf := bufio.NewWriter(f.output)
	var err error
	if data, err := util.Json.Marshal(domain); nil == err {
		_, err = buf.Write(append(data, []byte("\n")...))
	}
	buf.Flush()
	return err
}
func (f *JsonOutImp) Close() {
	f.output.Close()
}

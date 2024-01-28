package output

import (
	"bufio"
	util "github.com/GhostTroops/go-utils"
	"github.com/GhostTroops/ksubdomain/runner/result"
	"os"
)

type JsonOutImp struct {
	FileOutPut
}

func NewJsonOutImp(filename string, onlyDomain bool) (*JsonOutImp, error) {
	f := new(JsonOutImp)
	if "" != filename {
		output, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			return nil, err
		}
		f.output = output
	}

	f.onlyDomain = onlyDomain
	return f, nil
}
func (f *JsonOutImp) WriteDomainResult(domain result.Result) error {
	var err error
	if nil != f.output {
		buf := bufio.NewWriter(f.output)
		if data, err := util.Json.Marshal(domain); nil == err {
			_, err = buf.Write(append(data, []byte("\n")...))
		}
		buf.Flush()
	}
	return err
}
func (f *JsonOutImp) Close() {
	if nil != f.output {
		f.output.Close()
	}
}

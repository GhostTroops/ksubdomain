package main

import (
	"fmt"
	"github.com/hktalent/ksubdomain/core"
	"testing"
)

func TestLinesInFile(t *testing.T) {
	//if f, err := os.Create("config/subdomain.zip"); nil == err {
	//	w := gzip.NewWriter(f)
	//	if f1, err := os.OpenFile("config/subdomain.txt", os.O_RDWR, 0666); nil == err {
	//		io.Copy(w, f1)
	//		f1.Close()
	//		w.Close()
	//	}
	//}
	if got, err := core.LinesInFile("config/subdomain.txt"); nil == err {
		for _, x := range got {
			fmt.Println(x)
		}
	} else {
		fmt.Println(err)
	}
}

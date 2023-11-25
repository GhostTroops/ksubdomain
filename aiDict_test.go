package main

import (
	util "github.com/hktalent/go-utils"
	"github.com/hktalent/ksubdomain/cmd/ksubdomain"
	"os"
	"testing"
)

func TestGetDict(t *testing.T) {
	os.Setenv("CacheName", ".xxxTest")
	util.DoInitAll()
	ksubdomain.GetDict()
	util.Wg.Wait()
	util.CloseAll()
}

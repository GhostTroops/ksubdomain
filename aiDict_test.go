package main

import (
	"github.com/GhostTroops/ksubdomain/cmd/ksubdomain"
	util "github.com/hktalent/go-utils"
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

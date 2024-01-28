package main

import (
	util "github.com/GhostTroops/go-utils"
	"github.com/GhostTroops/ksubdomain/cmd/ksubdomain"
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

package main

import (
	util "github.com/hktalent/go-utils"
	myCmd "github.com/hktalent/ksubdomain/cmd/ksubdomain"
	"net/http"
	_ "net/http/pprof"
	"os"
)

// go tool pprof -seconds=60 -http=:9999 http://127.0.0.1:6060/debug/pprof/heap
// go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=60
func main() {
	//os.Args = []string{"", "enum", "-d", "huazhu.com", "-o", "/Users/51pwn/huazhu.json", "-j", "-b", "5M"}
	os.RemoveAll(".DbCache")
	defer os.RemoveAll("ksubdomain.yaml")
	util.DoInitAll()
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	myCmd.Main()
	util.Wg.Wait()
	util.CloseAll()
}

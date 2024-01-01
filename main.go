package main

import (
	"embed"
	myCmd "github.com/GhostTroops/ksubdomain/cmd/ksubdomain"
	util "github.com/hktalent/go-utils"
	"io"
	"log"
	_ "net/http/pprof"
	"os"
	"strings"
	"sync"
)

//go:embed config/*
var config embed.FS

// docker run --rm -it -v $PWD:/app go1214 /bin/bash -c "cd /app/; CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags \"-linkmode external -extldflags '-static' -s -w\" -o ksubdomain_linux main.go ; exit"
// go tool pprof -seconds=60 -http=:9999 http://127.0.0.1:6060/debug/pprof/heap
// go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=60
func main() {
	if strings.Contains(strings.Join(os.Args[1:], " "), "--json") {
		os.Setenv("devDebug", "false")
		os.Setenv("ProductMod", "release")
		log.SetOutput(io.Discard)
	}
	os.RemoveAll(".DbCache")
	defer os.RemoveAll("ksubdomain.yaml")
	util.Wg = &sync.WaitGroup{}
	util.DoInit(&config)
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil)
	//}()
	myCmd.Main()
	util.Wg.Wait()
	util.CloseAll()
}

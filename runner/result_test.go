package runner

import (
	"log"
	"net"
	"testing"
)

func Test_runner_handleResult(t *testing.T) {
	oIp := net.ParseIP("0.0.0.1")
	if oIp.IsPrivate() {
		log.Println("not is ok")
	} else {
		log.Println("ok")
	}
}

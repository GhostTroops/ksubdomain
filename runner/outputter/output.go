package outputter

import (
	"github.com/GhostTroops/ksubdomain/runner/result"
)

type Output interface {
	WriteDomainResult(domain result.Result) error
	Close()
}

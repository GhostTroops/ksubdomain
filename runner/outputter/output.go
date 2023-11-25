package outputter

import (
	"github.com/hktalent/ksubdomain/runner/result"
)

type Output interface {
	WriteDomainResult(domain result.Result) error
	Close()
}

package optional_test

import (
	"testing"

	"github.com/reaperhero/instrument/dump"
	"github.com/reaperhero/instrument/stdutil/optional"
)

func TestOptFactory_of(t *testing.T) {
	opt := optional.Of(nil)

	dump.P(opt.OrElseGet(34))
}

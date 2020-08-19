package gosermver

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-courier/semver"
)

const (
	v1 = `1.1.2`
	v2 = `v2.1.3-alpha.1`
)

func Test_Semver(t *testing.T) {

	v := semver.MustParseVersion(v1)
	spew.Dump(v)

	v, err := semver.ParseVersion(v2)
	if err != nil {
		panic(err)
	}
	spew.Dump(v)
}

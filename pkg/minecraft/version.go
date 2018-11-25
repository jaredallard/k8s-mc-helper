package minecraft

import (
	"github.com/Masterminds/semver"
)

// Parse will parse a minecraft version or return an error if it's invalud
func Parse(version string) (*Version, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return nil, err
	}

	return &Version{
		Major: v.Major(),
		Minor: v.Minor(),
		Patch: v.Patch(),
	}, nil
}

// GTE checks if version base is GTE to compare
func GTE(base, compare Version) bool {
	// 1.7.2 -> 172
	bs := (base.Major * 100) + (base.Minor * 10) + base.Patch
	bc := (compare.Major * 100) + (compare.Minor * 10) + compare.Patch

	if bs >= bc {
		return true
	}

	return false
}

package stank_test

import (
	"testing"

	"github.com/mcandre/stank"
)

func TestVersion(t *testing.T) {
	if stank.Version == "" {
		t.Errorf("Expected stank version to be non-blank")
	}
}

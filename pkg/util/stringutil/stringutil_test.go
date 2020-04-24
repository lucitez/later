package stringutil_test

import (
	"later/pkg/util/stringutil"
	"testing"
)

func TestRandomNInt(t *testing.T) {
	rand := stringutil.RandomNInt(6)

	if len(rand) != 6 {
		t.Error(rand)
	}
}

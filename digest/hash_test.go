package digest

import "testing"

func Test_makeDigest(t *testing.T) {
	digest := makeDigest(1561379366, 308630752, "secret")
	if digest != "2c717930b6098365d85fc4bee9b6174d11cd4fe2d0e3753d61ad8004bbc9fcc4" {
		t.Error("invalid hash")
	}
}

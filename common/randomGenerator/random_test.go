package randomGenerator

import (
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	//AES key generate
	randomString,err := GenerateRandomString(32)
	if err != nil {
		t.Fatalf("Generated random string error: %+v",err.Error())
	}
	t.Logf("Generated random string: %+v",randomString)
}
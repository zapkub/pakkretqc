package frontend

import (
	"errors"
	"testing"

	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

func TestLogin(t *testing.T) {

	if !errors.Is(almsdk.InvalidCredential, almsdk.InvalidCredential) {
		t.Fatal("invalid")
	}

}

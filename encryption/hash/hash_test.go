/*
* BSD 3-Clause License
* Copyright © 2022. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package hash

import (
	"testing"
)

func TestEval(t *testing.T) {
	hash := GenerateHash("sha256", "action.12345.2208.hola")
	if hash != "d462a15092063b5ded3d045ed4e199ec5741b43baeed9a698fe173c1ebe32eb7" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", hash, "d462a15092063b5ded3d045ed4e199ec5741b43baeed9a698fe173c1ebe32eb7")
	}
}

package feda

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Log("Testing String()")

	fedaStr := ""
	for cont := int(SPACE); cont <= int(REVERSE_INITIATOR); cont++ {
		t.Logf("%d: %s", cont, Fid(cont).String())
		fedaStr += Fid(cont).String()
	}
	t.Logf("Feda: " + fedaStr)

}

func TestUintToFid(t *testing.T) {
	t.Log("Testing UintToFid()")

	for val := uint(20); val <= 40; val++ {
		feda := UintToFeda(val)

		fedaStr := ""
		for _, fid := range feda {
			fedaStr += fid.String()
		}

		newVal, err := FedaToUint(feda)
		if err != nil {
			t.Error(err)
			return
		}

		if val != newVal {
			t.Errorf("Error - Value: %d -> %s -> %d", val, fedaStr, newVal)
			return
		}

		t.Log("Value: ", val, " -> ", fedaStr, " -> ", newVal)
	}
}

/*
func TestString(t *testing.T) {
	t.Log("Testing String()")

	aicmeStr := ">"
	for cont := 1; cont <= int(Nuin); cont++ {
		t.Logf("%d: %s", cont, Aicme1(cont).String())
		aicmeStr += Aicme1(cont).String()
	}
	t.Logf("Aicme1: " + aicmeStr)

	aicmeStr = ">"
	for cont := 1; cont <= int(Cert); cont++ {
		t.Logf("%d: %s", cont, Aicme2(cont).String())
		aicmeStr += Aicme2(cont).String()
	}
	t.Logf("Aicme2: " + aicmeStr)

	aicmeStr = ">"
	for cont := 1; cont <= int(Ruis); cont++ {
		t.Logf("%d: %s", cont, Aicme3(cont).String())
		aicmeStr += Aicme3(cont).String()
	}
	t.Logf("Aicme3: " + aicmeStr)

	aicmeStr = ">"
	for cont := 1; cont <= int(Idad); cont++ {
		t.Logf("%d: %s", cont, Aicme4(cont).String())
		aicmeStr += Aicme4(cont).String()
	}
	t.Logf("Aicme4: " + aicmeStr)

	aicmeStr = ">"
	for cont := 0; cont <= int(Peith); cont++ {
		t.Logf("%d: %s", cont, Forfeda(cont).String())
		aicmeStr += Forfeda(cont).String()
	}
	t.Logf("Forfeda: " + aicmeStr)
}
*/

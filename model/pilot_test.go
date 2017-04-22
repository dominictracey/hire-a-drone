package model

import "testing"

func TestGCE(t *testing.T) {
	m := newPilot()
	m.Address = "123 Kiwi Court, Aukland, NZ"

	if m == nil {
		t.Error("Test not running on GCE, but error does not indicate that fact.")
	} else {
		t.Log("Ping! All is Well...")
	}
}

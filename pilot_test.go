package main

import (
	"log"
	"testing"
)

func TestPilot(t *testing.T) {
	m := NewPilot()
	m.Address = "123 Kiwi Court, Aukland, NZ"

	if m == nil {
		t.Error("Test not running on GCE, but error does not indicate that fact.")
	} else {
		log.Println("I am here!")
		t.Log("Ping! All is Well...")
	}
}

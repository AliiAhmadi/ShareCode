package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2023, 12, 3, 5, 0, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != "03 Dec 2023 at 05:00" {
		t.Errorf("walt: %s. got: %s", "3 Dec 2023 at 5:00", hd)
	}
}

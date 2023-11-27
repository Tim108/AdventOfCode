package main

import "testing"

func TestElapseSecondForReindeer(t *testing.T) {
	reindeer := Reindeer{"testReindeer", 14, 10, 127, 10, 0, 0, 0}
	for i := 0; i < 1000; i++ {
		reindeer.elapseSecond()
	}
	want := 1120
	if reindeer.distance != want {
		t.Fatalf("Reindeer %v should have traveled %d meters, but was %d", reindeer, want, reindeer.distance)
	}
}

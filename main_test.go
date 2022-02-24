package main

import "testing"

func TestUpgradelevel(t *testing.T) {
	fr := []filledRole{
		{&contributor{
			name:   "un",
			skills: map[string]int{"go": 5},
			delay:  0,
		}, role{
			skill: "go",
			level: 5,
		}},
	}
	upgradeLevel(fr)
	actual := fr[0].contributor.skills["go"]
	expected := 6
	if actual != expected {
		t.Errorf("got %d, want %d", actual, expected)
	}
}

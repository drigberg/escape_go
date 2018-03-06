package main

import (
	"escape_go/helpers"
	"testing"
)

func TestGetLocations(t *testing.T) {
	locations := helpers.GetLocations("./test/locations.yaml")

	if len(locations) != 2 {
		t.Errorf("Should have loaded two locations")
	}

	if locations["left"].Description != "to the left" {
		t.Errorf("Want %q, got %q", "to the left", locations["left"].Description)
	}

	if locations["left"].Query != "go right?" {
		t.Errorf("Want %q, got %q", "go right?", locations["left"].Query)
	}

	if len(locations["left"].Options) != 1 {
		t.Errorf("Want %q, got %q", 1, len(locations["left"].Options))
	}

	if locations["right"].Description != "to the right" {
		t.Errorf("Want %q, got %q", "to the right", locations["right"].Description)
	}

	if locations["right"].Query != "go left?" {
		t.Errorf("Want %q, got %q", "go left?", locations["right"].Query)
	}

	if len(locations["right"].Options) != 2 {
		t.Errorf("Want %q, got %q", 2, len(locations["right"].Options))
	}
}
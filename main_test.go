package main

import "testing"

func TestNoAccessFilename(t *testing.T) {
	tests := map[string]string{
		"sccmfiles.txt":        "sccmfiles_noaccess.txt",
		"inventory":            "inventory_noaccess",
		"loot/sccmfiles.index": "loot/sccmfiles_noaccess.index",
	}

	for input, want := range tests {
		if got := noAccessFilename(input); got != want {
			t.Errorf("noAccessFilename(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestMatchingReference(t *testing.T) {
	metadata := []byte("[User]\r\n P0100001.1 =\r\nP0100002.1=\r\n")
	entries := map[string]struct{}{"p0100001.1": {}}

	got, ok := matchingReference(metadata, entries)
	if !ok {
		t.Fatal("matchingReference did not find a reference")
	}
	if got != "P0100001.1" {
		t.Errorf("matchingReference() = %q, want %q", got, "P0100001.1")
	}
}

func TestMatchingReferenceIgnoresNonMatchingMetadata(t *testing.T) {
	metadata := []byte("[User]\nP0100002.1=\n")
	entries := map[string]struct{}{"p0100001.1": {}}

	if got, ok := matchingReference(metadata, entries); ok || got != "" {
		t.Errorf("matchingReference() = %q, %t; want no match", got, ok)
	}
}

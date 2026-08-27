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
	metadata := []byte("[File]\r\n secret.ps1.INI = 2\r\nother.xml.INI=1\r\n")
	entries := map[string]struct{}{"secret.ps1.ini": {}}

	got, ok := matchingReference(metadata, entries)
	if !ok {
		t.Fatal("matchingReference did not find a reference")
	}
	if got != "secret.ps1.INI" {
		t.Errorf("matchingReference() = %q, want %q", got, "secret.ps1.INI")
	}
}

func TestMatchingReferenceIgnoresNonMatchingMetadata(t *testing.T) {
	metadata := []byte("[File]\nother.xml.INI=1\n")
	entries := map[string]struct{}{"secret.ps1.ini": {}}

	if got, ok := matchingReference(metadata, entries); ok || got != "" {
		t.Errorf("matchingReference() = %q, %t; want no match", got, ok)
	}
}

package cmd

import "testing"

func TestEncodeURL(t *testing.T) {
	raw := "name=\"Lukas Herman\"&age=23"
	expected := "name=%22Lukas%20Herman%22&age=23"

	encoded := encodeURL(raw)
	if encoded != expected {
		t.Error("expected:", expected, "got:", encoded)
	}
}

func TestDecodeURL(t *testing.T) {
	encoded := "name=%22Lukas%20Herman%22&age=23"
	expected := "name=\"Lukas Herman\"&age=23"

	raw, _ := decodeURL(encoded)
	if raw != expected {
		t.Error("expected:", expected, "got:", raw)
	}
}

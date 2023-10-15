package mp3_test

import (
	"testing"

	"github.com/danduggan98/oscillot-go/audio/mp3"
)

func TestParseHeader(t *testing.T) {
	bits := 0xFFFBA040 // Example from MP3 Wikipedia article
	h := mp3.ParseHeader(bits)

	for _, v := range h.SyncWord {
		if !v {
			t.Fatalf("sync word should be all 1s but got %v", h.SyncWord)
		}
	}
	if !h.Version {
		t.Fatalf("version bit should be 1 but got 0")
	}
	if h.Layer[0] || !h.Layer[1] {
		t.Fatalf("layer should be 01 but got %v", h.Layer)
	}
	if !h.ErrorProtection {
		t.Fatalf("error protection bit should be 1 but got 0")
	}
	if !h.BitRate[0] || h.BitRate[1] || !h.BitRate[2] || h.BitRate[3] {
		t.Fatalf("bit rate should be 1010 but got %v", h.BitRate)
	}
	for _, v := range h.Frequency {
		if v {
			t.Fatalf("frequency should be all 0s but got %v", h.Frequency)
		}
	}
	if h.PaddingBit {
		t.Fatalf("padding bit should be 0 but got 1")
	}
	if h.PrivacyBit {
		t.Fatalf("version bit should be 0 but got 1")
	}
	if h.Mode[0] || !h.Mode[1] {
		t.Fatalf("mode should be 01 but got %v", h.Mode)
	}
	for _, v := range h.ModeExtension {
		if v {
			t.Fatalf("mode extension should be all 0s but got %v", h.ModeExtension)
		}
	}
	if h.Copyrighted {
		t.Fatalf("copyright bit should be 0 but got 1")
	}
	if h.PrivacyBit {
		t.Fatalf("privacy bit should be 0 but got 1")
	}
	for _, v := range h.Emphasis {
		if v {
			t.Fatalf("emphasis should be all 0s but got %v", h.Emphasis)
		}
	}
}

package mp3_test

import (
	"testing"

	"github.com/danduggan98/oscillot-go/audio/mp3"
)

var bits uint32 = 0xFFFBA040 // Example from MP3 Wikipedia article
var h *mp3.Header = mp3.ParseHeader(bits)

func TestParseHeader(t *testing.T) {
	if h.SyncWord != 0b0111111111111 {
		t.Fatalf("sync word should be 111111111111 but was %b", h.SyncWord)
	}
	if h.Version != 0b1 {
		t.Fatalf("version bit should be 1 but got 0")
	}
	if h.Layer != 0b01 {
		t.Fatalf("layer should be 01 but got %b", h.Layer)
	}
	if h.ErrorProtection != 0b1 {
		t.Fatalf("error protection bit should be 1 but got 0")
	}
	if h.BitRate != 0b1010 {
		t.Fatalf("bit rate should be 1010 but got %b", h.BitRate)
	}
	if h.Frequency != 0b00 {
		t.Fatalf("frequency should be 00 but got %b", h.Frequency)
	}
	if h.PaddingBit != 0b0 {
		t.Fatalf("padding bit should be 0 but got 1")
	}
	if h.PrivacyBit != 0b0 {
		t.Fatalf("version bit should be 0 but got 1")
	}
	if h.Mode != 0b01 {
		t.Fatalf("mode should be 01 but got %b", h.Mode)
	}
	if h.ModeExtension != 0b00 {
		t.Fatalf("mode extension should be 00 but got %b", h.ModeExtension)
	}
	if h.Copyrighted != 0b0 {
		t.Fatalf("copyright bit should be 0 but got 1")
	}
	if h.PrivacyBit != 0b0 {
		t.Fatalf("privacy bit should be 0 but got 1")
	}
	if h.Emphasis != 0b00 {
		t.Fatalf("emphasis should be 00 but got %b", h.Emphasis)
	}
}

func TestCalculateBitRate(t *testing.T) {
	bitRate := h.CalculateBitrate()
	if bitRate != 160 {
		t.Fatalf("bit rate should be 160 but got %d", bitRate)
	}
}

func TestCalculateSampleRate(t *testing.T) {
	sampleRate := h.CalculateSampleRate()
	if sampleRate != 44100 {
		t.Fatalf("sample rate should be 44100 Hz but got %d", sampleRate)
	}
}

func TestCalculateModeExtension(t *testing.T) {
	iStereo, msStereo := h.CalculateModeExtension()
	if iStereo || msStereo {
		t.Fatalf("mode extension should be false false but got %v %v", iStereo, msStereo)
	}
}

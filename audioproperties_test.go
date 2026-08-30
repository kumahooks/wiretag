package wiretag

import (
	"errors"
	"testing"
)

func TestAudioProperties(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath    string
		wantLength       int
		wantBitrate      int
		wantSampleRate   int
		wantAudioChannel int
	}{
		{boaFilePath, 1, 446, 44100, 2},
		{venetianFilePath, 1, 446, 44100, 2},
		{wiynFilePath, 1, 45, 44100, 1},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotLength, err := audioFile.AudioLength()
		if err != nil {
			t.Fatalf("AudioLength(%s): %v", testCase.audioFilePath, err)
		}
		if gotLength != testCase.wantLength {
			t.Fatalf("AudioLength(%s): got %d, want %d", testCase.audioFilePath, gotLength, testCase.wantLength)
		}

		gotBitrate, err := audioFile.AudioBitrate()
		if err != nil {
			t.Fatalf("AudioBitrate(%s): %v", testCase.audioFilePath, err)
		}
		if gotBitrate != testCase.wantBitrate {
			t.Fatalf("AudioBitrate(%s): got %d, want %d", testCase.audioFilePath, gotBitrate, testCase.wantBitrate)
		}

		gotSampleRate, err := audioFile.AudioSampleRate()
		if err != nil {
			t.Fatalf("AudioSampleRate(%s): %v", testCase.audioFilePath, err)
		}
		if gotSampleRate != testCase.wantSampleRate {
			t.Fatalf(
				"AudioSampleRate(%s): got %d, want %d",
				testCase.audioFilePath,
				gotSampleRate,
				testCase.wantSampleRate,
			)
		}

		gotChannels, err := audioFile.AudioChannels()
		if err != nil {
			t.Fatalf("AudioChannels(%s): %v", testCase.audioFilePath, err)
		}
		if gotChannels != testCase.wantAudioChannel {
			t.Fatalf(
				"AudioChannels(%s): got %d, want %d",
				testCase.audioFilePath,
				gotChannels,
				testCase.wantAudioChannel,
			)
		}
	}
}

func TestAudioPropertiesClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	if got, err := audioFile.AudioLength(); !errors.Is(err, ErrFileClosed) || got != 0 {
		t.Fatalf("AudioLength(closed): got (%d, %v), want (0, ErrFileClosed)", got, err)
	}

	if got, err := audioFile.AudioBitrate(); !errors.Is(err, ErrFileClosed) || got != 0 {
		t.Fatalf("AudioBitrate(closed): got (%d, %v), want (0, ErrFileClosed)", got, err)
	}

	if got, err := audioFile.AudioSampleRate(); !errors.Is(err, ErrFileClosed) || got != 0 {
		t.Fatalf("AudioSampleRate(closed): got (%d, %v), want (0, ErrFileClosed)", got, err)
	}

	if got, err := audioFile.AudioChannels(); !errors.Is(err, ErrFileClosed) || got != 0 {
		t.Fatalf("AudioChannels(closed): got (%d, %v), want (0, ErrFileClosed)", got, err)
	}
}

func TestAudioPropertiesNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.AudioLength(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("AudioLength(nil): got %v, want ErrFileClosed", err)
	}

	if _, err := file.AudioBitrate(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("AudioBitrate(nil): got %v, want ErrFileClosed", err)
	}

	if _, err := file.AudioSampleRate(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("AudioSampleRate(nil): got %v, want ErrFileClosed", err)
	}

	if _, err := file.AudioChannels(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("AudioChannels(nil): got %v, want ErrFileClosed", err)
	}
}

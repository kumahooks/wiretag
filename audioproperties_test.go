package wiretag

import (
	"errors"
	"testing"
)

func TestAudioProperties(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath  string
		wantLength     int
		wantBitrate    int
		wantSampleRate int
		wantChannels   int
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
		if gotChannels != testCase.wantChannels {
			t.Fatalf(
				"AudioChannels(%s): got %d, want %d",
				testCase.audioFilePath,
				gotChannels,
				testCase.wantChannels,
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

	for _, testCase := range []struct {
		getter func() (int, error)
		name   string
	}{
		{audioFile.AudioLength, "AudioLength"},
		{audioFile.AudioBitrate, "AudioBitrate"},
		{audioFile.AudioSampleRate, "AudioSampleRate"},
		{audioFile.AudioChannels, "AudioChannels"},
	} {
		gotValue, err := testCase.getter()
		if !errors.Is(err, ErrFileClosed) {
			t.Fatalf("%s(closed): got %v, want ErrFileClosed", testCase.name, err)
		}
		if gotValue != 0 {
			t.Fatalf("%s(closed): got %d, want 0", testCase.name, gotValue)
		}
	}
}

func TestAudioPropertiesNilFile(t *testing.T) {
	var file *AudioFile

	for _, testCase := range []struct {
		getter func() (int, error)
		name   string
	}{
		{file.AudioLength, "AudioLength"},
		{file.AudioBitrate, "AudioBitrate"},
		{file.AudioSampleRate, "AudioSampleRate"},
		{file.AudioChannels, "AudioChannels"},
	} {
		if _, err := testCase.getter(); !errors.Is(err, ErrFileClosed) {
			t.Fatalf("%s(nil): got %v, want ErrFileClosed", testCase.name, err)
		}
	}
}

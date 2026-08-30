package wiretag

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	for _, path := range []string{completeBoaFilePath, completeVenetianFilePath, completeWiynFilePath} {
		audioFile, err := Open(path)
		if err != nil {
			t.Fatalf("Open(%s): %v", path, err)
		}
		t.Cleanup(audioFile.Close)

		if audioFile.handle == nil {
			t.Fatalf("Open(%s): handle not opened", path)
		}

		audioFile.Close()

		if audioFile.handle != nil {
			t.Fatalf("Close(%s): handle still open", path)
		}
	}
}

func TestOpenMissingFile(t *testing.T) {
	_, err := Open("testdata/does-not-exist.flac")
	if !errors.Is(err, ErrInvalid) {
		t.Fatalf("Open(missing): got %v, want ErrInvalid", err)
	}
}

func TestOpenEmptyPath(t *testing.T) {
	_, err := Open("")
	if !errors.Is(err, ErrInvalid) {
		t.Fatalf("Open(empty): got %v, want ErrInvalid", err)
	}
}

func TestOpenInvalidFile(t *testing.T) {
	invalidAudioFile := filepath.Join(t.TempDir(), "very-much-not-audio.flac")
	if err := os.WriteFile(invalidAudioFile, []byte("this is definitely not an audio file"), 0o644); err != nil {
		t.Fatalf("write invalid file: %v", err)
	}

	_, err := Open(invalidAudioFile)
	if !errors.Is(err, ErrInvalid) {
		t.Fatalf("Open(invalid): got %v, want ErrInvalid", err)
	}
}

func TestCloseIdempotent(t *testing.T) {
	file, err := Open(completeBoaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	file.Close()
	file.Close()

	if file.handle != nil || file.tag != nil || file.audioProperties != nil {
		t.Fatal("Close: handle, tag or audioProperties not nilled")
	}
}

func TestIsValid(t *testing.T) {
	t.Run("opened valid file", func(t *testing.T) {
		for _, path := range []string{completeBoaFilePath, completeVenetianFilePath, completeWiynFilePath} {
			audioFile, err := Open(path)
			if err != nil {
				t.Fatalf("Open(%s): %v", path, err)
			}
			t.Cleanup(audioFile.Close)

			if !audioFile.IsValid() {
				t.Errorf("IsValid(%s): got false, want true", path)
			}
		}
	})

	t.Run("closed file", func(t *testing.T) {
		audioFile, err := Open(completeBoaFilePath)
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		audioFile.Close()

		if audioFile.IsValid() {
			t.Error("IsValid(closed): got true, want false")
		}
	})

	t.Run("nil receiver", func(t *testing.T) {
		var audioFile *AudioFile

		if audioFile.IsValid() {
			t.Error("IsValid(nil): got true, want false")
		}
	})
}

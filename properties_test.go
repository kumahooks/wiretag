package wiretag

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func expectedProperties(audioFilePath string) (map[string][]string, error) {
	base := filepath.Base(audioFilePath)
	goldenPath := filepath.Join("testdata", "properties", base[:len(base)-len(filepath.Ext(base))]+".json")

	goldenData, err := os.ReadFile(goldenPath)
	if err != nil {
		return nil, err
	}

	wantProperties := map[string][]string{}
	if err := json.Unmarshal(goldenData, &wantProperties); err != nil {
		return nil, err
	}

	return wantProperties, nil
}

func TestPropertiesMatchesGolden(t *testing.T) {
	for _, audioFilePath := range []string{boaFilePath, venetianFilePath, wiynFilePath} {
		audioFile, err := Open(audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotProperties, err := audioFile.Properties()
		if err != nil {
			t.Fatalf("Properties(%s): %v", audioFilePath, err)
		}

		wantProperties, err := expectedProperties(audioFilePath)
		if err != nil {
			t.Fatalf("properties golden for %s: %v", audioFilePath, err)
		}

		if len(gotProperties) != len(wantProperties) {
			t.Fatalf("Properties(%s): got %d keys, want %d", audioFilePath, len(gotProperties), len(wantProperties))
		}

		for propertyKey, wantValues := range wantProperties {
			gotValues, ok := gotProperties[propertyKey]
			if !ok {
				t.Fatalf("Properties(%s): missing key %s", audioFilePath, propertyKey)
			}

			if len(gotValues) != len(wantValues) {
				t.Fatalf(
					"Properties(%s)[%s]: got %d values, want %d",
					audioFilePath,
					propertyKey,
					len(gotValues),
					len(wantValues),
				)
			}

			for i, wantValue := range wantValues {
				if gotValues[i] != wantValue {
					t.Fatalf(
						"Properties(%s)[%s][%d]: got %q, want %q",
						audioFilePath,
						propertyKey,
						i,
						gotValues[i],
						wantValue,
					)
				}
			}
		}
	}
}

func TestPropertiesClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	audioFileProperties, err := audioFile.Properties()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Properties(closed): got %v, want ErrFileClosed", err)
	}
	if audioFileProperties != nil {
		t.Fatalf("Properties(closed): got %v, want nil map", audioFileProperties)
	}
}

func TestPropertiesNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.Properties(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Properties(nil): got %v, want ErrFileClosed", err)
	}
}

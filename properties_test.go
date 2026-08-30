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

func TestPropertyValues(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		propertyKey   string
		wantValues    []string
	}{
		{boaFilePath, "ARTIST", []string{"bôa"}},
		{boaFilePath, "PERFORMER", []string{
			"Ben Henderson (percussion)",
			"Jasmine Rodgers (percussion)",
			"Paul Turrell (percussion)",
			"Ben Henderson (saxophone)",
			"Paul Turrell (piano)",
			"Paul Turrell (guitar)",
			"Paul Turrell (keyboard)",
		}},
		{boaFilePath, "MISSING_KEY", []string{}},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotValues, err := audioFile.PropertyValues(testCase.propertyKey)
		if err != nil {
			t.Fatalf("PropertyValues(%s, %s): %v", testCase.audioFilePath, testCase.propertyKey, err)
		}

		if len(gotValues) != len(testCase.wantValues) {
			t.Fatalf(
				"PropertyValues(%s, %s): got %d values, want %d",
				testCase.audioFilePath,
				testCase.propertyKey,
				len(gotValues),
				len(testCase.wantValues),
			)
		}

		for i, wantValue := range testCase.wantValues {
			if gotValues[i] != wantValue {
				t.Fatalf(
					"PropertyValues(%s, %s)[%d]: got %q, want %q",
					testCase.audioFilePath,
					testCase.propertyKey,
					i,
					gotValues[i],
					wantValue,
				)
			}
		}
	}
}

func TestPropertyValuesCaseInsensitive(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(audioFile.Close)

	gotValues, err := audioFile.PropertyValues("artist")
	if err != nil {
		t.Fatalf("PropertyValues(lowercase): %v", err)
	}
	if len(gotValues) != 1 || gotValues[0] != "bôa" {
		t.Fatalf("PropertyValues(lowercase): got %v, want [bôa]", gotValues)
	}
}

func TestPropertyValuesClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotValues, err := audioFile.PropertyValues("ARTIST")
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("PropertyValues(closed): got %v, want ErrFileClosed", err)
	}
	if gotValues != nil {
		t.Fatalf("PropertyValues(closed): got %v, want nil slice", gotValues)
	}
}

func TestPropertyValuesNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.PropertyValues("ARTIST"); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("PropertyValues(nil): got %v, want ErrFileClosed", err)
	}
}

func TestPropertyKeys(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantTotalKeys int
	}{
		{boaFilePath, 36},
		{venetianFilePath, 60},
		{wiynFilePath, 8},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotKeys, err := audioFile.PropertyKeys()
		if err != nil {
			t.Fatalf("PropertyKeys(%s): %v", testCase.audioFilePath, err)
		}
		if len(gotKeys) != testCase.wantTotalKeys {
			t.Fatalf(
				"PropertyKeys(%s): got %d keys %v, want %d",
				testCase.audioFilePath,
				len(gotKeys),
				gotKeys,
				testCase.wantTotalKeys,
			)
		}
	}
}

func TestPropertyKeysMatchGolden(t *testing.T) {
	for _, audioFilePath := range []string{boaFilePath, venetianFilePath, wiynFilePath} {
		audioFile, err := Open(audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotKeys, err := audioFile.PropertyKeys()
		if err != nil {
			t.Fatalf("PropertyKeys(%s): %v", audioFilePath, err)
		}

		wantProperties, err := expectedProperties(audioFilePath)
		if err != nil {
			t.Fatalf("properties golden for %s: %v", audioFilePath, err)
		}

		if len(gotKeys) != len(wantProperties) {
			t.Fatalf("PropertyKeys(%s): got %d keys, want %d", audioFilePath, len(gotKeys), len(wantProperties))
		}

		for _, gotKey := range gotKeys {
			if _, ok := wantProperties[gotKey]; !ok {
				t.Fatalf("PropertyKeys(%s): got key %s not in golden", audioFilePath, gotKey)
			}
		}
	}
}

func TestPropertyKeysClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotKeys, err := audioFile.PropertyKeys()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("PropertyKeys(closed): got %v, want ErrFileClosed", err)
	}
	if gotKeys != nil {
		t.Fatalf("PropertyKeys(closed): got %v, want nil slice", gotKeys)
	}
}

func TestPropertyKeysNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.PropertyKeys(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("PropertyKeys(nil): got %v, want ErrFileClosed", err)
	}
}

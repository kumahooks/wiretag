package wiretag

import (
	"errors"
	"testing"
)

func TestTitle(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantTitle     string
	}{
		{boaFilePath, "Fool"},
		{venetianFilePath, "Kétsarkú mozgalom"},
		{wiynFilePath, "we met a long long time ago."},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotTitle, err := audioFile.Title()
		if err != nil {
			t.Fatalf("Title(%s): %v", testCase.audioFilePath, err)
		}
		if gotTitle != testCase.wantTitle {
			t.Fatalf("Title(%s): got %q, want %q", testCase.audioFilePath, gotTitle, testCase.wantTitle)
		}
	}
}

func TestTitleClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotTitle, err := audioFile.Title()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Title(closed): got %v, want ErrFileClosed", err)
	}
	if gotTitle != "" {
		t.Fatalf("Title(closed): got %q, want empty string", gotTitle)
	}
}

func TestTitleNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.Title(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Title(nil): got %v, want ErrFileClosed", err)
	}
}

func TestArtist(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantArtist    string
	}{
		{boaFilePath, "bôa"},
		{venetianFilePath, "Venetian Snares"},
		{wiynFilePath, "What is Your Name?"},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotArtist, err := audioFile.Artist()
		if err != nil {
			t.Fatalf("Artist(%s): %v", testCase.audioFilePath, err)
		}
		if gotArtist != testCase.wantArtist {
			t.Fatalf("Artist(%s): got %q, want %q", testCase.audioFilePath, gotArtist, testCase.wantArtist)
		}
	}
}

func TestArtistClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotArtist, err := audioFile.Artist()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Artist(closed): got %v, want ErrFileClosed", err)
	}
	if gotArtist != "" {
		t.Fatalf("Artist(closed): got %q, want empty string", gotArtist)
	}
}

func TestArtistNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.Artist(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Artist(nil): got %v, want ErrFileClosed", err)
	}
}

func TestAlbum(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantAlbum     string
	}{
		{boaFilePath, "The Race of a Thousand Camels"},
		{venetianFilePath, "Rossz csillag alatt született"},
		{wiynFilePath, "beyond old names; everyone's song."},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotAlbum, err := audioFile.Album()
		if err != nil {
			t.Fatalf("Album(%s): %v", testCase.audioFilePath, err)
		}
		if gotAlbum != testCase.wantAlbum {
			t.Fatalf("Album(%s): got %q, want %q", testCase.audioFilePath, gotAlbum, testCase.wantAlbum)
		}
	}
}

func TestAlbumClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotAlbum, err := audioFile.Album()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Album(closed): got %v, want ErrFileClosed", err)
	}
	if gotAlbum != "" {
		t.Fatalf("Album(closed): got %q, want empty string", gotAlbum)
	}
}

func TestAlbumNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.Album(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Album(nil): got %v, want ErrFileClosed", err)
	}
}

func TestComment(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantComment   string
	}{
		{boaFilePath, "freedbID : 8A0AB60B"},
		{venetianFilePath, "Remember to support artists."},
		{wiynFilePath, "Visit https://whatisyourname.bandcamp.com"},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotComment, err := audioFile.Comment()
		if err != nil {
			t.Fatalf("Comment(%s): %v", testCase.audioFilePath, err)
		}
		if gotComment != testCase.wantComment {
			t.Fatalf("Comment(%s): got %q, want %q", testCase.audioFilePath, gotComment, testCase.wantComment)
		}
	}
}

func TestCommentClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotComment, err := audioFile.Comment()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Comment(closed): got %v, want ErrFileClosed", err)
	}
	if gotComment != "" {
		t.Fatalf("Comment(closed): got %q, want empty string", gotComment)
	}
}

func TestCommentNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.Comment(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Comment(nil): got %v, want ErrFileClosed", err)
	}
}

func TestGenre(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantGenre     string
	}{
		{boaFilePath, "misc"},
		{venetianFilePath, "Electronic"},
		{wiynFilePath, ""},
	} {
		audioFile, err := Open(testCase.audioFilePath)
		if err != nil {
			t.Fatalf("Open(%s): %v", testCase.audioFilePath, err)
		}
		t.Cleanup(audioFile.Close)

		gotGenre, err := audioFile.Genre()
		if err != nil {
			t.Fatalf("Genre(%s): %v", testCase.audioFilePath, err)
		}
		if gotGenre != testCase.wantGenre {
			t.Fatalf("Genre(%s): got %q, want %q", testCase.audioFilePath, gotGenre, testCase.wantGenre)
		}
	}
}

func TestGenreClosedFile(t *testing.T) {
	audioFile, err := Open(boaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	gotGenre, err := audioFile.Genre()
	if !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Genre(closed): got %v, want ErrFileClosed", err)
	}
	if gotGenre != "" {
		t.Fatalf("Genre(closed): got %q, want empty string", gotGenre)
	}
}

func TestGenreNilFile(t *testing.T) {
	var file *AudioFile
	if _, err := file.Genre(); !errors.Is(err, ErrFileClosed) {
		t.Fatalf("Genre(nil): got %v, want ErrFileClosed", err)
	}
}

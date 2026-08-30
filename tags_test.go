package wiretag

import (
	"errors"
	"testing"
)

func TestTags(t *testing.T) {
	for _, testCase := range []struct {
		audioFilePath string
		wantTitle     string
		wantArtist    string
		wantAlbum     string
		wantComment   string
		wantGenre     string
		wantYear      int
		wantTrack     int
	}{
		{
			audioFilePath: completeBoaFilePath,
			wantTitle:     "Fool",
			wantArtist:    "bôa",
			wantAlbum:     "The Race of a Thousand Camels",
			wantComment:   "freedbID : 8A0AB60B",
			wantGenre:     "misc",
			wantYear:      1998,
			wantTrack:     1,
		},
		{
			audioFilePath: completeVenetianFilePath,
			wantTitle:     "Kétsarkú mozgalom",
			wantArtist:    "Venetian Snares",
			wantAlbum:     "Rossz csillag alatt született",
			wantComment:   "Remember to support artists.",
			wantGenre:     "Electronic",
			wantYear:      2005,
			wantTrack:     10,
		},
		{
			audioFilePath: completeWiynFilePath,
			wantTitle:     "we met a long long time ago.",
			wantArtist:    "What is Your Name?",
			wantAlbum:     "beyond old names; everyone's song.",
			wantComment:   "Visit https://whatisyourname.bandcamp.com",
			wantGenre:     "",
			wantYear:      2023,
			wantTrack:     10,
		},
		{
			audioFilePath: completeTogawaFilePath,
			wantTitle:     "好き好き大好き",
			wantArtist:    "戸川純",
			wantAlbum:     "わたしが鳴こうホトトギス",
			wantComment:   "Virgin Babylon Records - VBR-038",
			wantGenre:     "Avant-Garde Pop",
			wantYear:      2016,
			wantTrack:     2,
		},
		{
			audioFilePath: missingTogawaFilePath,
			wantTitle:     "",
			wantArtist:    "",
			wantAlbum:     "",
			wantComment:   "",
			wantGenre:     "",
			wantYear:      0,
			wantTrack:     0,
		},
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

		gotArtist, err := audioFile.Artist()
		if err != nil {
			t.Fatalf("Artist(%s): %v", testCase.audioFilePath, err)
		}
		if gotArtist != testCase.wantArtist {
			t.Fatalf("Artist(%s): got %q, want %q", testCase.audioFilePath, gotArtist, testCase.wantArtist)
		}

		gotAlbum, err := audioFile.Album()
		if err != nil {
			t.Fatalf("Album(%s): %v", testCase.audioFilePath, err)
		}
		if gotAlbum != testCase.wantAlbum {
			t.Fatalf("Album(%s): got %q, want %q", testCase.audioFilePath, gotAlbum, testCase.wantAlbum)
		}

		gotComment, err := audioFile.Comment()
		if err != nil {
			t.Fatalf("Comment(%s): %v", testCase.audioFilePath, err)
		}
		if gotComment != testCase.wantComment {
			t.Fatalf("Comment(%s): got %q, want %q", testCase.audioFilePath, gotComment, testCase.wantComment)
		}

		gotGenre, err := audioFile.Genre()
		if err != nil {
			t.Fatalf("Genre(%s): %v", testCase.audioFilePath, err)
		}
		if gotGenre != testCase.wantGenre {
			t.Fatalf("Genre(%s): got %q, want %q", testCase.audioFilePath, gotGenre, testCase.wantGenre)
		}

		gotYear, err := audioFile.Year()
		if err != nil {
			t.Fatalf("Year(%s): %v", testCase.audioFilePath, err)
		}
		if gotYear != testCase.wantYear {
			t.Fatalf("Year(%s): got %d, want %d", testCase.audioFilePath, gotYear, testCase.wantYear)
		}

		gotTrack, err := audioFile.Track()
		if err != nil {
			t.Fatalf("Track(%s): %v", testCase.audioFilePath, err)
		}
		if gotTrack != testCase.wantTrack {
			t.Fatalf("Track(%s): got %d, want %d", testCase.audioFilePath, gotTrack, testCase.wantTrack)
		}
	}
}

func TestTagStringGettersClosedFile(t *testing.T) {
	audioFile, err := Open(completeBoaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	for _, testCase := range []struct {
		getter func() (string, error)
		name   string
	}{
		{audioFile.Title, "Title"},
		{audioFile.Artist, "Artist"},
		{audioFile.Album, "Album"},
		{audioFile.Comment, "Comment"},
		{audioFile.Genre, "Genre"},
	} {
		gotValue, err := testCase.getter()
		if !errors.Is(err, ErrFileClosed) {
			t.Fatalf("%s(closed): got %v, want ErrFileClosed", testCase.name, err)
		}
		if gotValue != "" {
			t.Fatalf("%s(closed): got %q, want empty string", testCase.name, gotValue)
		}
	}
}

func TestTagIntGettersClosedFile(t *testing.T) {
	audioFile, err := Open(completeBoaFilePath)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	audioFile.Close()

	for _, testCase := range []struct {
		getter func() (int, error)
		name   string
	}{
		{audioFile.Year, "Year"},
		{audioFile.Track, "Track"},
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

func TestTagGettersNilFile(t *testing.T) {
	var file *AudioFile

	for _, testCase := range []struct {
		getter func() (string, error)
		name   string
	}{
		{file.Title, "Title"},
		{file.Artist, "Artist"},
		{file.Album, "Album"},
		{file.Comment, "Comment"},
		{file.Genre, "Genre"},
	} {
		if _, err := testCase.getter(); !errors.Is(err, ErrFileClosed) {
			t.Fatalf("%s(nil): got %v, want ErrFileClosed", testCase.name, err)
		}
	}

	for _, testCase := range []struct {
		getter func() (int, error)
		name   string
	}{
		{file.Year, "Year"},
		{file.Track, "Track"},
	} {
		if _, err := testCase.getter(); !errors.Is(err, ErrFileClosed) {
			t.Fatalf("%s(nil): got %v, want ErrFileClosed", testCase.name, err)
		}
	}
}

package wiretag

// #include <taglib/tag_c.h>
import "C"

// Pending TODO:
// - taglib_tag_set_title
// - taglib_tag_set_artist
// - taglib_tag_set_album
// - taglib_tag_set_comment
// - taglib_tag_set_genre
// - taglib_tag_set_year
// - taglib_tag_set_track

func (file *AudioFile) Title() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cTitleString := C.taglib_tag_title(file.tag)
	if cTitleString == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cTitleString), nil
}

func (file *AudioFile) Artist() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cArtistString := C.taglib_tag_artist(file.tag)
	if cArtistString == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cArtistString), nil
}

func (file *AudioFile) Album() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cAlbumString := C.taglib_tag_album(file.tag)
	if cAlbumString == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cAlbumString), nil
}

func (file *AudioFile) Comment() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cCommentString := C.taglib_tag_comment(file.tag)
	if cCommentString == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cCommentString), nil
}

func (file *AudioFile) Genre() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cGenreString := C.taglib_tag_genre(file.tag)
	if cGenreString == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cGenreString), nil
}

func (file *AudioFile) Year() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	return int(C.taglib_tag_year(file.tag)), nil
}

func (file *AudioFile) Track() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	return int(C.taglib_tag_track(file.tag)), nil
}

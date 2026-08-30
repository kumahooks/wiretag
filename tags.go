package wiretag

// #include <taglib/tag_c.h>
import "C"

import "unsafe"

// Pending TODO:
// - taglib_tag_set_title
// - taglib_tag_set_artist
// - taglib_tag_set_album
// - taglib_tag_set_comment
// - taglib_tag_set_genre
// - taglib_tag_set_year
// - taglib_tag_set_track

// Title returns the track title of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Title() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	if file.tag == nil {
		return "", ErrInvalid
	}

	cTitleString := C.taglib_tag_title(file.tag)
	if cTitleString == nil {
		return "", nil
	}
	defer C.taglib_free(unsafe.Pointer(cTitleString))

	return C.GoString(cTitleString), nil
}

// Artist returns the artist name of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Artist() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	if file.tag == nil {
		return "", ErrInvalid
	}

	cArtistString := C.taglib_tag_artist(file.tag)
	if cArtistString == nil {
		return "", nil
	}
	defer C.taglib_free(unsafe.Pointer(cArtistString))

	return C.GoString(cArtistString), nil
}

// Album returns the album name of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Album() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	if file.tag == nil {
		return "", ErrInvalid
	}

	cAlbumString := C.taglib_tag_album(file.tag)
	if cAlbumString == nil {
		return "", nil
	}
	defer C.taglib_free(unsafe.Pointer(cAlbumString))

	return C.GoString(cAlbumString), nil
}

// Comment returns the comment of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Comment() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	if file.tag == nil {
		return "", ErrInvalid
	}

	cCommentString := C.taglib_tag_comment(file.tag)
	if cCommentString == nil {
		return "", nil
	}
	defer C.taglib_free(unsafe.Pointer(cCommentString))

	return C.GoString(cCommentString), nil
}

// Genre returns the genre of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Genre() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	if file.tag == nil {
		return "", ErrInvalid
	}

	cGenreString := C.taglib_tag_genre(file.tag)
	if cGenreString == nil {
		return "", nil
	}
	defer C.taglib_free(unsafe.Pointer(cGenreString))

	return C.GoString(cGenreString), nil
}

// Year returns the year of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Year() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	if file.tag == nil {
		return 0, ErrInvalid
	}

	return int(C.taglib_tag_year(file.tag)), nil
}

// Track returns the track number of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) Track() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	if file.tag == nil {
		return 0, ErrInvalid
	}

	return int(C.taglib_tag_track(file.tag)), nil
}

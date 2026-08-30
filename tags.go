package wiretag

// #include <stdlib.h>
// #include <taglib/tag_c.h>
import "C"

func (file *AudioFile) Title() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cTitle := C.taglib_tag_title(file.tag)
	if cTitle == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cTitle), nil
}

func (file *AudioFile) Artist() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cArtist := C.taglib_tag_artist(file.tag)
	if cArtist == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cArtist), nil
}

func (file *AudioFile) Album() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cAlbum := C.taglib_tag_album(file.tag)
	if cAlbum == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cAlbum), nil
}

func (file *AudioFile) Comment() (string, error) {
	if !file.isFileOpened() {
		return "", ErrFileClosed
	}

	cComment := C.taglib_tag_comment(file.tag)
	if cComment == nil {
		C.taglib_tag_free_strings()
		return "", nil
	}
	defer C.taglib_tag_free_strings()

	return C.GoString(cComment), nil
}

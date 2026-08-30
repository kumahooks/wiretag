// Package wiretag exposes Go bindings for popular taglib C++ library.
package wiretag

// #cgo pkg-config: taglib
// #cgo LDFLAGS: -ltag_c
// #include <stdlib.h>
// #include <taglib/tag_c.h>
import "C"

import (
	"errors"
	"unsafe"
)

var (
	ErrOpen    = errors.New("could not open file")
	ErrInvalid = errors.New("invalid audio file")
)

// AudioFile represents the taglib's TagLib_File, TagLib_Tag, and TagLib_AudioProperties of a given audio file.
type AudioFile struct {
	handle          *C.TagLib_File
	tag             *C.TagLib_Tag
	audioProperties *C.TagLib_AudioProperties
}

// Open receives a file path, and either return AudioFile, or an error.
//
// there are two ways we can fail here:
// - taglib_file_new returns null: in this case, taglib could not determine or open the file;
// - taglib_file_is_valid returns false: in this case, taglib either deems the file not opened, unreadable, or could not
// find valid informations for TagLib_Tag or TagLib_AudioProperties.
//
// in case taglib_file_is_valid is false, the byte stream has already been created by taglib with taglib_file_new, so we
// free it using taglib_file_free(file) before we return ErrInvalid.
func Open(filePath string) (*AudioFile, error) {
	cString := C.CString(filePath)
	defer C.free(unsafe.Pointer(cString))

	fileHandle := C.taglib_file_new(cString)
	if fileHandle == nil {
		return nil, ErrOpen
	}

	if C.taglib_file_is_valid(fileHandle) == 0 {
		C.taglib_file_free(fileHandle)
		return nil, ErrInvalid
	}

	return &AudioFile{
		handle:          fileHandle,
		tag:             C.taglib_file_tag(fileHandle),
		audioProperties: C.taglib_file_audioproperties(fileHandle),
	}, nil
}

// Close frees the memory allocated to the audio file's bytestream with taglib's taglib_file_free function.
func (file *AudioFile) Close() {
	if file == nil || file.handle == nil {
		return
	}

	C.taglib_file_free(file.handle)

	file.handle = nil
	file.tag = nil
	file.audioProperties = nil
}

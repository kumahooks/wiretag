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

type AudioFile struct {
	handle          *C.TagLib_File
	tag             *C.TagLib_Tag
	audioProperties *C.TagLib_AudioProperties
}

func Open(filepath string) (*AudioFile, error) {
	cString := C.CString(filepath)
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

func (file *AudioFile) Close() {
	if file == nil || file.handle == nil {
		return
	}

	C.taglib_file_free(file.handle)

	file.handle = nil
	file.tag = nil
	file.audioProperties = nil
}

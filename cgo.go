// Package wiretag exposes Go bindings for popular taglib C++ library.
package wiretag

// #cgo pkg-config: taglib
// #cgo LDFLAGS: -ltag_c
// #include <stdlib.h>
// #include <taglib/tag_c.h>
//
// static inline const char* wiretag_strarray_at(char** array, int index) {
//     return array[index];
// }
import "C"

import (
	"errors"
	"unsafe"
)

var (
	ErrOpen       = errors.New("could not open file")
	ErrInvalid    = errors.New("invalid audio file")
	ErrFileClosed = errors.New("audio file is closed")
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

// propertyValues fetches all values of a single property key, returning as an array of strings.
// taglib_property_get returns NULL if the values are empty, and in this case, we simply return an empty array.
func (file *AudioFile) propertyValues(propertyKey string) []string {
	cPropertyKey := C.CString(propertyKey)
	defer C.free(unsafe.Pointer(cPropertyKey))

	cPropertyValues := C.taglib_property_get(file.handle, cPropertyKey)
	if cPropertyValues == nil {
		return []string{}
	}
	defer C.taglib_property_free(cPropertyValues)

	return cStringSlice(cPropertyValues)
}

// Properties returns every property of the file keyed by property name, and string arrays for their values.
// If the file is closed (either file or file.handle is nil), we return an error indicating this is undefined behavior.
//
// If the file has no property whatsoever (taglib returns NULL at taglib_property_keys), we return an empty map. Same behavior
// for a property with no values: we store an empty array at that key.
func (file *AudioFile) Properties() (map[string][]string, error) {
	if !file.isFileOpened() {
		return nil, ErrFileClosed
	}

	cKeys := C.taglib_property_keys(file.handle)
	if cKeys == nil {
		return map[string][]string{}, nil
	}
	defer C.taglib_property_free(cKeys)

	propertyMap := make(map[string][]string)
	for _, property := range cStringSlice(cKeys) {
		propertyMap[property] = file.propertyValues(property)
	}

	return propertyMap, nil
}

func (file *AudioFile) isFileOpened() bool {
	return file != nil && file.handle != nil
}

// cStringSlice copies a NULL-terminated char** into a Go string slice.
func cStringSlice(arr **C.char) []string {
	values := []string{}
	for i := 0; ; i++ {
		value := C.wiretag_strarray_at(arr, C.int(i))
		if value == nil {
			break
		}

		values = append(values, C.GoString(value))
	}

	return values
}

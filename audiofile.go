package wiretag

// #include <stdlib.h>
// #include <taglib/tag_c.h>
import "C"

import "unsafe"

// Pending TODO:
// - taglib_memory_iostream_new
// - taglib_iostream_free
// - TagLib_File_Type
// - taglib_file_new_type
// - taglib_file_new_iostream
// - taglib_file_save

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

	if C.taglib_file_is_valid(fileHandle) != 1 {
		freeAudioFile(fileHandle)
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

	freeAudioFile(file.handle)

	file.handle = nil
	file.tag = nil
	file.audioProperties = nil
}

// IsValid reports whether the audio file is open, readable, and valid information for the Tag and/or AudioProperties
// was found. It is the direct equivalent of taglib_file_is_valid.
func (file *AudioFile) IsValid() bool {
	return file.isFileOpened() && C.taglib_file_is_valid(file.handle) == 1
}

func (file *AudioFile) isFileOpened() bool {
	return file != nil && file.handle != nil
}

func freeAudioFile(fileHandle *C.TagLib_File) {
	C.taglib_file_free(fileHandle)
}

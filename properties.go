package wiretag

// #include <stdlib.h>
// #include <taglib/tag_c.h>
import "C"

import "unsafe"

// Pending TODO:
// - taglib_property_set
// - taglib_property_set_append

// Properties returns every property of the file keyed by property name, and slice of strings for their values.
// If the file is closed (either file or file.handle is nil), we return an error indicating this is undefined behavior.
//
// If the file has no property whatsoever (taglib returns NULL at taglib_property_keys), we return an empty map. Same behavior
// for a property with no values: we store an empty slice at that key.
func (file *AudioFile) Properties() (map[string][]string, error) {
	if !file.isFileOpened() {
		return nil, ErrFileClosed
	}

	propertyMap := make(map[string][]string)
	for _, property := range file.propertyKeys() {
		propertyMap[property] = file.propertyValues(property)
	}

	return propertyMap, nil
}

// PropertyKeys returns every property key present in the file.
// If the file is closed (either file or file.handle is nil), we return an error indicating this is undefined behavior.
//
// If the file has no property whatsoever (taglib returns NULL at taglib_property_keys), we return an empty slice.
func (file *AudioFile) PropertyKeys() ([]string, error) {
	if !file.isFileOpened() {
		return nil, ErrFileClosed
	}

	return file.propertyKeys(), nil
}

// PropertyValues returns all values of a single property key as a slice of strings.
// If the file is closed (either file or file.handle is nil), we return an error indicating this is undefined behavior.
//
// The key is case-insensitive. If the key is absent or holds no values, we return an empty slice.
func (file *AudioFile) PropertyValues(propertyKey string) ([]string, error) {
	if !file.isFileOpened() {
		return nil, ErrFileClosed
	}

	return file.propertyValues(propertyKey), nil
}

// propertyKeys returns every property key present in the file as a slice of strings.
// taglib_property_keys returns NULL if the file has no properties, and in this case, we return an empty slice.
func (file *AudioFile) propertyKeys() []string {
	cKeys := C.taglib_property_keys(file.handle)
	if cKeys == nil {
		return []string{}
	}
	defer C.taglib_property_free(cKeys)

	return stringSliceFromCStringArray(cKeys)
}

// propertyValues fetches all values of a single property key, returning as a slice of strings.
// taglib_property_get returns NULL if the values are empty, and in this case, we simply return an empty slice.
func (file *AudioFile) propertyValues(propertyKey string) []string {
	cPropertyKey := C.CString(propertyKey)
	defer C.free(unsafe.Pointer(cPropertyKey))

	cPropertyValues := C.taglib_property_get(file.handle, cPropertyKey)
	if cPropertyValues == nil {
		return []string{}
	}
	defer C.taglib_property_free(cPropertyValues)

	return stringSliceFromCStringArray(cPropertyValues)
}

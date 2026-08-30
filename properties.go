package wiretag

// #include <stdlib.h>
// #include <taglib/tag_c.h>
import "C"

import "unsafe"

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

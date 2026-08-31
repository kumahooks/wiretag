// shims.c defines C helper functions for some specific taglib functionalities.
#include <taglib/tag_c.h>

// Implemented in shims.cpp: swaps taglib's default debug listener (stderr) for a silent one.
extern void wiretag_silence_warnings(void);

const char* wiretag_strarray_at(char** array, int index)
{
	return array[index];
}

void wiretag_init(void)
{
	taglib_set_string_management_enabled(0);
	wiretag_silence_warnings();
}

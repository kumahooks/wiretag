package wiretag

// #include <taglib/tag_c.h>
import "C"

func (file *AudioFile) AudioLength() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	return int(C.taglib_audioproperties_length(file.audioProperties)), nil
}

func (file *AudioFile) AudioBitRate() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	return int(C.taglib_audioproperties_bitrate(file.audioProperties)), nil
}

func (file *AudioFile) AudioSampleRate() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	return int(C.taglib_audioproperties_samplerate(file.audioProperties)), nil
}

func (file *AudioFile) AudioChannels() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	return int(C.taglib_audioproperties_channels(file.audioProperties)), nil
}

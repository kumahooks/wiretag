package wiretag

// #include <taglib/tag_c.h>
import "C"

// AudioLength returns the duration of the file in seconds.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) AudioLength() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	if file.audioProperties == nil {
		return 0, ErrInvalid
	}

	return int(C.taglib_audioproperties_length(file.audioProperties)), nil
}

// AudioBitrate returns the bitrate of the file in kbps.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) AudioBitrate() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	if file.audioProperties == nil {
		return 0, ErrInvalid
	}

	return int(C.taglib_audioproperties_bitrate(file.audioProperties)), nil
}

// AudioSampleRate returns the sample rate of the file in Hz.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) AudioSampleRate() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	if file.audioProperties == nil {
		return 0, ErrInvalid
	}

	return int(C.taglib_audioproperties_samplerate(file.audioProperties)), nil
}

// AudioChannels returns the number of audio channels of the file.
// If the file is closed (either file or file.handle is nil), we return ErrFileClosed.
func (file *AudioFile) AudioChannels() (int, error) {
	if !file.isFileOpened() {
		return 0, ErrFileClosed
	}

	if file.audioProperties == nil {
		return 0, ErrInvalid
	}

	return int(C.taglib_audioproperties_channels(file.audioProperties)), nil
}

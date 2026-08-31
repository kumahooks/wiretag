// shims.cpp implements the C++-only portion of the shims
#include <taglib/tag_c.h>
#include <taglib/tdebuglistener.h>

extern "C" void wiretag_silence_warnings(void);

namespace {
class SilentListener : public TagLib::DebugListener {
  public:
	void printMessage(const TagLib::String&) override
	{
	}
};
} // namespace

void wiretag_silence_warnings(void)
{
	static SilentListener listener;
	TagLib::setDebugListener(&listener);
}


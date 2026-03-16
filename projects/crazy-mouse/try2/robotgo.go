package robotgo

/*
#include <CoreGraphics/CoreGraphics.h>
*/
import "C"

// GetMainId get the main display id
func GetMainId() int {
	return int(C.CGMainDisplayID())
}

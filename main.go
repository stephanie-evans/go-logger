package go_logger

/*
#cgo LDFLAGS: -L./lib/target/release -lrlogger
#include <stdlib.h>
#include <stdint.h>

extern void init(const char* message);
*/
import "C"

func initialize(prefix string) {
	message := C.CString(prefix)
	C.init(message)
}

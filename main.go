package go_logger

/*
#cgo LDFLAGS: -L./lib/target/release -lrlogger
#include <stdlib.h>
#include <stdint.h>

extern void init(const char* message);
*/
import "C"

func initialize(genesisNodeConfig string) {
	result := C.CString(genesisNodeConfig)
	C.init(result)
}

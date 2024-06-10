package main

// #include <uuid/uuid.h>
// #include <stdlib.h>
//
// // create a uuid function in C to return a uuid char*
// char* _go_uuid() {
//   uuid_t uuid;
//   uuid_generate_random(uuid);
//   char *str = malloc(37);
//   uuid_unparse_lower(uuid, str);
//   return str;
// }
import "C"
import (
	"fmt"
	"time"
)

// uuid generates a UUID using the C shared library.
// It returns a Go string.
func uuid() string {
	return C.GoString(C._go_uuid())
}

func memoryLeak() {
	for {
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		_ = uuid

		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	// and now it's simple to use
	myuuid := uuid() // this is a go string now
	fmt.Println(myuuid)

	memoryLeak()
}

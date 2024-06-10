package main

// #include <uuid/uuid.h>
import "C"
import "fmt"

func main() {
	{
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		var uuid_str *C.char
		uuid_str = (*C.char)(C.malloc(37))
		C.uuid_generate_random(uuid)
		C.uuid_unparse(uuid, uuid_str)
		fmt.Println(C.GoString(uuid_str))
	}

	{
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		var uuid_str *C.char
		uuid_str = (*C.char)(C.malloc(37))
		C.uuid_generate_random(uuid)
		C.uuid_unparse(uuid, uuid_str)
		fmt.Println(C.GoString(uuid_str))
	}
}

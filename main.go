package main

// #include <uuid/uuid.h>
import "C"
import "fmt"

//// In C:
//// we need a uuid_t variable to initalize
//uuid uuid_t;
//// then we generate the random string. I use the random form
//// but you can use other generate methods.
//uuid_generate_random(uuid);
//// To get a uuid string, we need to "unparse"
//char uuid_str[37];
//uuid_unparse_lower(uuid, uuid_str);
//// In the uuid_str char*, there is a uuid

func main() {
	var uuid *C.uchar
	uuid = (*C.uchar)(C.malloc(16))
	var uuid_str *C.char
	uuid_str = (*C.char)(C.malloc(37))
	C.uuid_generate_random(uuid)
	C.uuid_unparse(uuid, uuid_str)
	fmt.Println(C.GoString(uuid_str))
}

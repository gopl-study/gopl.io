package main

// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <stdlib.h>
// #include <string.h>
// #include <mysql.h>
import "C"
import (
	"io/ioutil"
	"net/http"
	"unicode/utf8"
	"unsafe"
)

const arrLength = 1 << 30

//export http_get_init
func http_get_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.char {
	if args.arg_count != 1 {
		C.strcpy(message, C.CString("only one param"))
		return 1
	}

	return 0
}

//export http_get
func http_get(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.int,
	null_value *C.char, message *C.char) *C.char {
	gArg_count := uint(args.arg_count)

	var ret string
	gArgs := ((*[arrLength]*C.char)(unsafe.Pointer(args.args)))[:gArg_count:gArg_count]
	resp, err := http.Get(C.GoString(gArgs[0]))
	if err != nil {
		ret = err.Error()
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		ret = string(data)
	}

	result = C.CString(ret)
	*length = C.int(utf8.RuneCountInString(ret))
	return result
}


func main() {
}
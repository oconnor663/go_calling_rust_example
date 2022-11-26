package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

// #cgo LDFLAGS: -Lrust_wrapper/target/release -lrust_wrapper
// #include <stddef.h>
// void blake3_wrapper(const char* input, size_t input_len, char* output);
import "C"

func main() {
	input := []byte("foo")
	output := make([]byte, 32)
	C.blake3_wrapper(
		(*C.char)(unsafe.Pointer(&input[0])),
		C.size_t(len(input)),
		(*C.char)(unsafe.Pointer(&output[0])),
	)
	fmt.Printf("BLAKE3(\"foo\") = %s\n", hex.EncodeToString(output))
}

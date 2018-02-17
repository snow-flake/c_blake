
package main

/*
#cgo CFLAGS: -I./DyLib
#cgo LDFLAGS: -L. -lc_blake
#include "./include/library.h"
*/
import "C"
import "unsafe"

func main() {
    var output [64]byte
    input := []byte("The quick brown fox jumps over the lazy dog")
    C.blake256_hash(
		unsafe.Pointer(&output),
		unsafe.Pointer(&input),
		uint64(len(input)),
	)
}

// func hash(input string) string {
// 	input_buffer := [64]byte{}
// 	//
// }

// func bin2hex(input [32]byte)
// inline static std::string bin2hex([32]byte) {
// 	std::vector<char> buffer;
// 	buffer.reserve(len);
// 	for (unsigned int i = 0; i < len / 2; i++) {
// 		buffer[i * 2] = hf_bin2hex((in[i] & 0xF0) >> 4);
// 		buffer[i * 2 + 1] = hf_bin2hex(in[i] & 0x0F);
// 	}
// 	auto output = std::string(buffer.data(), len);
// 	return output;
// }

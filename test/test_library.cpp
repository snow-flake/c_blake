// Let Catch provide main():
#define CATCH_CONFIG_MAIN

#include "catch.hpp"
#include "../src/library.c"
#include <array>
#include <string>
#include <iostream>
#include <iomanip>


int Factorial( int number ) {
	// return number <= 1 ? number : Factorial( number - 1 ) * number;  // fail
	return number <= 1 ? 1      : Factorial( number - 1 ) * number;  // pass
}


inline static char hf_bin2hex(unsigned char c) {
	if (c <= 0x9)
		return '0' + c;
	else
		return 'a' - 0xA + c;
}

inline static std::string bin2hex(const unsigned char *in, unsigned int len) {
	std::vector<char> buffer;
	buffer.reserve(len);
	for (unsigned int i = 0; i < len / 2; i++) {
		buffer[i * 2] = hf_bin2hex((in[i] & 0xF0) >> 4);
		buffer[i * 2 + 1] = hf_bin2hex(in[i] & 0x0F);
	}
	auto output = std::string(buffer.data(), len);
	return output;
}

const std::string hash_string(const std::string input_str) {
	assert(sizeof(unsigned char) == sizeof(uint8_t));

	std::array<uint8_t, 65> input;
	input.fill(0);
	strcpy((char *) &input.data()[0], input_str.c_str());

	std::array<uint8_t, 64> output;
	output.fill(0);

	blake256_hash(output.data(), input.data(), input_str.length());

	auto output_str = bin2hex(output.data(), 64);
	return output_str;
}


TEST_CASE( "Factorial of 0 is 1 (fail)", "[single-file]" ) {
    REQUIRE( Factorial(0) == 1 );
}

TEST_CASE( "Factorials of 1 and higher are computed (pass)", "[single-file]" ) {
    REQUIRE( Factorial(1) == 1 );
    REQUIRE( Factorial(2) == 2 );
    REQUIRE( Factorial(3) == 6 );
    REQUIRE( Factorial(10) == 3628800 );
}

TEST_CASE( "Blake hash should match", "[single-file]" ) {
    REQUIRE( hash_string("The quick brown fox jumps over the lazy dog") == "7576698ee9cad30173080678e5965916adbb11cb5245d386bf1ffda1cb26c9d7" );
}


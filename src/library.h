#ifndef C_BLAKE_LIBRARY_H
#define C_BLAKE_LIBRARY_H

#include <stdint.h>

void blake256_hash(uint8_t * output, const uint8_t * input, const uint64_t input_len);

#endif
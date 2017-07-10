/*
  This file is part of elhash.

  elhash is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  elhash is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with elhash.  If not, see <http://www.gnu.org/licenses/>.
*/

/** @file sha3.cpp
* 
* @date 2016
*/
#include <stdint.h>
#include <cryptopp/sha3.h>

extern "C" {
struct elhash_h256;
typedef struct elhash_h256 elhash_h256_t;
void SHA3_256(elhash_h256_t const* ret, uint8_t const* data, size_t size)
{
	CryptoPP::SHA3_256().CalculateDigest((uint8_t*)ret, data, size);
}

void SHA3_512(uint8_t* const ret, uint8_t const* data, size_t size)
{
	CryptoPP::SHA3_512().CalculateDigest(ret, data, size);
}
}

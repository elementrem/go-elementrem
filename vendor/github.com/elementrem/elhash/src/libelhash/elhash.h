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

/** @file elhash.h

*/
#pragma once

#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include <stddef.h>
#include "compiler.h"

#define ELHASH_REVISION 23
#define ELHASH_DATASET_BYTES_INIT 1073741824U // 2**30
#define ELHASH_DATASET_BYTES_GROWTH 8388608U  // 2**23
#define ELHASH_CACHE_BYTES_INIT 1073741824U // 2**24
#define ELHASH_CACHE_BYTES_GROWTH 131072U  // 2**17
#define ELHASH_EPOCH_LENGTH 30000U
#define ELHASH_MIX_BYTES 128
#define ELHASH_HASH_BYTES 64
#define ELHASH_DATASET_PARENTS 256
#define ELHASH_CACHE_ROUNDS 3
#define ELHASH_ACCESSES 64
#define ELHASH_DAG_MAGIC_NUM_SIZE 8
#define ELHASH_DAG_MAGIC_NUM 0xFEE1DEADBADDCAFE

#ifdef __cplusplus
extern "C" {
#endif

/// Type of a seedhash/blockhash e.t.c.
typedef struct elhash_h256 { uint8_t b[32]; } elhash_h256_t;

// convenience macro to statically initialize an h256_t
// usage:
// elhash_h256_t a = elhash_h256_static_init(1, 2, 3, ... )
// have to provide all 32 values. If you don't provide all the rest
// will simply be unitialized (not guranteed to be 0)
#define elhash_h256_static_init(...)			\
	{ {__VA_ARGS__} }

struct elhash_light;
typedef struct elhash_light* elhash_light_t;
struct elhash_full;
typedef struct elhash_full* elhash_full_t;
typedef int(*elhash_callback_t)(unsigned);

typedef struct elhash_return_value {
	elhash_h256_t result;
	elhash_h256_t mix_hash;
	bool success;
} elhash_return_value_t;

/**
 * Allocate and initialize a new elhash_light handler
 *
 * @param block_number   The block number for which to create the handler
 * @return               Newly allocated elhash_light handler or NULL in case of
 *                       ERRNOMEM or invalid parameters used for @ref elhash_compute_cache_nodes()
 */
elhash_light_t elhash_light_new(uint64_t block_number);
/**
 * Frees a previously allocated elhash_light handler
 * @param light        The light handler to free
 */
void elhash_light_delete(elhash_light_t light);
/**
 * Calculate the light client data
 *
 * @param light          The light client handler
 * @param header_hash    The header hash to pack into the mix
 * @param nonce          The nonce to pack into the mix
 * @return               an object of elhash_return_value_t holding the return values
 */
elhash_return_value_t elhash_light_compute(
	elhash_light_t light,
	elhash_h256_t const header_hash,
	uint64_t nonce
);

/**
 * Allocate and initialize a new elhash_full handler
 *
 * @param light         The light handler containing the cache.
 * @param callback      A callback function with signature of @ref elhash_callback_t
 *                      It accepts an unsigned with which a progress of DAG calculation
 *                      can be displayed. If all goes well the callback should return 0.
 *                      If a non-zero value is returned then DAG generation will stop.
 *                      Be advised. A progress value of 100 means that DAG creation is
 *                      almost complete and that this function will soon return succesfully.
 *                      It does not mean that the function has already had a succesfull return.
 * @return              Newly allocated elhash_full handler or NULL in case of
 *                      ERRNOMEM or invalid parameters used for @ref elhash_compute_full_data()
 */
elhash_full_t elhash_full_new(elhash_light_t light, elhash_callback_t callback);

/**
 * Frees a previously allocated elhash_full handler
 * @param full    The light handler to free
 */
void elhash_full_delete(elhash_full_t full);
/**
 * Calculate the full client data
 *
 * @param full           The full client handler
 * @param header_hash    The header hash to pack into the mix
 * @param nonce          The nonce to pack into the mix
 * @return               An object of elhash_return_value to hold the return value
 */
elhash_return_value_t elhash_full_compute(
	elhash_full_t full,
	elhash_h256_t const header_hash,
	uint64_t nonce
);
/**
 * Get a pointer to the full DAG data
 */
void const* elhash_full_dag(elhash_full_t full);
/**
 * Get the size of the DAG data
 */
uint64_t elhash_full_dag_size(elhash_full_t full);

/**
 * Calculate the seedhash for a given block number
 */
elhash_h256_t elhash_get_seedhash(uint64_t block_number);

#ifdef __cplusplus
}
#endif

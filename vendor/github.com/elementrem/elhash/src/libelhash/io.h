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
/** @file io.h


 */
#pragma once
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>
#ifdef __cplusplus
#define __STDC_FORMAT_MACROS 1
#endif
#include <inttypes.h>
#include "endian.h"
#include "elhash.h"

#ifdef __cplusplus
extern "C" {
#endif
// Maximum size for mutable part of DAG file name
// 6 is for "full-R", the suffix of the filename
// 10 is for maximum number of digits of a uint32_t (for REVISION)
// 1 is for - and 16 is for the first 16 hex digits for first 8 bytes of
// the seedhash and last 1 is for the null terminating character
// Reference: https://github.com/elementrem
#define DAG_MUTABLE_NAME_MAX_SIZE (6 + 10 + 1 + 16 + 1)
/// Possible return values of @see elhash_io_prepare
enum elhash_io_rc {
	ELHASH_IO_FAIL = 0,           ///< There has been an IO failure
	ELHASH_IO_MEMO_SIZE_MISMATCH, ///< DAG with revision/hash match, but file size was wrong.
	ELHASH_IO_MEMO_MISMATCH,      ///< The DAG file did not exist or there was revision/hash mismatch
	ELHASH_IO_MEMO_MATCH,         ///< DAG file existed and revision/hash matched. No need to do anything
};

// small hack for windows. I don't feel I should use va_args and forward just
// to have this one function properly cross-platform abstracted
#if defined(_WIN32) && !defined(__GNUC__)
#define snprintf(...) sprintf_s(__VA_ARGS__)
#endif

/**
 * Logs a critical error in important parts of elhash. Should mostly help
 * figure out what kind of problem (I/O, memory e.t.c.) causes a NULL
 * elhash_full_t
 */
#ifdef ELHASH_PRINT_CRITICAL_OUTPUT
#define ELHASH_CRITICAL(...)							\
	do													\
	{													\
		printf("ELHASH CRITICAL ERROR: "__VA_ARGS__);	\
		printf("\n");									\
		fflush(stdout);									\
	} while (0)
#else
#define ELHASH_CRITICAL(...)          
#endif

/**
 * Prepares io for elhash
 *
 * Create the DAG directory and the DAG file if they don't exist.
 *
 * @param[in] dirname        A null terminated c-string of the path of the elhash
 *                           data directory. If it does not exist it's created.
 * @param[in] seedhash       The seedhash of the current block number, used in the
 *                           naming of the file as can be seen from the spec at:
 *                           https://github.com/elementrem
 * @param[out] output_file   If there was no failure then this will point to an open
 *                           file descriptor. User is responsible for closing it.
 *                           In the case of memo match then the file is open on read
 *                           mode, while on the case of mismatch a new file is created
 *                           on write mode
 * @param[in] file_size      The size that the DAG file should have on disk
 * @param[out] force_create  If true then there is no check to see if the file
 *                           already exists
 * @return                   For possible return values @see enum elhash_io_rc
 */
enum elhash_io_rc elhash_io_prepare(
	char const* dirname,
	elhash_h256_t const seedhash,
	FILE** output_file,
	uint64_t file_size,
	bool force_create
);

/**
 * An fopen wrapper for no-warnings crossplatform fopen.
 *
 * Msvc compiler considers fopen to be insecure and suggests to use their
 * alternative. This is a wrapper for this alternative. Another way is to
 * #define _CRT_SECURE_NO_WARNINGS, but disabling all security warnings does
 * not sound like a good idea.
 *
 * @param file_name        The path to the file to open
 * @param mode             Opening mode. Check fopen()
 * @return                 The FILE* or NULL in failure
 */
FILE* elhash_fopen(char const* file_name, char const* mode);

/**
 * An strncat wrapper for no-warnings crossplatform strncat.
 *
 * Msvc compiler considers strncat to be insecure and suggests to use their
 * alternative. This is a wrapper for this alternative. Another way is to
 * #define _CRT_SECURE_NO_WARNINGS, but disabling all security warnings does
 * not sound like a good idea.
 *
 * @param des              Destination buffer
 * @param dest_size        Maximum size of the destination buffer. This is the
 *                         extra argument for the MSVC secure strncat
 * @param src              Souce buffer
 * @param count            Number of bytes to copy from source
 * @return                 If all is well returns the dest buffer. If there is an
 *                         error returns NULL
 */
char* elhash_strncat(char* dest, size_t dest_size, char const* src, size_t count);

/**
 * A cross-platform mkdir wrapper to create a directory or assert it's there
 *
 * @param dirname        The full path of the directory to create
 * @return               true if the directory was created or if it already
 *                       existed
 */
bool elhash_mkdir(char const* dirname);

/**
 * Get a file's size
 *
 * @param[in] f        The open file stream whose size to get
 * @param[out] size    Pass a size_t by reference to contain the file size
 * @return             true in success and false if there was a failure
 */
bool elhash_file_size(FILE* f, size_t* ret_size);

/**
 * Get a file descriptor number from a FILE stream
 *
 * @param f            The file stream whose fd to get
 * @return             Platform specific fd handler
 */
int elhash_fileno(FILE* f);

/**
 * Create the filename for the DAG.
 *
 * @param dirname            The directory name in which the DAG file should reside
 *                           If it does not end with a directory separator it is appended.
 * @param filename           The actual name of the file
 * @param filename_length    The length of the filename in bytes
 * @return                   A char* containing the full name. User must deallocate.
 */
char* elhash_io_create_filename(
	char const* dirname,
	char const* filename,
	size_t filename_length
);

/**
 * Gets the default directory name for the DAG depending on the system
 *
 * The spec defining this directory is here: https://github.com/elementrem
 *
 * @param[out] strbuf          A string buffer of sufficient size to keep the
 *                             null termninated string of the directory name
 * @param[in]  buffsize        Size of @a strbuf in bytes
 * @return                     true for success and false otherwise
 */
bool elhash_get_default_dirname(char* strbuf, size_t buffsize);

static inline bool elhash_io_mutable_name(
	uint32_t revision,
	elhash_h256_t const* seed_hash,
	char* output
)
{
    uint64_t hash = *((uint64_t*)seed_hash);
#if LITTLE_ENDIAN == BYTE_ORDER
    hash = elhash_swap_u64(hash);
#endif
    return snprintf(output, DAG_MUTABLE_NAME_MAX_SIZE, "full-R%u-%016" PRIx64, revision, hash) >= 0;
}

#ifdef __cplusplus
}
#endif

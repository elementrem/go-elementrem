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
/*
 */

#include "io.h"
#include <direct.h>
#include <errno.h>
#include <stdio.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <shlobj.h>

FILE* elhash_fopen(char const* file_name, char const* mode)
{
	FILE* f;
	return fopen_s(&f, file_name, mode) == 0 ? f : NULL;
}

char* elhash_strncat(char* dest, size_t dest_size, char const* src, size_t count)
{
	return strncat_s(dest, dest_size, src, count) == 0 ? dest : NULL;
}

bool elhash_mkdir(char const* dirname)
{
	int rc = _mkdir(dirname);
	return rc != -1 || errno == EEXIST;
}

int elhash_fileno(FILE* f)
{
	return _fileno(f);
}

char* elhash_io_create_filename(
	char const* dirname,
	char const* filename,
	size_t filename_length
)
{
	size_t dirlen = strlen(dirname);
	size_t dest_size = dirlen + filename_length + 1;
	if (dirname[dirlen] != '\\' || dirname[dirlen] != '/') {
		dest_size += 1;
	}
	char* name = malloc(dest_size);
	if (!name) {
		return NULL;
	}

	name[0] = '\0';
	elhash_strncat(name, dest_size, dirname, dirlen);
	if (dirname[dirlen] != '\\' || dirname[dirlen] != '/') {
		elhash_strncat(name, dest_size, "\\", 1);
	}
	elhash_strncat(name, dest_size, filename, filename_length);
	return name;
}

bool elhash_file_size(FILE* f, size_t* ret_size)
{
	struct _stat st;
	int fd;
	if ((fd = _fileno(f)) == -1 || _fstat(fd, &st) != 0) {
		return false;
	}
	*ret_size = st.st_size;
	return true;
}

bool elhash_get_default_dirname(char* strbuf, size_t buffsize)
{
	static const char dir_suffix[] = "Elhash\\";
	strbuf[0] = '\0';
	if (!SUCCEEDED(SHGetFolderPathA(NULL, CSIDL_LOCAL_APPDATA, NULL, 0, (CHAR*)strbuf))) {
		return false;
	}
	if (!elhash_strncat(strbuf, buffsize, "\\", 1)) {
		return false;
	}

	return elhash_strncat(strbuf, buffsize, dir_suffix, sizeof(dir_suffix));
}

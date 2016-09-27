#!/usr/bin/env python
import os
from distutils.core import setup, Extension
sources = [
    'src/python/core.c',
    'src/libelhash/io.c',
    'src/libelhash/internal.c',
    'src/libelhash/sha3.c']
if os.name == 'nt':
    sources += [
        'src/libelhash/util_win32.c',
        'src/libelhash/io_win32.c',
        'src/libelhash/mmap_win32.c',
    ]
else:
    sources += [
        'src/libelhash/io_posix.c'
    ]
depends = [
    'src/libelhash/elhash.h',
    'src/libelhash/compiler.h',
    'src/libelhash/data_sizes.h',
    'src/libelhash/endian.h',
    'src/libelhash/elhash.h',
    'src/libelhash/io.h',
    'src/libelhash/fnv.h',
    'src/libelhash/internal.h',
    'src/libelhash/sha3.h',
    'src/libelhash/util.h',
]
pyelhash = Extension('pyelhash',
                     sources=sources,
                     depends=depends,
                     extra_compile_args=["-Isrc/", "-std=gnu99", "-Wall"])

setup(
    name='pyelhash',
    author="Matthew Wampler-Doty",
    author_email="matthew.wampler.doty@gmail.com",
    license='GPL',
    version='0.1.23',
    url='https://github.com/elementrem/',
    download_url='https://github.com/elementrem/',
    description=('Python wrappers for elhash, the elementrem proof of work'
                 'hashing function'),
    ext_modules=[pyelhash],
)

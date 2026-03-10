// Copyright ©2012 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bgzf implements BGZF format reading and writing according to the
// SAM specification.
//
// The specification is available at https://github.com/samtools/hts-specs.
package bgzf

import (
	"errors"
	"time"
)

const (
	BlockSize    = 0x0ff00 // The maximum size of an uncompressed input data block.
	MaxBlockSize = 0x10000 // The maximum size of a compressed output block.
)

const (
	bgzfExtra = "BC\x02\x00\x00\x00"
	minFrame  = 20 + len(bgzfExtra) // Minimum bgzf header+footer length.
)

var (
	bgzfExtraPrefix = []byte(bgzfExtra[:4])
	unixEpoch       = time.Unix(0, 0)
)

func compressBound(srcLen int) int {
	return srcLen + srcLen>>12 + srcLen>>14 + srcLen>>25 + 13 + minFrame
}

func init() {
	if compressBound(BlockSize) > MaxBlockSize {
		panic("bam: BlockSize too large")
	}
}

var (
	ErrCorrupt           = errors.New("bgzf: corrupt block")
	ErrNotASeeker        = errors.New("bgzf: not a seeker")
	ErrContaminatedCache = errors.New("bgzf: cache owner mismatch")
	ErrNoBlockSize       = errors.New("bgzf: could not determine block size")
)

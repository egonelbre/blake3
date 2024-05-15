//go:build !purego && arm64
// +build !purego,arm64

package compress

import (
	"github.com/zeebo/blake3/internal/alg/compress/compress_arm64"
)

func Compress(chain *[8]uint32, block *[16]uint32, counter uint64, blen uint32, flags uint32, out *[16]uint32) {
	compress_arm64.Compress(chain, block, counter, blen, flags, out)
}

package compress_arm64

import "github.com/zeebo/blake3/internal/consts"

func Compress(
	chain *[8]uint32,
	block *[16]uint32,
	counter uint64,
	blen uint32,
	flags uint32,
	out *[16]uint32,
) {

	*out = [16]uint32{
		chain[0], chain[1], chain[2], chain[3],
		chain[4], chain[5], chain[6], chain[7],
		consts.IV0, consts.IV1, consts.IV2, consts.IV3,
		uint32(counter), uint32(counter >> 32), blen, flags,
	}

	rcompress(out, block)
}

func rcompress(state, block *[16]uint32)

//go:noinline
func rcompress2(state, block *[16]uint32) {
	state[0] = 100
}

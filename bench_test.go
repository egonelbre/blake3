package blake3

import (
	"fmt"
	"testing"
)

func BenchmarkIncremental(b *testing.B) {
	run := func(b *testing.B, size int) {
		h := new(hasher)
		out := make([]byte, 32)
		buf := make([]byte, size)

		b.ReportAllocs()
		b.SetBytes(int64(len(buf)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			h.update(buf)
			h.finalize(out)
			h.reset()
		}
	}

	for _, n := range []int{
		1, 4, 8, 12, 16,
	} {
		b.Run(fmt.Sprintf("%04d_block", n), func(b *testing.B) { run(b, n*64) })
	}

	for _, n := range []int{
		1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024,
	} {
		b.Run(fmt.Sprintf("%04d_kib", n), func(b *testing.B) { run(b, n*1024) })
	}
	for _, n := range []int{
		1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024,
	} {
		b.Run(fmt.Sprintf("%04d_kib+512", n), func(b *testing.B) { run(b, n*1024+512) })
	}
}

func BenchmarkHashF_1(b *testing.B) {
	var input [8192]byte
	var chain [8]uint32
	var out chainVector

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hashF(&input, 1, 0, 0, &out, &chain)
	}
}

func BenchmarkHashF_1536(b *testing.B) {
	var input [8192]byte
	var chain [8]uint32
	var out chainVector

	b.SetBytes(1536)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hashF(&input, 1536, 0, 0, &out, &chain)
	}
}

func BenchmarkHashF_8K(b *testing.B) {
	var input [8192]byte
	var chain [8]uint32
	var out chainVector

	b.SetBytes(8192)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hashF(&input, 8192, 0, 0, &out, &chain)
	}
}

func BenchmarkHashP(b *testing.B) {
	var left chainVector
	var right chainVector
	var out chainVector

	for n := 1; n <= 8; n++ {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.SetBytes(int64(64 * n))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				hashP(&left, &right, 0, &out, n)
			}
		})
	}
}

func BenchmarkCompress(b *testing.B) {
	var c [8]uint32
	var m, o [16]uint32

	b.SetBytes(64)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		compress(&c, &m, 0, 0, 0, &o)
	}
}

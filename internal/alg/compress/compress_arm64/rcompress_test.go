package compress_arm64

import (
	"math/bits"
	"testing"

	"github.com/zeebo/blake3/internal/alg/compress/compress_pure"
)

func TestDebug(t *testing.T) {
	var a, b [16]uint32
	a[0] = 0x00112233
	a[1] = 0x44556677
	a[2] = 0x8899aabb
	a[3] = 0xccddeeff
	a[4] = 0x00112233
	a[5] = 0x44556677
	a[6] = 0x8899aabb
	a[7] = 0xccddeeff

	rcompress(&a, &b)

	t.Logf("%08x %08x %08x %08x", a[0], a[1], a[2], a[3])
}

func BenchmarkRCompress(b *testing.B) {
	var state, block [16]uint32
	for i := 0; i < b.N; i++ {
		rcompress(&state, &block)
	}
}

func BenchmarkRCompressGo(b *testing.B) {
	var state, block [16]uint32
	for i := 0; i < b.N; i++ {
		compress_pure.RCompress(&state, &block)
	}
}

func TestRCompress(t *testing.T) {
	type test struct {
		a, b [16]uint32
	}

	const q = 0x12345678
	const u = 0x87654321

	tests := []test{
		{
			a: [16]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint32{q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q},
			b: [16]uint32{u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			asma, asmb := test.a, test.b
			rcompress(&asma, &asmb)

			goa, gob := test.a, test.b
			compress_pure.RCompress(&goa, &gob)

			if asma != goa || asmb != gob || asmb != test.b {
				t.Error("different result")

				if asma != goa {
					t.Logf(" in A: %08x", test.a)
					t.Logf("got A: %08x", asma)
					t.Logf("exp A: %08x", goa)
				}

				if asmb != gob || asmb != test.b {
					t.Logf(" in B: %08x", test.b)
					t.Logf("got B: %08x", asmb)
					t.Logf("exp B: %08x", gob)
				}
			}
		})
	}
}

func run(a, b [16]uint32, fn func(a, b *[16]uint32)) ([16]uint32, [16]uint32) {
	fn(&a, &b)
	return a, b
}

type Vector [4]*uint32

func (v Vector) Rot(n int) (r Vector) {
	for i, x := range v {
		r[(i-n+len(v))%len(v)] = x
	}
	return r
}

func rcompress3(state, block *[16]uint32) {
	const a, b, c, d, e, f = 10, 11, 12, 13, 14, 15

	v0 := Vector{&state[0], &state[1], &state[2], &state[3]}
	v1 := Vector{&state[4], &state[5], &state[6], &state[7]}
	v2 := Vector{&state[8], &state[9], &state[a], &state[b]}
	v3 := Vector{&state[c], &state[d], &state[e], &state[f]}

	p := func(q1, q2, q3, q4 int) Vector {
		return Vector{&block[q1], &block[q2], &block[q3], &block[q4]}
	}

	// round 1
	mix(v0, v1, v2, v3, p(0, 2, 4, 6), p(1, 3, 5, 7))
	mix(v0, v1.Rot(1), v2.Rot(2), v3.Rot(3), p(8, a, c, e), p(9, b, d, f))

	mix(v0, v1, v2, v3, p(0, 2, 4, 6), p(1, 3, 5, 7))
	mix(v0, v1.Rot(1), v2.Rot(2), v3.Rot(3), p(8, a, c, e), p(9, b, d, f))

	//mix(v0, v1, v2, v3, p(8, a, c, e), p(9, b, d, f))
}

func mix(a, b, c, d, mx, my Vector) {
	for i := 0; i < 4; i++ {
		*a[i] += *b[i] + *mx[i]

		*d[i] ^= *a[i]
		*d[i] = bits.RotateLeft32(*d[i], -16)

		*c[i] += *d[i]

		*b[i] ^= *c[i]
		*b[i] = bits.RotateLeft32(*b[i], -12)

		*a[i] += *b[i] + *my[i]

		*d[i] ^= *a[i]
		*d[i] = bits.RotateLeft32(*d[i], -8)

		*c[i] += *d[i]

		*b[i] = bits.RotateLeft32(*b[i]^*c[i], -7)
	}
}

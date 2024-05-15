package main

import (
	"flag"
	"fmt"
	"os"
)

var mixcall = flag.Bool("mixcall", false, "emit a mix function") // for experimenting

func main() {
	purego := flag.String("purego", "", "purego output file")
	arm64 := flag.String("arm64", "", "arm64 output file")
	flag.Parse()

	if *purego == "" {
		flag.Usage()
		os.Exit(1)
	}

	{
		code := &Go{}
		emitPureGo(code)

		err := code.SaveTo(*purego)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	{
		code := &Arm64{}
		emitArm64(code)

		err := code.SaveTo(*arm64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

type MixFunc func(code Code, a, b, c, d, mx, my Register)

func rcompress(code Code, statep, blockp Register, s [16]Register, m [16]Register, mix MixFunc) {
	const a, b, c, d, e, f = 10, 11, 12, 13, 14, 15

	for i := range s {
		code.Load(s[i], statep.Ix(i))
	}
	for i := range m {
		if !m[i].Eq(blockp.Ix(i)) {
			code.Load(m[i], blockp.Ix(i))
		}
	}

	code.Doc("round 1")
	mix(code, s[0], s[4], s[8], s[c], m[0], m[1])
	mix(code, s[1], s[5], s[9], s[d], m[2], m[3])
	mix(code, s[2], s[6], s[a], s[e], m[4], m[5])
	mix(code, s[3], s[7], s[b], s[f], m[6], m[7])

	mix(code, s[0], s[5], s[a], s[f], m[8], m[9])
	mix(code, s[1], s[6], s[b], s[c], m[a], m[b])
	mix(code, s[2], s[7], s[8], s[d], m[c], m[d])
	mix(code, s[3], s[4], s[9], s[e], m[e], m[f])

	code.Doc("round 2")
	mix(code, s[0], s[4], s[8], s[c], m[2], m[6])
	mix(code, s[1], s[5], s[9], s[d], m[3], m[a])
	mix(code, s[2], s[6], s[a], s[e], m[7], m[0])
	mix(code, s[3], s[7], s[b], s[f], m[4], m[d])

	mix(code, s[0], s[5], s[a], s[f], m[1], m[b])
	mix(code, s[1], s[6], s[b], s[c], m[c], m[5])
	mix(code, s[2], s[7], s[8], s[d], m[9], m[e])
	mix(code, s[3], s[4], s[9], s[e], m[f], m[8])

	code.Doc("round 3")
	mix(code, s[0], s[4], s[8], s[c], m[3], m[4])
	mix(code, s[1], s[5], s[9], s[d], m[a], m[c])
	mix(code, s[2], s[6], s[a], s[e], m[d], m[2])
	mix(code, s[3], s[7], s[b], s[f], m[7], m[e])

	mix(code, s[0], s[5], s[a], s[f], m[6], m[5])
	mix(code, s[1], s[6], s[b], s[c], m[9], m[0])
	mix(code, s[2], s[7], s[8], s[d], m[b], m[f])
	mix(code, s[3], s[4], s[9], s[e], m[8], m[1])

	code.Doc("round 4")
	mix(code, s[0], s[4], s[8], s[c], m[a], m[7])
	mix(code, s[1], s[5], s[9], s[d], m[c], m[9])
	mix(code, s[2], s[6], s[a], s[e], m[e], m[3])
	mix(code, s[3], s[7], s[b], s[f], m[d], m[f])

	mix(code, s[0], s[5], s[a], s[f], m[4], m[0])
	mix(code, s[1], s[6], s[b], s[c], m[b], m[2])
	mix(code, s[2], s[7], s[8], s[d], m[5], m[8])
	mix(code, s[3], s[4], s[9], s[e], m[1], m[6])

	code.Doc("round 5")
	mix(code, s[0], s[4], s[8], s[c], m[c], m[d])
	mix(code, s[1], s[5], s[9], s[d], m[9], m[b])
	mix(code, s[2], s[6], s[a], s[e], m[f], m[a])
	mix(code, s[3], s[7], s[b], s[f], m[e], m[8])

	mix(code, s[0], s[5], s[a], s[f], m[7], m[2])
	mix(code, s[1], s[6], s[b], s[c], m[5], m[3])
	mix(code, s[2], s[7], s[8], s[d], m[0], m[1])
	mix(code, s[3], s[4], s[9], s[e], m[6], m[4])

	code.Doc("round 6")
	mix(code, s[0], s[4], s[8], s[c], m[9], m[e])
	mix(code, s[1], s[5], s[9], s[d], m[b], m[5])
	mix(code, s[2], s[6], s[a], s[e], m[8], m[c])
	mix(code, s[3], s[7], s[b], s[f], m[f], m[1])

	mix(code, s[0], s[5], s[a], s[f], m[d], m[3])
	mix(code, s[1], s[6], s[b], s[c], m[0], m[a])
	mix(code, s[2], s[7], s[8], s[d], m[2], m[6])
	mix(code, s[3], s[4], s[9], s[e], m[4], m[7])

	code.Doc("round 7")
	mix(code, s[0], s[4], s[8], s[c], m[b], m[f])
	mix(code, s[1], s[5], s[9], s[d], m[5], m[0])
	mix(code, s[2], s[6], s[a], s[e], m[1], m[9])
	mix(code, s[3], s[7], s[b], s[f], m[8], m[6])

	mix(code, s[0], s[5], s[a], s[f], m[e], m[a])
	mix(code, s[1], s[6], s[b], s[c], m[2], m[c])
	mix(code, s[2], s[7], s[8], s[d], m[3], m[4])
	mix(code, s[3], s[4], s[9], s[e], m[7], m[d])

	code.Doc("mix upper and lower halves")

	for i := 0; i < 8; i++ {
		// s[8+i] = s8 ^ s[i]
		code.Xor2(statep.Ix(8+i), s[8+i], statep.Ix(i))
	}

	for i := 0; i < 8; i++ {
		// s[0] = s0 ^ s8
		code.Xor2(statep.Ix(i), s[i], s[8+i])
	}
}

func mix(code Code, a, b, c, d, mx, my Register) {
	code.Doc(fmt.Sprintf("mix(%v, %v, %v, %v, %v, %v)", a.Go(), b.Go(), c.Go(), d.Go(), mx.Go(), my.Go()))
	// a += b + mx
	code.AddInto2(a, b, mx)
	// d = bits.RotateLeft32(d^a, -16)
	code.XorIntoAndRotateRight(d, a, 16)
	// c += d
	code.AddInto(c, d)
	// b = bits.RotateLeft32(b^c, -12)
	code.XorIntoAndRotateRight(b, c, 12)
	// a += b + my
	code.AddInto2(a, b, my)
	// d = bits.RotateLeft32(d^a, -8)
	code.XorIntoAndRotateRight(d, a, 8)
	// c += d
	code.AddInto(c, d)
	// b = bits.RotateLeft32(b^c, -7)
	code.XorIntoAndRotateRight(b, c, 7)
}

type Register struct {
	Name  string
	Index int

	Indexed bool
}

func R(name string) Register            { return Register{Name: name} }
func I(name string, index int) Register { return Register{Name: name, Index: index, Indexed: true} }

func (r Register) Eq(v Register) bool {
	return r.Name == v.Name && r.Index == v.Index
}

func (r Register) Ix(k int) Register {
	r.Index += k
	r.Indexed = true
	return r
}

type Code interface {
	Doc(string)
	Printf(format string, args ...any)
	Load(dst, src Register)
	Store(dst Register, src Register)
	AddInto(dst, x Register)
	AddInto2(dst, x, y Register)
	XorInto(dst, x Register)
	Xor2(dst, x, y Register)
	XorIntoAndRotateRight(dst, x Register, amount int)
	RotateRight(dst Register, amount int)
}

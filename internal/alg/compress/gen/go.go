package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"strconv"
)

func (r Register) Go() string {
	if !r.Indexed {
		return r.Name
	}
	return r.Name + "[" + strconv.Itoa(r.Index) + "]"
}

func emitPureGo(code *Go) {
	const a, b, c, d, e, f = 10, 11, 12, 13, 14, 15

	code.Printf("package compress_pure\n\n")

	code.Printf("import \"math/bits\"\n\n")

	code.Printf("func rcompress(s *[16]uint32, m *[16]uint32) {\n")
	code.Printf("	_, _ = m[15], s[15]\n")
	defer code.Printf("}\n")

	var mixer MixFunc = mix
	if *mixcall {
		mixer = mixGoCall

		emitMixGo(code)
	}

	rcompress(code, R("s"), R("m"),
		[16]Register{
			R("s0"), R("s1"), R("s2"), R("s3"),
			R("s4"), R("s5"), R("s6"), R("s7"),
			R("s8"), R("s9"), R("sa"), R("sb"),
			R("sc"), R("sd"), R("se"), R("sf"),
		},
		[16]Register{
			I("m", 0), I("m", 1), I("m", 2), I("m", 3),
			I("m", 4), I("m", 5), I("m", 6), I("m", 7),
			I("m", 8), I("m", 9), I("m", a), I("m", b),
			I("m", c), I("m", d), I("m", e), I("m", f),
		}, mixer)
}

func emitMixGo(code Code) {
	code.Printf("\nfunc mix(a, b, c, d, mx, my uint32) (uint32, uint32, uint32, uint32) {\n")
	mix(code, R("a"), R("b"), R("c"), R("d"), R("mx"), R("my"))
	defer code.Printf("return a, b, c, d\n}\n\b")
}

func mixGoCall(code Code, a, b, c, d, mx, my Register) {
	code.Printf("%v, %v, %v, %v = mix(%v, %v, %v, %v, %v, %v)\n", a, b, c, d, a, b, c, d, mx, my)
}

type Go struct {
	buf bytes.Buffer
}

func (code *Go) Printf(format string, args ...any) {
	fmt.Fprintf(&code.buf, format, args...)
}

func (code *Go) Doc(s string) {
	fmt.Fprintf(&code.buf, "// %v\n", s)
}

func (code *Go) SaveTo(p string) error {
	unformatted := code.buf.Bytes()
	formatted, err := format.Source(unformatted)
	if err != nil {
		err2 := os.WriteFile(p, unformatted, 0644)
		return errors.Join(err, err2)
	}
	return os.WriteFile(p, formatted, 0644)
}

// dst := src
func (code *Go) Load(dst, src Register) {
	code.Printf("%v := %v\n", dst.Go(), src.Go())
}

// dst = src
func (code *Go) Store(dst Register, src Register) {
	code.Printf("%v = %v\n", dst.Go(), src.Go())
}

// dst = dst + x
func (code *Go) AddInto(dst, x Register) {
	code.Printf("%v += %v\n", dst.Go(), x.Go())
}

// dst = dst + x + y
func (code *Go) AddInto2(dst, x, y Register) {
	code.Printf("%v += %v + %v\n", dst.Go(), x.Go(), y.Go())
}

// dst = dst ^ x
func (code *Go) XorInto(dst, x Register) {
	code.Printf("%v ^= %v\n", dst.Go(), x.Go())
}

// dst = x ^ y
func (code *Go) Xor2(dst, x, y Register) {
	code.Printf("%v = %v ^ %v\n", dst.Go(), x.Go(), y.Go())
}

// dst = rotateRight(dst ^ x, amount)
func (code *Go) XorIntoAndRotateRight(dst, x Register, amount int) {
	code.Printf("%v = bits.RotateLeft32(%v ^ %v, %v)\n", dst.Go(), dst.Go(), x.Go(), -amount)
}

// dst = rotateRight(x, amount)
func (code *Go) RotateRight(dst Register, amount int) {
	code.Printf("%v = bits.RotateLeft32(%v, %v)\n", dst.Go(), dst.Go(), -amount)
}

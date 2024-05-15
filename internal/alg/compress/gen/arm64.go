package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func emitArm64(code *Arm64) {
	const a, b, c, d, e, f = 10, 11, 12, 13, 14, 15

	code.Printf("#include \"textflag.h\"\n\n")

	code.Printf(`// func rcompress(state, block *[16]uint32)` + "\n")
	code.Printf("TEXT ·rcompress(SB), NOSPLIT, $0\n")
	defer code.Printf("	RET\n")

	available := []Register{}
	for i := 2; i <= 26; i++ {
		if i == 18 {
			continue
		}

		available = append(available, R("R"+strconv.Itoa(i)))
	}
	pickRegister := func() (Register, bool) {
		if len(available) == 0 {
			return R(""), false
		}
		r := available[0]
		available = available[1:]
		return r, true
	}

	const offset = 2
	sreg := [16]Register{}
	for i := range sreg {
		r, ok := pickRegister()
		if ok {
			sreg[i] = r
		} else {
			sreg[i] = I("R0", i)
		}
	}

	mreg := [16]Register{}
	for i := range sreg {
		r, ok := pickRegister()
		if ok {
			mreg[i] = r
		} else {
			mreg[i] = I("R1", i)
		}
	}

	rcompress(code, R("R0"), R("R1"), sreg, mreg, mix)
}

func (r Register) Arm64() string {
	if !r.Indexed {
		return r.Name
	}
	return strconv.Itoa(r.Index*4) + "(" + r.Name + ")"
}

type Arm64 struct {
	buf bytes.Buffer
}

func (code *Arm64) Printf(format string, args ...any) {
	fmt.Fprintf(&code.buf, format, args...)
}

func (code *Arm64) Doc(s string) {
	fmt.Fprintf(&code.buf, "// %v\n", s)
}

func (code *Arm64) SaveTo(p string) error {
	return os.WriteFile(p, code.buf.Bytes(), 0644)
}

// dst := src
func (code *Arm64) Load(dst, src Register) {
	code.Printf("\tMOVWU %v, %v\n", src.Arm64(), dst.Arm64())
}

// dst = src
func (code *Arm64) Store(dst Register, src Register) {
	code.Printf("\tMOVWU %v, %v\n", src.Arm64(), dst.Arm64())
}

// dst = dst + x
func (code *Arm64) AddInto(dst, x Register) {
	if x.Indexed {
		code.Load(R("R30"), x)
		code.AddInto(dst, R("R30"))
		return
	}

	code.Printf("\tADD %v, %v, %v\n", x.Arm64(), dst.Arm64(), dst.Arm64())
}

// dst = dst + x + y
func (code *Arm64) AddInto2(dst, x, y Register) {
	code.AddInto(dst, x)
	code.AddInto(dst, y)
}

// dst = dst ^ x
func (code *Arm64) XorInto(dst, x Register) {
	if dst.Indexed {
		code.Printf("\tMOVW %v, R30\n", dst.Arm64())
		code.Printf("\tEOR %v, R30, R30\n", x.Arm64())
		code.Printf("\tMOVW R30, %v\n", dst.Arm64())
	} else {
		code.Printf("\tEOR %v, %v, %v\n", x.Arm64(), dst.Arm64(), dst.Arm64())
	}
}

// dst = x ^ y
func (code *Arm64) Xor2(dst, x, y Register) {
	if y.Indexed || dst.Indexed {
		code.Printf("\tMOVW %v, R30\n", y.Arm64())
		code.Printf("\tEOR %v, R30, R30\n", x.Arm64())
		code.Printf("\tMOVW R30, %v\n", dst.Arm64())
	} else {
		code.Printf("\tMOVW R30, %v\n", dst.Arm64())
	}
}

// dst = rotateRight(x, amount)
func (code *Arm64) RotateRight(dst Register, amount int) {
	code.Printf("\tRORW $%v, %v, %v\n", amount, dst.Arm64(), dst.Arm64())
}

// dst = rotateRight(dst ^ x, amount)
func (code *Arm64) XorIntoAndRotateRight(dst, x Register, amount int) {
	code.XorInto(dst, x)
	code.RotateRight(dst, amount)
}

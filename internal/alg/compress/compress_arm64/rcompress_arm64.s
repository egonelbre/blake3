#include "textflag.h"

#define ST0 V0
#define ST1 V1
#define ST2 V2
#define ST3 V3

#define BL0 V4
#define BL1 V5
#define BL2 V6
#define BL3 V7

#define M0 BL0.S[0]
#define M1 BL0.S[1]
#define M2 BL0.S[2]
#define M3 BL0.S[3]
#define M4 BL1.S[0]
#define M5 BL1.S[1]
#define M6 BL1.S[2]
#define M7 BL1.S[3]
#define M8 BL2.S[0]
#define M9 BL2.S[1]
#define MA BL2.S[2]
#define MB BL2.S[3]
#define MC BL3.S[0]
#define MD BL3.S[1]
#define ME BL3.S[2]
#define MF BL3.S[3]

#define MX V8
#define MY V9

#define tmp1 V10
#define tmp2 V11

#define MIX(a, b, c, d, mx, my) \
	\ // a = a + b
	VADD   b.S4,  a.S4, a.S4 \
	\ // a = a + mx
	VADD   mx.S4, a.S4, a.S4 \
	\ // d = d ^ a
	VEOR   a.B16, d.B16, d.B16 \
	\ // a = a + my
	VADD   my.S4, a.S4, a.S4  \
	\ // d = bits.RotateLeft32(d, -16)
	VREV32 d.H8,  d.H8 \
	\ // c = c + d
	VADD   d.S4, c.S4, c.S4  \
	\ // b = b ^ c
	VEOR   c.B16, b.B16, b.B16 \
	\ // b = bits.RotateLeft32(b, -12)
	VSHL   $20, b.S4,  tmp1.S4 \
	VSRI   $12, b.S4,  tmp1.S4 \
	VMOV   tmp1.B16, b.B16 \
	\ // a = a + b
	VADD   b.S4, a.S4, a.S4  \
	\ // d = d ^ a
	VEOR   a.B16, d.B16, d.B16 \
	\ // d = bits.RotateLeft32(d, -8)
	VSHL   $24, d.S4,  tmp1.S4 \
	VSRI   $8,  d.S4,  tmp1.S4 \
	VMOV   tmp1.B16, d.B16 \
	\ // c = c + d
	VADD   d.S4, c.S4, c.S4  \
	\ // b = b ^ c
	VEOR   c.B16, b.B16, b.B16 \
	\ // b = bits.RotateLeft32(b, -7)
	VSHL   $25, b.S4,  tmp1.S4 \
	VSRI   $7,  b.S4,  tmp1.S4 \
	VMOV   tmp1.B16, b.B16

#define SET(into, a, b, c, d) \
	VMOV a, into.S[0] \
	VMOV b, into.S[1] \
	VMOV c, into.S[2] \
	VMOV d, into.S[3]

#define ROT1(x) VEXT  $4, x.B16, x.B16, x.B16
#define ROT2(x) VEXT  $8, x.B16, x.B16, x.B16
#define ROT3(x) VEXT $12, x.B16, x.B16, x.B16

#define ROUNDA(q1, q2, q3, q4, u1, u2, u3, u4) \
	SET(MX, q1, q2, q3, q4) \
	SET(MY, u1, u2, u3, u4) \
	MIX(ST0, ST1, ST2, ST3, MX, MY)

#define ROUNDB(q1, q2, q3, q4, u1, u2, u3, u4) \
	SET(MX, q1, q2, q3, q4) \
	SET(MY, u1, u2, u3, u4) \
	ROT1(ST1); ROT2(ST2); ROT3(ST3) \
	MIX(ST0, ST1, ST2, ST3, MX, MY) \
	ROT3(ST1); ROT2(ST2); ROT1(ST3)

// func rcompress(state, block *[16]uint32)
TEXT ·rcompress(SB), NOSPLIT, $0-16
	MOVD state+0(FP), R0
	MOVD block+8(FP), R1

	VLD1 (R0), [ST0.S4, ST1.S4, ST2.S4, ST3.S4]
	VLD1 (R1), [BL0.S4, BL1.S4, BL2.S4, BL3.S4]

	// Round 1
	ROUNDA(M0, M2, M4, M6, M1, M3, M5, M7)
	ROUNDB(M8, MA, MC, ME, M9, MB, MD, MF)

	// Round 2
	ROUNDA(M2, M3, M7, M4, M6, MA, M0, MD)
	ROUNDB(M1, MC, M9, MF, MB, M5, ME, M8)

	// Round 3
	ROUNDA(M3, MA, MD, M7, M4, MC, M2, ME)
	ROUNDB(M6, M9, MB, M8, M5, M0, MF, M1)

	// Round 4
	ROUNDA(MA, MC, ME, MD, M7, M9, M3, MF)
	ROUNDB(M4, MB, M5, M1, M0, M2, M8, M6)

	// Round 5
	ROUNDA(MC, M9, MF, ME, MD, MB, MA, M8)
	ROUNDB(M7, M5, M0, M6, M2, M3, M1, M4)

	// Round 6
	ROUNDA(M9, MB, M8, MF, ME, M5, MC, M1)
	ROUNDB(MD, M0, M2, M4, M3, MA, M6, M7)

	// Round 7
	ROUNDA(MB, M5, M1, M8, MF, M0, M9, M6)
	ROUNDB(ME, M2, M3, M7, MA, MC, M4, MD)

	// mix upper and lower halves

	VLD1 (R0), [V16.S4, V17.S4, V18.S4, V19.S4]

	VEOR ST2.B16, V16.B16, V18.B16
	VEOR ST3.B16, V17.B16, V19.B16

	VEOR ST0.B16, ST2.B16, V16.B16
	VEOR ST1.B16, ST3.B16, V17.B16

	//VST1 [ST0.S4, ST1.S4, ST2.S4, ST3.S4], (R0)
	VST1 [V16.S4, V17.S4, V18.S4, V19.S4], (R0)

	RET

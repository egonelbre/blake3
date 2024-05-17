#include "textflag.h"

// #define USE_VTBL_ROTATION

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

#define RSCHEDULETABLE R12

#define MAX V8
#define MAY V9
#define MBX V10
#define MBY V11
#define MX V8
#define MY V9

#define tmp1 V30
#define tmp2 V31

#define RotS1 V12
#define RotS2 V13
#define RotS3 V14

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


#ifdef USE_VTBL_ROTATION
#define ROT1(x) VTBL RotS1.B16, [x.B16], x.B16
#define ROT2(x) VTBL RotS2.B16, [x.B16], x.B16
#define ROT3(x) VTBL RotS3.B16, [x.B16], x.B16
#else
#define ROT1(x) VEXT  $4, x.B16, x.B16, x.B16
#define ROT2(x) VEXT  $8, x.B16, x.B16, x.B16
#define ROT3(x) VEXT $12, x.B16, x.B16, x.B16
#endif

#define ROTATE(x, y, z) \
	ROT1(x); ROT2(y); ROT3(z)
#define UNROTATE(x, y, z) \
	ROT3(x); ROT2(y); ROT1(z)

// func rcompress(state, block *[16]uint32)
TEXT ·rcompress(SB), NOSPLIT, $0-16
	MOVD state+0(FP), R0
	MOVD block+8(FP), R1

	#ifdef USE_VTBL_ROTATION
	MOVD $·rotationTable(SB), R11
	VLD1 (R11), [RotS1.B16, RotS2.B16, RotS3.B16]
	#endif

	MOVD $·messageSchedule(SB), RSCHEDULETABLE

	VLD1 (R0), [ST0.S4, ST1.S4, ST2.S4, ST3.S4]
	VLD1 (R1), [BL0.S4, BL1.S4, BL2.S4, BL3.S4]

	// Round 1
	MOVD $0x7, R5
rounds:

	VLD1 (RSCHEDULETABLE), [V16.B16, V17.B16, V18.B16, V19.B16]
	VTBL V16.B16, [BL0.B16, BL1.B16, BL2.B16, BL3.B16], MAX.B16
	VTBL V17.B16, [BL0.B16, BL1.B16, BL2.B16, BL3.B16], MAY.B16
	MIX(ST0, ST1, ST2, ST3, MAX, MAY)

	ROTATE(ST1, ST2, ST3)
	VTBL V18.B16, [BL0.B16, BL1.B16, BL2.B16, BL3.B16], MBX.B16
	VTBL V19.B16, [BL0.B16, BL1.B16, BL2.B16, BL3.B16], MBY.B16
	MIX(ST0, ST1, ST2, ST3, MBX, MBY)
	UNROTATE(ST1, ST2, ST3)

	ADD $0x40, RSCHEDULETABLE, RSCHEDULETABLE
	SUB $1, R5
	CBNZ R5, rounds

	// mix upper and lower halves

	VLD1 (R0), [V16.S4, V17.S4, V18.S4, V19.S4]

	VEOR ST2.B16, V16.B16, V18.B16
	VEOR ST3.B16, V17.B16, V19.B16

	VEOR ST0.B16, ST2.B16, V16.B16
	VEOR ST1.B16, ST3.B16, V17.B16

	VST1 [V16.S4, V17.S4, V18.S4, V19.S4], (R0)

	RET

DATA	·rotationTable+0x00(SB)/4, $0x07060504
DATA	·rotationTable+0x04(SB)/4, $0x0b0a0908
DATA	·rotationTable+0x08(SB)/4, $0x0f0e0d0c
DATA	·rotationTable+0x0c(SB)/4, $0x03020100

DATA	·rotationTable+0x10(SB)/4, $0x0b0a0908
DATA	·rotationTable+0x14(SB)/4, $0x0f0e0d0c
DATA	·rotationTable+0x18(SB)/4, $0x03020100
DATA	·rotationTable+0x1c(SB)/4, $0x07060504

DATA	·rotationTable+0x20(SB)/4, $0x0f0e0d0c
DATA	·rotationTable+0x24(SB)/4, $0x03020100
DATA	·rotationTable+0x28(SB)/4, $0x07060504
DATA	·rotationTable+0x2c(SB)/4, $0x0b0a0908
GLOBL	·rotationTable(SB), NOPTR|RODATA, $48

// MessageSchedule table
// Round 1, Part 1
DATA    ·messageSchedule+0x000(SB)/4, $0x03020100 // MX[0] = 0
DATA    ·messageSchedule+0x004(SB)/4, $0x0b0a0908 // MX[1] = 2
DATA    ·messageSchedule+0x008(SB)/4, $0x13121110 // MX[2] = 4
DATA    ·messageSchedule+0x00c(SB)/4, $0x1b1a1918 // MX[3] = 6
DATA    ·messageSchedule+0x010(SB)/4, $0x07060504 // MY[0] = 1
DATA    ·messageSchedule+0x014(SB)/4, $0x0f0e0d0c // MY[1] = 3
DATA    ·messageSchedule+0x018(SB)/4, $0x17161514 // MY[2] = 5
DATA    ·messageSchedule+0x01c(SB)/4, $0x1f1e1d1c // MY[3] = 7
// Round 1, Part 2
DATA    ·messageSchedule+0x020(SB)/4, $0x23222120 // MX[0] = 8
DATA    ·messageSchedule+0x024(SB)/4, $0x2b2a2928 // MX[1] = 10
DATA    ·messageSchedule+0x028(SB)/4, $0x33323130 // MX[2] = 12
DATA    ·messageSchedule+0x02c(SB)/4, $0x3b3a3938 // MX[3] = 14
DATA    ·messageSchedule+0x030(SB)/4, $0x27262524 // MY[0] = 9
DATA    ·messageSchedule+0x034(SB)/4, $0x2f2e2d2c // MY[1] = 11
DATA    ·messageSchedule+0x038(SB)/4, $0x37363534 // MY[2] = 13
DATA    ·messageSchedule+0x03c(SB)/4, $0x3f3e3d3c // MY[3] = 15
// Round 2, Part 1
DATA    ·messageSchedule+0x040(SB)/4, $0x0b0a0908 // MX[0] = 2
DATA    ·messageSchedule+0x044(SB)/4, $0x0f0e0d0c // MX[1] = 3
DATA    ·messageSchedule+0x048(SB)/4, $0x1f1e1d1c // MX[2] = 7
DATA    ·messageSchedule+0x04c(SB)/4, $0x13121110 // MX[3] = 4
DATA    ·messageSchedule+0x050(SB)/4, $0x1b1a1918 // MY[0] = 6
DATA    ·messageSchedule+0x054(SB)/4, $0x2b2a2928 // MY[1] = 10
DATA    ·messageSchedule+0x058(SB)/4, $0x03020100 // MY[2] = 0
DATA    ·messageSchedule+0x05c(SB)/4, $0x37363534 // MY[3] = 13
// Round 2, Part 2
DATA    ·messageSchedule+0x060(SB)/4, $0x07060504 // MX[0] = 1
DATA    ·messageSchedule+0x064(SB)/4, $0x33323130 // MX[1] = 12
DATA    ·messageSchedule+0x068(SB)/4, $0x27262524 // MX[2] = 9
DATA    ·messageSchedule+0x06c(SB)/4, $0x3f3e3d3c // MX[3] = 15
DATA    ·messageSchedule+0x070(SB)/4, $0x2f2e2d2c // MY[0] = 11
DATA    ·messageSchedule+0x074(SB)/4, $0x17161514 // MY[1] = 5
DATA    ·messageSchedule+0x078(SB)/4, $0x3b3a3938 // MY[2] = 14
DATA    ·messageSchedule+0x07c(SB)/4, $0x23222120 // MY[3] = 8
// Round 3, Part 1
DATA    ·messageSchedule+0x080(SB)/4, $0x0f0e0d0c // MX[0] = 3
DATA    ·messageSchedule+0x084(SB)/4, $0x2b2a2928 // MX[1] = 10
DATA    ·messageSchedule+0x088(SB)/4, $0x37363534 // MX[2] = 13
DATA    ·messageSchedule+0x08c(SB)/4, $0x1f1e1d1c // MX[3] = 7
DATA    ·messageSchedule+0x090(SB)/4, $0x13121110 // MY[0] = 4
DATA    ·messageSchedule+0x094(SB)/4, $0x33323130 // MY[1] = 12
DATA    ·messageSchedule+0x098(SB)/4, $0x0b0a0908 // MY[2] = 2
DATA    ·messageSchedule+0x09c(SB)/4, $0x3b3a3938 // MY[3] = 14
// Round 3, Part 2
DATA    ·messageSchedule+0x0a0(SB)/4, $0x1b1a1918 // MX[0] = 6
DATA    ·messageSchedule+0x0a4(SB)/4, $0x27262524 // MX[1] = 9
DATA    ·messageSchedule+0x0a8(SB)/4, $0x2f2e2d2c // MX[2] = 11
DATA    ·messageSchedule+0x0ac(SB)/4, $0x23222120 // MX[3] = 8
DATA    ·messageSchedule+0x0b0(SB)/4, $0x17161514 // MY[0] = 5
DATA    ·messageSchedule+0x0b4(SB)/4, $0x03020100 // MY[1] = 0
DATA    ·messageSchedule+0x0b8(SB)/4, $0x3f3e3d3c // MY[2] = 15
DATA    ·messageSchedule+0x0bc(SB)/4, $0x07060504 // MY[3] = 1
// Round 4, Part 1
DATA    ·messageSchedule+0x0c0(SB)/4, $0x2b2a2928 // MX[0] = 10
DATA    ·messageSchedule+0x0c4(SB)/4, $0x33323130 // MX[1] = 12
DATA    ·messageSchedule+0x0c8(SB)/4, $0x3b3a3938 // MX[2] = 14
DATA    ·messageSchedule+0x0cc(SB)/4, $0x37363534 // MX[3] = 13
DATA    ·messageSchedule+0x0d0(SB)/4, $0x1f1e1d1c // MY[0] = 7
DATA    ·messageSchedule+0x0d4(SB)/4, $0x27262524 // MY[1] = 9
DATA    ·messageSchedule+0x0d8(SB)/4, $0x0f0e0d0c // MY[2] = 3
DATA    ·messageSchedule+0x0dc(SB)/4, $0x3f3e3d3c // MY[3] = 15
// Round 4, Part 2
DATA    ·messageSchedule+0x0e0(SB)/4, $0x13121110 // MX[0] = 4
DATA    ·messageSchedule+0x0e4(SB)/4, $0x2f2e2d2c // MX[1] = 11
DATA    ·messageSchedule+0x0e8(SB)/4, $0x17161514 // MX[2] = 5
DATA    ·messageSchedule+0x0ec(SB)/4, $0x07060504 // MX[3] = 1
DATA    ·messageSchedule+0x0f0(SB)/4, $0x03020100 // MY[0] = 0
DATA    ·messageSchedule+0x0f4(SB)/4, $0x0b0a0908 // MY[1] = 2
DATA    ·messageSchedule+0x0f8(SB)/4, $0x23222120 // MY[2] = 8
DATA    ·messageSchedule+0x0fc(SB)/4, $0x1b1a1918 // MY[3] = 6
// Round 5, Part 1
DATA    ·messageSchedule+0x100(SB)/4, $0x33323130 // MX[0] = 12
DATA    ·messageSchedule+0x104(SB)/4, $0x27262524 // MX[1] = 9
DATA    ·messageSchedule+0x108(SB)/4, $0x3f3e3d3c // MX[2] = 15
DATA    ·messageSchedule+0x10c(SB)/4, $0x3b3a3938 // MX[3] = 14
DATA    ·messageSchedule+0x110(SB)/4, $0x37363534 // MY[0] = 13
DATA    ·messageSchedule+0x114(SB)/4, $0x2f2e2d2c // MY[1] = 11
DATA    ·messageSchedule+0x118(SB)/4, $0x2b2a2928 // MY[2] = 10
DATA    ·messageSchedule+0x11c(SB)/4, $0x23222120 // MY[3] = 8
// Round 5, Part 2
DATA    ·messageSchedule+0x120(SB)/4, $0x1f1e1d1c // MX[0] = 7
DATA    ·messageSchedule+0x124(SB)/4, $0x17161514 // MX[1] = 5
DATA    ·messageSchedule+0x128(SB)/4, $0x03020100 // MX[2] = 0
DATA    ·messageSchedule+0x12c(SB)/4, $0x1b1a1918 // MX[3] = 6
DATA    ·messageSchedule+0x130(SB)/4, $0x0b0a0908 // MY[0] = 2
DATA    ·messageSchedule+0x134(SB)/4, $0x0f0e0d0c // MY[1] = 3
DATA    ·messageSchedule+0x138(SB)/4, $0x07060504 // MY[2] = 1
DATA    ·messageSchedule+0x13c(SB)/4, $0x13121110 // MY[3] = 4
// Round 6, Part 1
DATA    ·messageSchedule+0x140(SB)/4, $0x27262524 // MX[0] = 9
DATA    ·messageSchedule+0x144(SB)/4, $0x2f2e2d2c // MX[1] = 11
DATA    ·messageSchedule+0x148(SB)/4, $0x23222120 // MX[2] = 8
DATA    ·messageSchedule+0x14c(SB)/4, $0x3f3e3d3c // MX[3] = 15
DATA    ·messageSchedule+0x150(SB)/4, $0x3b3a3938 // MY[0] = 14
DATA    ·messageSchedule+0x154(SB)/4, $0x17161514 // MY[1] = 5
DATA    ·messageSchedule+0x158(SB)/4, $0x33323130 // MY[2] = 12
DATA    ·messageSchedule+0x15c(SB)/4, $0x07060504 // MY[3] = 1
// Round 6, Part 2
DATA    ·messageSchedule+0x160(SB)/4, $0x37363534 // MX[0] = 13
DATA    ·messageSchedule+0x164(SB)/4, $0x03020100 // MX[1] = 0
DATA    ·messageSchedule+0x168(SB)/4, $0x0b0a0908 // MX[2] = 2
DATA    ·messageSchedule+0x16c(SB)/4, $0x13121110 // MX[3] = 4
DATA    ·messageSchedule+0x170(SB)/4, $0x0f0e0d0c // MY[0] = 3
DATA    ·messageSchedule+0x174(SB)/4, $0x2b2a2928 // MY[1] = 10
DATA    ·messageSchedule+0x178(SB)/4, $0x1b1a1918 // MY[2] = 6
DATA    ·messageSchedule+0x17c(SB)/4, $0x1f1e1d1c // MY[3] = 7
// Round 7, Part 1
DATA    ·messageSchedule+0x180(SB)/4, $0x2f2e2d2c // MX[0] = 11
DATA    ·messageSchedule+0x184(SB)/4, $0x17161514 // MX[1] = 5
DATA    ·messageSchedule+0x188(SB)/4, $0x07060504 // MX[2] = 1
DATA    ·messageSchedule+0x18c(SB)/4, $0x23222120 // MX[3] = 8
DATA    ·messageSchedule+0x190(SB)/4, $0x3f3e3d3c // MY[0] = 15
DATA    ·messageSchedule+0x194(SB)/4, $0x03020100 // MY[1] = 0
DATA    ·messageSchedule+0x198(SB)/4, $0x27262524 // MY[2] = 9
DATA    ·messageSchedule+0x19c(SB)/4, $0x1b1a1918 // MY[3] = 6
// Round 7, Part 2
DATA    ·messageSchedule+0x1a0(SB)/4, $0x3b3a3938 // MX[0] = 14
DATA    ·messageSchedule+0x1a4(SB)/4, $0x0b0a0908 // MX[1] = 2
DATA    ·messageSchedule+0x1a8(SB)/4, $0x0f0e0d0c // MX[2] = 3
DATA    ·messageSchedule+0x1ac(SB)/4, $0x1f1e1d1c // MX[3] = 7
DATA    ·messageSchedule+0x1b0(SB)/4, $0x2b2a2928 // MY[0] = 10
DATA    ·messageSchedule+0x1b4(SB)/4, $0x33323130 // MY[1] = 12
DATA    ·messageSchedule+0x1b8(SB)/4, $0x13121110 // MY[2] = 4
DATA    ·messageSchedule+0x1bc(SB)/4, $0x37363534 // MY[3] = 13
GLOBL   ·messageSchedule(SB), NOPTR|RODATA, $0x1c0

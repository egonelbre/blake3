package compress_pure

import "math/bits"

func RCompress(s *[16]uint32, m *[16]uint32) { rcompress(s, m) }

func rcompress(s *[16]uint32, m *[16]uint32) {
	_, _ = m[15], s[15]
	s0 := s[0]
	s1 := s[1]
	s2 := s[2]
	s3 := s[3]
	s4 := s[4]
	s5 := s[5]
	s6 := s[6]
	s7 := s[7]
	s8 := s[8]
	s9 := s[9]
	sa := s[10]
	sb := s[11]
	sc := s[12]
	sd := s[13]
	se := s[14]
	sf := s[15]
	// round 1
	// mix(s0, s4, s8, sc, m[0], m[1])
	s0 += s4 + m[0]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[1]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[2], m[3])
	s1 += s5 + m[2]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[3]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[4], m[5])
	s2 += s6 + m[4]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[5]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[6], m[7])
	s3 += s7 + m[6]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[7]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[8], m[9])
	s0 += s5 + m[8]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[9]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[10], m[11])
	s1 += s6 + m[10]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[11]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[12], m[13])
	s2 += s7 + m[12]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[13]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[14], m[15])
	s3 += s4 + m[14]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[15]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// round 2
	// mix(s0, s4, s8, sc, m[2], m[6])
	s0 += s4 + m[2]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[6]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[3], m[10])
	s1 += s5 + m[3]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[10]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[7], m[0])
	s2 += s6 + m[7]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[0]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[4], m[13])
	s3 += s7 + m[4]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[13]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[1], m[11])
	s0 += s5 + m[1]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[11]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[12], m[5])
	s1 += s6 + m[12]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[5]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[9], m[14])
	s2 += s7 + m[9]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[14]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[15], m[8])
	s3 += s4 + m[15]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[8]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// round 3
	// mix(s0, s4, s8, sc, m[3], m[4])
	s0 += s4 + m[3]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[4]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[10], m[12])
	s1 += s5 + m[10]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[12]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[13], m[2])
	s2 += s6 + m[13]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[2]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[7], m[14])
	s3 += s7 + m[7]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[14]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[6], m[5])
	s0 += s5 + m[6]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[5]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[9], m[0])
	s1 += s6 + m[9]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[0]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[11], m[15])
	s2 += s7 + m[11]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[15]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[8], m[1])
	s3 += s4 + m[8]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[1]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// round 4
	// mix(s0, s4, s8, sc, m[10], m[7])
	s0 += s4 + m[10]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[7]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[12], m[9])
	s1 += s5 + m[12]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[9]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[14], m[3])
	s2 += s6 + m[14]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[3]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[13], m[15])
	s3 += s7 + m[13]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[15]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[4], m[0])
	s0 += s5 + m[4]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[0]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[11], m[2])
	s1 += s6 + m[11]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[2]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[5], m[8])
	s2 += s7 + m[5]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[8]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[1], m[6])
	s3 += s4 + m[1]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[6]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// round 5
	// mix(s0, s4, s8, sc, m[12], m[13])
	s0 += s4 + m[12]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[13]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[9], m[11])
	s1 += s5 + m[9]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[11]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[15], m[10])
	s2 += s6 + m[15]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[10]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[14], m[8])
	s3 += s7 + m[14]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[8]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[7], m[2])
	s0 += s5 + m[7]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[2]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[5], m[3])
	s1 += s6 + m[5]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[3]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[0], m[1])
	s2 += s7 + m[0]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[1]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[6], m[4])
	s3 += s4 + m[6]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[4]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// round 6
	// mix(s0, s4, s8, sc, m[9], m[14])
	s0 += s4 + m[9]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[14]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[11], m[5])
	s1 += s5 + m[11]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[5]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[8], m[12])
	s2 += s6 + m[8]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[12]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[15], m[1])
	s3 += s7 + m[15]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[1]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[13], m[3])
	s0 += s5 + m[13]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[3]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[0], m[10])
	s1 += s6 + m[0]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[10]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[2], m[6])
	s2 += s7 + m[2]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[6]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[4], m[7])
	s3 += s4 + m[4]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[7]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// round 7
	// mix(s0, s4, s8, sc, m[11], m[15])
	s0 += s4 + m[11]
	sc = bits.RotateLeft32(sc^s0, -16)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -12)
	s0 += s4 + m[15]
	sc = bits.RotateLeft32(sc^s0, -8)
	s8 += sc
	s4 = bits.RotateLeft32(s4^s8, -7)
	// mix(s1, s5, s9, sd, m[5], m[0])
	s1 += s5 + m[5]
	sd = bits.RotateLeft32(sd^s1, -16)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -12)
	s1 += s5 + m[0]
	sd = bits.RotateLeft32(sd^s1, -8)
	s9 += sd
	s5 = bits.RotateLeft32(s5^s9, -7)
	// mix(s2, s6, sa, se, m[1], m[9])
	s2 += s6 + m[1]
	se = bits.RotateLeft32(se^s2, -16)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -12)
	s2 += s6 + m[9]
	se = bits.RotateLeft32(se^s2, -8)
	sa += se
	s6 = bits.RotateLeft32(s6^sa, -7)
	// mix(s3, s7, sb, sf, m[8], m[6])
	s3 += s7 + m[8]
	sf = bits.RotateLeft32(sf^s3, -16)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -12)
	s3 += s7 + m[6]
	sf = bits.RotateLeft32(sf^s3, -8)
	sb += sf
	s7 = bits.RotateLeft32(s7^sb, -7)
	// mix(s0, s5, sa, sf, m[14], m[10])
	s0 += s5 + m[14]
	sf = bits.RotateLeft32(sf^s0, -16)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -12)
	s0 += s5 + m[10]
	sf = bits.RotateLeft32(sf^s0, -8)
	sa += sf
	s5 = bits.RotateLeft32(s5^sa, -7)
	// mix(s1, s6, sb, sc, m[2], m[12])
	s1 += s6 + m[2]
	sc = bits.RotateLeft32(sc^s1, -16)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -12)
	s1 += s6 + m[12]
	sc = bits.RotateLeft32(sc^s1, -8)
	sb += sc
	s6 = bits.RotateLeft32(s6^sb, -7)
	// mix(s2, s7, s8, sd, m[3], m[4])
	s2 += s7 + m[3]
	sd = bits.RotateLeft32(sd^s2, -16)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -12)
	s2 += s7 + m[4]
	sd = bits.RotateLeft32(sd^s2, -8)
	s8 += sd
	s7 = bits.RotateLeft32(s7^s8, -7)
	// mix(s3, s4, s9, se, m[7], m[13])
	s3 += s4 + m[7]
	se = bits.RotateLeft32(se^s3, -16)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -12)
	s3 += s4 + m[13]
	se = bits.RotateLeft32(se^s3, -8)
	s9 += se
	s4 = bits.RotateLeft32(s4^s9, -7)
	// mix upper and lower halves
	s[8] = s8 ^ s[0]
	s[9] = s9 ^ s[1]
	s[10] = sa ^ s[2]
	s[11] = sb ^ s[3]
	s[12] = sc ^ s[4]
	s[13] = sd ^ s[5]
	s[14] = se ^ s[6]
	s[15] = sf ^ s[7]
	s[0] = s0 ^ s8
	s[1] = s1 ^ s9
	s[2] = s2 ^ sa
	s[3] = s3 ^ sb
	s[4] = s4 ^ sc
	s[5] = s5 ^ sd
	s[6] = s6 ^ se
	s[7] = s7 ^ sf
}

package main

import "fmt"

func main() {
	fmt.Printf("// MessageSchedule table\n")
	tableOffset := 0

	tableName := "·messageSchedule"
	prefix := "DATA\t" + tableName + "+0x"
	for round, schedule := range Schedule {
		for part := 0; part < 2; part++ {
			fmt.Printf("// Round %v, Part %v\n", round+1, part+1)

			mx := [4]byte{schedule[0+part*8], schedule[2+part*8], schedule[4+part*8], schedule[6+part*8]}
			my := [4]byte{schedule[1+part*8], schedule[3+part*8], schedule[5+part*8], schedule[7+part*8]}

			for i := range mx {
				fmt.Printf("%s%03x(SB)/4, $0x%02x%02x%02x%02x // MX[%v] = %v\n", prefix, tableOffset, mx[i]*4+3, mx[i]*4+2, mx[i]*4+1, mx[0]*4, i, mx[i])
				tableOffset += 4
			}
			for i := range my {
				fmt.Printf("%s%03x(SB)/4, $0x%02x%02x%02x%02x // MY[%v] = %v\n", prefix, tableOffset, my[i]*4+3, my[i]*4+2, my[i]*4+1, my[0]*4, i, my[i])
				tableOffset += 4
			}
		}
	}
	fmt.Printf("GLOBL\t%v(SB), NOPTR|RODATA, $0x%x", tableName, tableOffset)
}

var Schedule = [7][16]byte{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	{2, 6, 3, 10, 7, 0, 4, 13, 1, 11, 12, 5, 9, 14, 15, 8},
	{3, 4, 10, 12, 13, 2, 7, 14, 6, 5, 9, 0, 11, 15, 8, 1},
	{10, 7, 12, 9, 14, 3, 13, 15, 4, 0, 11, 2, 5, 8, 1, 6},
	{12, 13, 9, 11, 15, 10, 14, 8, 7, 2, 5, 3, 0, 1, 6, 4},
	{9, 14, 11, 5, 8, 12, 15, 1, 13, 3, 0, 10, 2, 6, 4, 7},
	{11, 15, 5, 0, 1, 9, 8, 6, 14, 10, 2, 12, 3, 4, 7, 13},
}

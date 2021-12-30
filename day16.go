package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parse4(strs []string) (ret []int64) {
	ret = make([]int64, len(strs))
	for i := range strs {
		ret[i], _ = strconv.ParseInt(strs[i], 10, 0)
	}
	return
}

type Instr int64

const (
	addr Instr = iota
	addi
	mulr
	muli
	banr
	bani
	borr
	bori
	setr
	seti
	gtir
	gtri
	gtrr
	eqir
	eqri
	eqrr
)

func exec1(op []int64, regs []int64, oo map[int64]Instr) {
	switch oo[op[0]] {
	case addr:
		regs[op[3]] = regs[op[1]] + regs[op[2]]
	case addi:
		regs[op[3]] = regs[op[1]] + op[2]
	case mulr:
		regs[op[3]] = regs[op[1]] * regs[op[2]]
	case muli:
		regs[op[3]] = regs[op[1]] * op[2]
	case banr:
		regs[op[3]] = regs[op[1]] & regs[op[2]]
	case bani:
		regs[op[3]] = regs[op[1]] & op[2]
	case borr:
		regs[op[3]] = regs[op[1]] | regs[op[2]]
	case bori:
		regs[op[3]] = regs[op[1]] | op[2]
	case setr:
		regs[op[3]] = regs[op[1]]
	case seti:
		regs[op[3]] = op[1]
	case gtir:
		if op[1] > regs[op[2]] {
			regs[op[3]] = 1
		} else {
			regs[op[3]] = 0
		}
	case gtri:
		if regs[op[1]] > op[2] {
			regs[op[3]] = 1
		} else {
			regs[op[3]] = 0
		}
	case gtrr:
		if regs[op[1]] > regs[op[2]] {
			regs[op[3]] = 1
		} else {
			regs[op[3]] = 0
		}
	case eqir:
		if op[1] == regs[op[2]] {
			regs[op[3]] = 1
		} else {
			regs[op[3]] = 0
		}
	case eqri:
		if regs[op[1]] == op[2] {
			regs[op[3]] = 1
		} else {
			regs[op[3]] = 0
		}
	case eqrr:
		if regs[op[1]] == regs[op[2]] {
			regs[op[3]] = 1
		} else {
			regs[op[3]] = 0
		}
	}
}

func (x Aoc) Day16(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`(Before|After):\s+\[(\d+),\s*(\d+),\s*(\d+),\s*(\d+)\]`)
	capture := false
	var before, after, opcode []int64
	ooo := make(map[int64]Instr, 16)
	program := make([][]int64, 0, 1000)
	bad := make([][]bool, 16)
	for i := addr; i <= eqrr; i++ {
		bad[i] = make([]bool, 16)
	}
	tot1 := 0
	descr := map[Instr]string{addr: "ADD.R", addi: "ADD.I", mulr: "MUL.R", muli: "MUL.I",
		banr: "BAN.R", bani: "BAN.I", borr: "BOR.R", bori: "BOR.I", setr: "SET.R", seti: "SET.I",
		gtir: "GT.IR", gtri: "GT.RI", gtrr: "GT.RR", eqir: "EQ.IR", eqri: "EQ.RI", eqrr: "EQ.RR"}
	good := addr

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			g := r.FindStringSubmatch(scanner.Text())
			if capture {
				after = parse4(g[2:])
				tsame := 0
				for i := addr; i <= eqrr; i++ {
					ooo[opcode[0]] = i
					tmp := make([]int64, 4)
					copy(tmp, before)
					exec1(opcode, tmp, ooo)
					same := true
					for j := range after {
						if after[j] != tmp[j] {
							same = false
							break
						}
					}
					if same {
						good = i
						tsame++
					} else {
						bad[opcode[0]][i] = true
					}
				}
				if tsame >= 3 {
					tot1++
				}
				capture = false
			} else {
				before = parse4(g[2:])
				capture = true
			}
		} else {
			nums := strings.Split(scanner.Text(), " ")
			if len(nums) == 4 {
				opcode = parse4(nums)
				if !capture {
					program = append(program, opcode)
				}
			}
		}
	}
	ooo = make(map[int64]Instr, 16)
	for i := int64(0); i < 16; i++ {
		_, done := ooo[i]
		if done {
			continue
		}
		gcnt := 0
		for j := addr; j <= eqrr; j++ {
			if !bad[i][j] {
				gcnt++
				good = j
			}
		}
		if gcnt == 1 {
			fmt.Printf("%d must be %s (%d)\n", i, descr[good], good)
			ooo[i] = good
			for j := 0; j < 16; j++ {
				bad[j][good] = true
			}
			i = -1
		}
	}
	fmt.Printf("Mapped %d opcodes\n", len(ooo))
	regs := []int64{0, 0, 0, 0}
	for _, op := range program {
		// fmt.Printf("[%02d] %s %d %d %d | %v\n", op[0], descr[ooo[op[0]]], op[1], op[2], op[3], regs)
		exec1(op, regs, ooo)
	}
	fmt.Printf("Final registers: %v\n", regs)
	fmt.Printf("Day 16 Part1: %d\n", tot1)
	fmt.Printf("Day 16 Part2: %d\n", regs[0])
}

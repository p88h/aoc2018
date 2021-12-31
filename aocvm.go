package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Opcode int64

const (
	addr Opcode = iota
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

type Instruction struct {
	op      Opcode
	a, b, c int64
}

type AocVM struct {
	regs []int64
	ip   int
	rip  int
}

var opdescr = map[Opcode]string{addr: "ADD.R", addi: "ADD.I", mulr: "MUL.R", muli: "MUL.I",
	banr: "BAN.R", bani: "BAN.I", borr: "BOR.R", bori: "BOR.I", setr: "SET.R", seti: "SET.I",
	gtir: "GT.IR", gtri: "GT.RI", gtrr: "GT.RR", eqir: "EQ.IR", eqri: "EQ.RI", eqrr: "EQ.RR"}

func (vm *AocVM) Describe(insn Instruction) string {
	return fmt.Sprintf("[%02d@%04d] %s %d %d %d | %v", insn.op, vm.ip, opdescr[insn.op], insn.a, insn.b, insn.c, vm.regs)
}

var opnames = map[string]Opcode{"addr": addr, "addi": addi, "mulr": mulr, "muli": muli,
	"banr": banr, "bani": bani, "borr": borr, "bori": bori, "setr": setr, "seti": seti,
	"gtir": gtir, "gtri": gtri, "gtrr": gtrr, "eqir": eqir, "eqri": eqri, "eqrr": eqrr}

func (vm *AocVM) Parse(str string) (ret Instruction) {
	strs := strings.Split(str, " ")
	ret.op = opnames[strs[0]]
	ret.a, _ = strconv.ParseInt(strs[1], 10, 0)
	ret.b, _ = strconv.ParseInt(strs[2], 10, 0)
	ret.c, _ = strconv.ParseInt(strs[3], 10, 0)
	return
}

func (vm *AocVM) Copy() (ret AocVM) {
	ret.regs = make([]int64, len(vm.regs))
	copy(ret.regs, vm.regs)
	ret.ip, ret.rip = vm.ip, vm.rip
	return
}

func (vm *AocVM) Equals(other *AocVM) bool {
	for j := range vm.regs {
		if vm.regs[j] != other.regs[j] {
			return false
		}
	}
	return vm.ip == other.ip && vm.rip == other.rip
}

func (vm *AocVM) ExecOne(insn Instruction) {
	if vm.rip >= 0 {
		vm.regs[vm.rip] = int64(vm.ip)
	}
	switch insn.op {
	case addr:
		vm.regs[insn.c] = vm.regs[insn.a] + vm.regs[insn.b]
	case addi:
		vm.regs[insn.c] = vm.regs[insn.a] + insn.b
	case mulr:
		vm.regs[insn.c] = vm.regs[insn.a] * vm.regs[insn.b]
	case muli:
		vm.regs[insn.c] = vm.regs[insn.a] * insn.b
	case banr:
		vm.regs[insn.c] = vm.regs[insn.a] & vm.regs[insn.b]
	case bani:
		vm.regs[insn.c] = vm.regs[insn.a] & insn.b
	case borr:
		vm.regs[insn.c] = vm.regs[insn.a] | vm.regs[insn.b]
	case bori:
		vm.regs[insn.c] = vm.regs[insn.a] | insn.b
	case setr:
		vm.regs[insn.c] = vm.regs[insn.a]
	case seti:
		vm.regs[insn.c] = insn.a
	case gtir:
		if insn.a > vm.regs[insn.b] {
			vm.regs[insn.c] = 1
		} else {
			vm.regs[insn.c] = 0
		}
	case gtri:
		if vm.regs[insn.a] > insn.b {
			vm.regs[insn.c] = 1
		} else {
			vm.regs[insn.c] = 0
		}
	case gtrr:
		if vm.regs[insn.a] > vm.regs[insn.b] {
			vm.regs[insn.c] = 1
		} else {
			vm.regs[insn.c] = 0
		}
	case eqir:
		if insn.a == vm.regs[insn.b] {
			vm.regs[insn.c] = 1
		} else {
			vm.regs[insn.c] = 0
		}
	case eqri:
		if vm.regs[insn.a] == insn.b {
			vm.regs[insn.c] = 1
		} else {
			vm.regs[insn.c] = 0
		}
	case eqrr:
		if vm.regs[insn.a] == vm.regs[insn.b] {
			vm.regs[insn.c] = 1
		} else {
			vm.regs[insn.c] = 0
		}
	}
	if vm.rip >= 0 {
		vm.ip = int(vm.regs[vm.rip])
	}
}

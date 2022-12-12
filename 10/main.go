package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

func main() {
	contents := shared.ReadFileContents("input.txt")
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	vm := NewVirtualMachine()
	for _, line := range lines {
		vm.ParseAndQueueInstruction(line)
	}
	score := 0
	nextCycleCheck := 20
	crt := [6][40]bool{}
	for {
		if !vm.Tick() {
			break
		}
		if vm.CurrentCycle+1 == nextCycleCheck {
			score += (vm.CurrentCycle + 1) * vm.Register
			fmt.Printf("%v: %v\n", vm.CurrentCycle, vm.Register)
			nextCycleCheck += 40
		}
		drawCrtPixel(&crt, vm.CurrentCycle, vm.Register)

	}
	solution := score
	fmt.Printf("Solution: %v\n", solution)
	printCrt(crt)
}

func drawCrtPixel(crt *[6][40]bool, index int, registerValue int) {
	col := index % 40
	row := index / 40
	if col == registerValue || col-1 == registerValue || col+1 == registerValue {
		crt[row][col] = true
	}
}

func printCrt(crt [6][40]bool) {
	for _, row := range crt {
		for _, val := range row {
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

type VirtualMachine struct {
	Register           int
	CurrentCycle       int
	instructionQueue   shared.Queue[Instruction]
	currentInstruction *Instruction
}

func NewVirtualMachine() VirtualMachine {
	return VirtualMachine{
		instructionQueue:   shared.NewQueue[Instruction](),
		currentInstruction: nil,
		Register:           1,
		CurrentCycle:       0,
	}
}

func (vm *VirtualMachine) ParseAndQueueInstruction(line string) {
	s := strings.Split(line, " ")
	rawInstructionType := s[0]
	parameters := s[1:]

	instructionType := ParseInstructionType(rawInstructionType)
	vm.instructionQueue.Offer(&Instruction{
		_type:      instructionType,
		parameters: parameters,
		cyclesLeft: instructionType.requiredCycles,
	})
}

func (vm *VirtualMachine) evaluateInstruction() {
	switch vm.currentInstruction._type.name {
	case addx:
		vm.evaluateAddx(vm.currentInstruction.parameters)
		break
	case noop:
		break
	default:
		panic("Unknown instrction " + vm.currentInstruction._type.name)
	}
}

func (vm *VirtualMachine) evaluateAddx(parameters []string) {
	n, _ := strconv.Atoi(parameters[0])
	vm.Register += n
}

func (vm *VirtualMachine) Tick() bool {
	if vm.currentInstruction == nil {
		if instruction, ok := vm.instructionQueue.Poll(); ok {
			vm.currentInstruction = instruction
		} else {
			return false
		}
	}
	vm.currentInstruction.cyclesLeft--
	vm.CurrentCycle++
	if vm.currentInstruction.cyclesLeft == 0 {
		vm.evaluateInstruction()
		vm.currentInstruction = nil
	}
	return true
}

func (vm *VirtualMachine) RunUntilExit() {
	for {
		if !vm.Tick() {
			break
		}
	}
}

type Instruction struct {
	_type      InstructionType
	parameters []string
	cyclesLeft int
}

type InstructionType struct {
	name           InstructionTypeName
	requiredCycles int
}

type InstructionTypeName string

const (
	addx InstructionTypeName = "addx"
	noop                     = "noop"
)

func ParseInstructionType(s string) InstructionType {
	switch s {
	case "addx":
		return InstructionType{
			name:           addx,
			requiredCycles: 2,
		}
	case "noop":
		return InstructionType{
			name:           noop,
			requiredCycles: 1,
		}
	default:
		panic("Unknown instruction type " + s)
	}
}

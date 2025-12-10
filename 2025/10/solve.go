package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/draffensperger/golp"
	"github.com/pglaum/aoc-go/util"
)

type Machine struct {
	lights         uint16
	buttons        []uint16
	desiredJoltage []uint32

	// simulation
	lightsSeen  []uint16
	lightStates []uint16
	joltage     []uint32
	steps       uint32
	done        bool

	// dfs
	leastSteps int
}

func (m *Machine) parseLights(str string) {
	if str[0] != '[' || str[len(str)-1] != ']' {
		fmt.Printf("invalid lights string %s\n", str)
		return
	}
	str = str[1 : len(str)-1]
	m.lights = 0
	for i := range str {
		if str[i] == '#' {
			m.lights |= 1 << (i)
		}
	}
}

func (m *Machine) parseButtons(strs []string) {
	m.buttons = []uint16{}

	for _, str := range strs {
		if str[0] != '(' || str[len(str)-1] != ')' {
			fmt.Printf("invalid button string %s\n", str)
			return
		}

		var button uint16
		nums := strings.SplitSeq(str[1:len(str)-1], ",")
		for num := range nums {
			i := int(num[0] - '0')
			button |= 1 << i
		}
		m.buttons = append(m.buttons, button)
	}
}

func (m *Machine) parseJoltage(str string) {
	if str[0] != '{' || str[len(str)-1] != '}' {
		fmt.Printf("invalid joltage string %s\n", str)
		return
	}

	m.desiredJoltage = []uint32{}
	nums := strings.SplitSeq(str[1:len(str)-1], ",")
	for num := range nums {
		j, _ := strconv.Atoi(num)
		m.desiredJoltage = append(m.desiredJoltage, uint32(j))
	}
}

var sample bool

func main() {
	start := time.Now()
	flag.BoolVar(&sample, "sample", false, "use sample input")
	flag.Parse()

	filename := "input.txt"
	if sample {
		filename = "sample.txt"
	}
	lines := util.ReadInputLines(filename, true)

	machines := []Machine{}
	for _, line := range lines {
		fields := strings.Fields(line)
		machine := Machine{
			lightStates: []uint16{0},
		}
		machine.parseLights(fields[0])
		machine.parseButtons(fields[1 : len(fields)-1])
		machine.parseJoltage(fields[len(fields)-1])
		machines = append(machines, machine)
	}

	elapsedParse := time.Since(start)

	start1 := time.Now()
	part1(machines)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	part2(machines)
	elapsed2 := time.Since(start2)
	elapsedFull := time.Since(start)

	util.PrintElapsed(elapsedParse, elapsed1, elapsed2, elapsedFull)
}

func (m *Machine) runButtons() {
	m.steps += 1
	newStates := []uint16{}
	for _, light := range m.lightStates {
		for _, button := range m.buttons {
			l := light ^ button
			if l == m.lights {
				fmt.Printf("found target in %d steps\n", m.steps)
				m.done = true
				return
			}
			if slices.Contains(m.lightsSeen, l) {
				continue
			}
			newStates = append(newStates, l)
		}
	}
	m.lightStates = newStates
}

func part1(machines []Machine) {
	count := 0
	for _, machine := range machines {
		for !machine.done {
			machine.runButtons()
		}
		count += int(machine.steps)
	}
	println("Part 1:", count)
}

func (m *Machine) ReachJoltage() int {
	numButtons := len(m.buttons)
	numJoltages := len(m.desiredJoltage)

	lp := golp.NewLP(0, numButtons)
	lp.SetVerboseLevel(golp.NEUTRAL)

	objectiveCoeffs := make([]float64, numButtons)
	for i := range numButtons {
		objectiveCoeffs[i] = 1.0
	}
	lp.SetObjFn(objectiveCoeffs)

	for i := range numButtons {
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, 1000.0)
	}

	for i := range numJoltages {
		var entries []golp.Entry
		for j, btn := range m.buttons {
			if (btn & (1 << i)) != 0 {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		targetValue := float64(m.desiredJoltage[i])
		if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
			panic(err)
		}
	}

	status := lp.Solve()
	if status != golp.OPTIMAL {
		println("not optimal")
		return 0
	}

	solution := lp.Variables()
	totalPresses := 0
	for _, val := range solution {
		totalPresses += int(val + 0.5)
	}

	return totalPresses
}

func part2(machines []Machine) {
	count := 0
	for _, machine := range machines {
		count += machine.ReachJoltage()
	}
	println("Part 2:", count)
}

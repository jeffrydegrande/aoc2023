package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "A brief description of your command",
	Run:   day3Main,
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func gridSize(b [][]byte) (int, int) {
	return len(b[0]) - 1, len(b) - 2
}

type Part struct {
	Digits []Digit
	GearX  int
	GearY  int
}

type Digit struct {
	B byte
	X int
	Y int
}

func (p *Part) Number() int {
	var bytes []byte
	for _, d := range p.Digits {
		bytes = append(bytes, d.B)
	}

	i, err := strconv.Atoi(string(bytes))
	failIf(err)
	return i
}

func (p *Part) String() string {
	var bytes []byte

	for _, d := range p.Digits {
		bytes = append(bytes, d.B)
	}

	return string(bytes)
}

func (p *Part) MaybeGear(bytes [][]byte, x, y int) bool {
	if bytes[y][x] == '*' {
		p.GearX = x
		p.GearY = y
		return true
	}
	return false
}

func (p *Part) Add(b byte, x, y int) {
	p.Digits = append(p.Digits, Digit{B: b, X: x, Y: y})
}

func at(bytes [][]byte, x, y int) bool {
	return !(bytes[y][x] == '.' || isDigit(bytes[y][x]))
}

func Check(b [][]byte, part *Part) bool {
	// x, y := 139, 139
	x, y := gridSize(b)

	for _, d := range part.Digits {
		//left
		if at(b, max(0, d.X-1), d.Y) {
			return true
		}

		// top left
		if at(b, max(0, d.X-1), max(0, d.Y-1)) {
			return true
		}

		// top
		if at(b, d.X, max(0, d.Y-1)) {
			return true
		}

		if at(b, min(x, d.X+1), max(0, d.Y-1)) {
			return true
		}
		if at(b, min(x, d.X+1), d.Y) {
			return true
		}

		if at(b, min(x, d.X+1), min(y, d.Y+1)) {
			return true
		}

		if at(b, d.X, min(y, d.Y+1)) {
			return true
		}
		if at(b, max(0, d.X-1), min(y, d.Y+1)) {
			return true
		}
	}

	return false
}

func CheckGear(b [][]byte, part *Part) {
	// x, y := 139, 139
	x, y := gridSize(b)

	for _, d := range part.Digits {
		if part.MaybeGear(b, max(0, d.X-1), d.Y) {
			break
		}
		if part.MaybeGear(b, max(0, d.X-1), max(0, d.Y-1)) {
			break
		}
		if part.MaybeGear(b, d.X, max(0, d.Y-1)) {
			break
		}
		if part.MaybeGear(b, min(x, d.X+1), max(0, d.Y-1)) {
			break
		}
		if part.MaybeGear(b, min(x, d.X+1), d.Y) {
			break
		}
		if part.MaybeGear(b, min(x, d.X+1), min(y, d.Y+1)) {
			break
		}
		if part.MaybeGear(b, d.X, min(y, d.Y+1)) {
			break
		}
		if part.MaybeGear(b, max(0, d.X-1), min(y, d.Y+1)) {
			break
		}
	}
}

func day3Main(cmd *cobra.Command, args []string) {
	// day3MainPart1(cmd, args)
	day3MainPart2(cmd, args)
}

func day3MainPart1(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile("day3.txt")
	failIf(err)

	// split into grid
	b := bytes.Split(data, []byte("\n"))

	// maxY := len(b) - 1 // last newline
	// maxX := len(b[0])

	var parts []*Part
	var part *Part = nil

	for y, line := range b {
		if len(line) == 0 {
			continue
		}

		// go through the line, collect digits
		for x, c := range line {
			if isDigit(c) {
				if part == nil {
					part = &Part{}
				}
				part.Add(c, x, y)
			} else {
				if part != nil {
					parts = append(parts, part)
				}
				part = nil
			}
		}
	}

	s := 0
	for _, p := range parts {
		if Check(b, p) {
			s += p.Number()
		}
	}
	// fmt.Println(s)
}

func day3MainPart2(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile("day3.txt")
	failIf(err)

	// split into grid
	b := bytes.Split(data, []byte("\n"))

	var parts []*Part
	var part *Part = nil

	for y, line := range b {
		if len(line) == 0 {
			continue
		}

		// go through the line, collect digits
		for x, c := range line {
			if isDigit(c) {
				if part == nil {
					part = &Part{}
				}
				part.Add(c, x, y)
			} else {
				if part != nil {
					parts = append(parts, part)
				}
				part = nil
			}
		}
	}

	// check for gear
	// s := 0
	for _, p := range parts {
		CheckGear(b, p)
	}

	m := make(map[string][]int)

	for _, p := range parts {
		key := fmt.Sprintf("%d,%d", p.GearX, p.GearY)
		if key == "0,0" {
			continue
		}
		m[key] = append(m[key], p.Number())
	}

	sum := 0
	for _, v := range m {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	fmt.Println(sum)
}

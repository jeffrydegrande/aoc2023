/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Day 2 of advent of code",
	Run:   day2Main,
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}

func part1(line string) int {
	parts := strings.Split(line, ":")

	/// extract game id
	gameID := strings.Split(parts[0], " ")
	id, err := strconv.Atoi(gameID[1])
	failIf(err)

	// extract the cubes
	r := regexp.MustCompile(`\d+ (blue|green|red)`)
	matches := r.FindAllString(parts[1], -1)

	for _, m := range matches {
		mm := strings.Split(m, " ")
		count, err := strconv.Atoi(mm[0])
		failIf(err)
		color := mm[1]

		switch color {
		case "red":
			if count > 12 {
				return 0
			}
		case "blue":
			if count > 14 {
				return 0
			}
		case "green":
			if count > 13 {
				return 0
			}
		}
	}
	return id
}

func part2(line string) int {
	parts := strings.Split(line, ":")

	// extract the cubes
	r := regexp.MustCompile(`\d+ (blue|green|red)`)
	matches := r.FindAllString(parts[1], -1)

	red, green, blue := 0, 0, 0

	for _, m := range matches {
		mm := strings.Split(m, " ")
		count, err := strconv.Atoi(mm[0])
		failIf(err)
		color := mm[1]

		switch color {
		case "red":
			red = max(red, count)
		case "blue":
			blue = max(blue, count)
		case "green":
			green = max(green, count)
		}
	}
	return red * green * blue
}

func day2Main(cmd *cobra.Command, args []string) {
	day2MainPart1(cmd, args)
	day2MainPart2(cmd, args)
}

func day2MainPart1(cmd *cobra.Command, args []string) {
	f, err := os.Open("day2.txt")
	failIf(err)

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += part1(line)
	}
	fmt.Println("Part1:", sum)
	failIf(scanner.Err())
}

/// solution to step 2

func day2MainPart2(cmd *cobra.Command, args []string) {
	f, err := os.Open("day2.txt")
	failIf(err)

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += part2(line)
	}
	fmt.Println("Part2:", sum)
	failIf(scanner.Err())
}

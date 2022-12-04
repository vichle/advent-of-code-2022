package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

type Range struct {
	start int
	end   int
}

func (a Range) completelyOverlaps(b Range) bool {
	return a.start <= b.start && a.end >= b.end
}

func (a Range) overlapsAtBeginning(b Range) bool {
	return a.start <= b.start && a.end >= b.start
}

type RangePair struct {
	a Range
	b Range
}

func (rp RangePair) completelyOverlaps() bool {
	return rp.a.completelyOverlaps(rp.b) || rp.b.completelyOverlaps(rp.a)
}

func (rp RangePair) overlaps() bool {
	return rp.a.overlapsAtBeginning(rp.b) || rp.b.overlapsAtBeginning(rp.a)
}

func main() {
	contents := shared.ReadFileContents("input.txt")
	pairs := parseContents(strings.TrimSpace(contents))

	completeOverlapCount := 0
	overlapCount := 0
	for _, pair := range pairs {
		if pair.completelyOverlaps() {
			completeOverlapCount += 1
		}
		if pair.overlaps() {
			overlapCount += 1
		}
	}
	fmt.Printf("Complete overlap: %v\nOverlaps: %v\n", completeOverlapCount, overlapCount)
}

func parseContents(contents string) []RangePair {
	rows := strings.Split(contents, "\n")
	pairs := make([]RangePair, 0)
	for _, row := range rows {
		ranges := strings.Split(row, ",")
		pairs = append(pairs, RangePair{
			a: parseRange(ranges[0]),
			b: parseRange(ranges[1]),
		})
	}

	return pairs
}

func parseRange(r string) Range {
	tmp := strings.Split(r, "-")
	start, _ := strconv.Atoi(tmp[0])
	end, _ := strconv.Atoi(tmp[1])
	return Range{
		start: start,
		end:   end,
	}
}

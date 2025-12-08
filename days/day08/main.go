package main

import (
	"aoc-2025/utils"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// +==================================================+
// |                 PARSE SECTION                    |
// +==================================================+

type coordinate struct {
	x int
	y int
	z int
}

func parseInput(input string) (ans []coordinate) {

	split_line := strings.SplitSeq(input, "\n")

	for line := range split_line {
		coord_xyz := strings.Split(line, ",")

		x, err := strconv.Atoi(coord_xyz[0])
		if err != nil {
			fmt.Println("Error: 1")
		}

		y, err := strconv.Atoi(coord_xyz[1])
		if err != nil {
			fmt.Println("Error: 2")
		}

		z, err := strconv.Atoi(coord_xyz[2])
		if err != nil {
			fmt.Println("Error: 3")
		}

		coord := coordinate{
			x: x,
			y: y,
			z: z,
		}

		ans = append(ans, coord)
	}

	return ans
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func (c coordinate) Dim(i int) int {
	switch i {
	case 0:
		return c.x
	case 1:
		return c.y
	default:
		return c.z
	}
}

func Dist2(a, b coordinate) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

type Node struct {
	c     coordinate
	Dim   int
	Left  *Node
	Right *Node
}

type KDTree struct {
	Root *Node
}

func Build(points []coordinate, depth int) *Node {
	if len(points) == 0 {
		return nil
	}

	dim := depth % 3

	sort.Slice(points, func(i, j int) bool {
		return points[i].Dim(dim) < points[j].Dim(dim)
	})

	median := len(points) / 2

	return &Node{
		c:     points[median],
		Dim:   dim,
		Left:  Build(points[:median], depth+1),
		Right: Build(points[median+1:], depth+1),
	}
}

func New(point []coordinate) *KDTree {
	return &KDTree{Root: Build(point, 0)}
}

func nearest(node *Node, target coordinate, best *Node, bestDist *int) *Node {
	if node == nil {
		return best
	}

	d := Dist2(target, node.c)
	if d < *bestDist {
		*bestDist = d
		best = node
	}

	var primary, secondary *Node

	if target.Dim(node.Dim) < node.c.Dim(node.Dim) {
		primary = node.Left
		secondary = node.Right
	} else {
		primary = node.Right
		secondary = node.Left
	}

	best = nearest(primary, target, best, bestDist)

	diff := target.Dim(node.Dim) - node.c.Dim(node.Dim)
	if diff*diff < *bestDist {
		best = nearest(secondary, target, best, bestDist)
	}

	return best
}

func (t *KDTree) Nearest(target coordinate) coordinate {
	bestDist := math.MaxInt
	best := nearest(t.Root, target, t.Root, &bestDist)
	return best.c
}

func Part1(input string) int {
	coordinates := parseInput(input)

	tree := New(coordinates)

	c := tree.Nearest(coordinate{})

	fmt.Println(coordinates)
	fmt.Println(c)

	return 1
}

func Part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 2
}

// +==================================================+
// |                  MAIN SECTION                    |
// +==================================================+

//go:embed input.txt
var input string

func init() {
	// for lines
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	logger := utils.NewPartLogger("day08")

	switch part {
	case 1:
		ans := Part1(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		logger.PrintPart(1, ans)
	case 2:
		ans := Part2(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		logger.PrintPart(2, ans)
	default:
		ans1 := Part1(input)
		ans2 := Part2(input)
		logger.PrintPart(1, ans1)
		logger.PrintPart(2, ans2)
	}
}

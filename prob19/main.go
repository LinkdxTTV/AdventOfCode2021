package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
	z int
}

type Scanner struct {
	num     int
	beacons []Point
}

func main() {
	now := time.Now()
	f, err := os.Open("./example")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	io := bufio.NewScanner(f)

	i := 0
	scanners := []Scanner{}
	scanner := Scanner{num: i, beacons: []Point{}}
	for io.Scan() {
		line := io.Text()
		if line == "" {
			scanners = append(scanners, scanner)
			continue
		}
		if strings.Contains(line, "---") {
			scanner = Scanner{num: i, beacons: []Point{}}
			i++
			continue
		}
		x, y, z := xyzSplitLine(line)
		scanner.beacons = append(scanner.beacons, Point{x, y, z})
	}
	// Grab the last one
	scanners = append(scanners, scanner)

	// Assume every scanner must have atleast 1 scanner in common with the list
	// Lest we have a rogue scanner ??
	// Map out scanner 0 as true grid / rotation
	beaconMap := map[[3]int]bool{}

	scanner0 := scanners[0]
	for _, beacon := range scanner0.beacons {
		point := [3]int{beacon.x, beacon.y, beacon.z}
		beaconMap[point] = true
	}

	// Keep an array of the scanners
	scannerPositions := []Point{}

	// Keep track of what scanners we have linked so far
	scannerMap := map[int]bool{0: true}
	for i := range scanners {
		if i == 0 {
			continue
		}
		scannerMap[i] = false
	}
	// Reset indexing and GO
	// We need to use a while loop essentially because scanners wont link in order
	i = 0
	for !allScannersLinked(scannerMap) {
		i++
		if i >= len(scannerMap) {
			i = 0
		}
		if scannerMap[i] {
			continue
		}

		stop := false
		for _, rotation := range scanners[i].returnEveryRotation() {
			// Translation
			// Assume this point matches some other point..
			for existingPoint := range beaconMap {
				for _, possibleInitial := range rotation.beacons {
					xDiff := possibleInitial.x - existingPoint[0]
					yDiff := possibleInitial.y - existingPoint[1]
					zDiff := possibleInitial.z - existingPoint[2]
					matchingPoints := 0
					missingPoints := 0
					for _, possiblePoint := range rotation.beacons {

						_, ok := beaconMap[[3]int{possiblePoint.x - xDiff, possiblePoint.y - yDiff, possiblePoint.z - zDiff}]
						if ok {
							matchingPoints++
						} else {
							missingPoints++
						}
						if missingPoints >= 15 {
							break
						}
						if matchingPoints >= 12 {
							scannerMap[i] = true
							fmt.Println("POGGERS linked scanner", i)
							// Add them to the true map
							for _, possiblePoint := range rotation.beacons {
								scannerPositions = append(scannerPositions, Point{xDiff, yDiff, zDiff})
								translatedPoint := [3]int{possiblePoint.x - xDiff, possiblePoint.y - yDiff, possiblePoint.z - zDiff}
								beaconMap[translatedPoint] = true
							}
							stop = true
							break
						}
					}
					if stop {
						break
					}
				}
				if stop {
					break
				}
			}
			if stop {
				break
			}
		}
	}

	fmt.Println("Part 1:", len(beaconMap))

	// Find manhattans
	max := 0
	for i, scannerPoint1 := range scannerPositions {
		for j, scannerPoint2 := range scannerPositions {
			if i == j {
				continue
			}
			dist := manhattanDistance(scannerPoint1, scannerPoint2)
			if dist > max {
				max = dist
			}
		}
	}
	fmt.Println("Part 2:", max)
	fmt.Println(time.Since(now))
}

func manhattanDistance(point1, point2 Point) int {
	return abs(point1.x-point2.x) + abs(point1.y-point2.y) + abs(point1.z-point2.z)
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func allScannersLinked(scannerMap map[int]bool) bool {
	for _, linked := range scannerMap {
		if !linked {
			return false
		}
	}
	return true
}

// Some of these rotations are illegal but I'm too lazy to fix it
func (s Scanner) returnEveryRotation() []Scanner {
	output := []Scanner{}

	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 1; k++ {
				xsign := 1
				ysign := 1
				zsign := 1
				if i%2 == 0 {
					xsign = -1
				}
				if j%2 == 0 {
					ysign = -1
				}
				if k%2 == 0 {
					zsign = -1
				}
				for reverse := 0; reverse <= 1; reverse++ {
					for n := 0; n <= 2; n++ {
						x := n     // 0 1 2
						y := n + 1 // 1 2 0
						z := n + 2 // 2 0 1
						if y > 2 {
							y = y - 3
						}
						if z > 2 {
							z = z - 3
						}
						newScanner := Scanner{num: s.num, beacons: []Point{}}
						for _, point := range s.beacons {
							points := []int{xsign * point.x, ysign * point.y, zsign * point.z}
							newPoint := Point{x: points[x], y: points[y], z: points[z]}
							if reverse == 1 {
								newPoint = Point{x: points[z], y: points[y], z: points[x]}
							}
							newScanner.beacons = append(newScanner.beacons, newPoint)
						}
						output = append(output, newScanner)
					}
				}
			}
		}
	}
	return output
}

func xyzSplitLine(line string) (int, int, int) {
	split := strings.Split(line, ",")
	x, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	z, err := strconv.ParseInt(split[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(x), int(y), int(z)
}

package main

import "fmt"

func main() {

	// Not gonna waste time parsing today

	example := targetArea{20, 30, -10, -5}
	input := targetArea{111, 161, -154, -101}

	fmt.Println(example, input)

	targetArea := input

	maxHeight := 0
	counter := 0
	for xdot := -1000; xdot < 1000; xdot++ {
		for ydot := -1000; ydot < 1000; ydot++ {
			fmt.Println(xdot, ydot)
			xvel, yvel := xdot, ydot
			tempHeight := 0
			point := point{0, 0}
			for {
				point.move(xvel, yvel)
				if point.y > tempHeight {
					tempHeight = point.y
				}
				if isPointInTarget(point, targetArea) {
					counter++
					if tempHeight > maxHeight {
						maxHeight = tempHeight
					}
					break
				}
				if isPointOvershoot(point, targetArea) {
					break
				}
				applyDragAndGravity(&xvel, &yvel)
			}
		}
	}
	fmt.Println(maxHeight)
	fmt.Println(counter)

}

type targetArea struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

type point struct {
	x int
	y int
}

func isPointInTarget(point point, targetArea targetArea) bool {
	if point.x < targetArea.xmin || point.x > targetArea.xmax || point.y < targetArea.ymin || point.y > targetArea.ymax {
		return false
	}
	return true
}

func isPointOvershoot(point point, targetArea targetArea) bool {
	if point.x > targetArea.xmax || point.y < targetArea.ymin {
		return true
	}
	return false
}

func applyDragAndGravity(xdot, ydot *int) {
	*ydot--
	if *xdot > 0 {
		*xdot--
	}
	if *xdot < 0 {
		*xdot++
	}
}

func (p *point) move(xdot, ydot int) {
	p.x += xdot
	p.y += ydot
}

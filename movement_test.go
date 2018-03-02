package main

import "testing"

func TestGetPointBasedOnDirection(t *testing.T) {
	directions := []int{1, 2, 3, 4}
	source_point := Point{X: 0, Y: 7}
	expected_points := []Point{
		Point{X: 0, Y: 6},
		Point{X: 1, Y: 7},
		Point{X: 0, Y: 8},
		Point{X: -1, Y: 7},
	}
	for point_index, point := range expected_points {
		output := GetNextPointBasedOnDirection(directions[point_index], source_point)
		if output != point {
			t.Errorf("Given point %v and direction %v, the expected point should be %v, not %v", source_point, directions[point_index], output, point)
		}
	}
}

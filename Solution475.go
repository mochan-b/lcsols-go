package lcsols

import (
	"sort"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func findHeater(house int, heaters []int) int {
	left := 0
	right := len(heaters) - 1

	for left <= right {
		mid := (left + right) / 2
		if heaters[mid] == house {
			return mid
		} else if heaters[mid] < house {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	var MaxVal = Max(houses[len(houses)-1], heaters[len(heaters)-1]) - Min(houses[0], heaters[0]) + 1

	radius := 0

	for i := 0; i < len(houses); i++ {
		house := houses[i]
		heaterIndex := findHeater(house, heaters)

		leftDist := MaxVal
		if heaterIndex != 0 {
			leftDist = house - heaters[heaterIndex-1]
		}

		rightDist := MaxVal
		if heaterIndex != len(heaters) {
			rightDist = heaters[heaterIndex] - house
		}

		radius = Max(radius, Min(leftDist, rightDist))
	}

	return radius
}

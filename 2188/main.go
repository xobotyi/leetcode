package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumFinishTime(
		[][]int{
			{3, 4},
			{84, 2},
			{63, 8},
			{72, 8},
			{82, 7},
			{83, 6},
			{23, 2},
			{77, 5},
			{51, 10},
			{28, 2},
			{47, 9},
			{8, 3},
			{48, 3},
			{56, 3},
			{8, 10},
			{66, 6},
			{92, 9},
			{44, 6},
			{23, 5},
			{5, 6},
			{86, 9},
			{13, 10},
			{91, 3},
			{2, 2},
			{8, 4},
			{67, 8},
			{63, 6},
			{52, 5},
			{42, 10},
			{3, 9},
			{66, 5},
			{35, 10},
			{63, 6},
			{65, 6},
			{22, 8},
			{40, 9},
			{43, 4},
			{73, 9},
			{81, 5},
			{32, 2},
			{30, 5},
			{80, 9},
			{50, 4},
			{35, 4},
			{52, 7},
			{11, 5},
			{7, 8},
			{68, 3},
			{54, 8},
			{49, 8},
		},
		90,
		87,
	))
}
func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	minSameTire := minTimeSameTire(&tires, changeTime)
	minTime := map[int]int{}

	return minFinishTime(minTime, minSameTire, changeTime, numLaps) - changeTime
}

func minFinishTime(minTime map[int]int, minSameTime map[int]int, penalty, laps int) int {
	if laps == 0 {
		return 0
	}

	res, ok := minTime[laps]
	if ok {
		return res
	}

	res = math.MaxInt

	for lap := 1; lap <= minInt(len(minSameTime), laps); lap++ {
		time := penalty + minSameTime[lap] + minFinishTime(minTime, minSameTime, penalty, laps-lap)

		if time < res {
			res = time
		}
	}

	minTime[laps] = res

	return res
}

// Calculates the fastest way to run number of laps without changing tires.
// Meaning that if tires A, B, C completes 4 laps in 4, 5, and 6 seconds
// respectively 4 will be stored for 4 laps as it is the minimum time of all.
func minTimeSameTire(tires *[][]int, penalty int) (res map[int]int) {
	res = make(map[int]int, len(*tires)) // presumably we're over-allocating here, but it is fine.

	for _, tire := range *tires {
		for lap := 1; lap <= lapsTillChange(tire[0], tire[1], penalty); lap++ {
			time := timeSameTire(tire[0], tire[1], lap)

			if min, ok := res[lap]; !ok || time < min {
				res[lap] = time
			}
		}
	}

	return
}

// Calculates time required to run certain number of laps on same tire.
func timeSameTire(time, degradation, laps int) (res int) {
	for lap := 1; lap <= laps; lap++ {
		res += getTireTime(time, degradation, lap)
	}

	return
}

const MaxLaps = 1000 // by the problem constrains

// At a certain point there is no sense to continue using same tire, as lap time
// is more than change penalty + 1st lap time.
func lapsTillChange(time, degradation, penalty int) int {
	changePoint := getTireTime(time, degradation, 1) + penalty

	// start from 2nd lap as first is obviously less that its time + penalty
	for lap := 2; lap <= MaxLaps; lap++ {
		if getTireTime(time, degradation, lap) > changePoint {
			return lap - 1
		}
	}

	return MaxLaps
}

func getTireTime(time, degradation, lap int) int {
	return time * powInt(degradation, lap-1)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// faster solution for integers that math.Pow
func powInt(n, m int) int {
	if m == 0 {
		return 1
	}

	res := n
	for i := 1; i < m; i++ {
		res *= n
	}

	return res
}

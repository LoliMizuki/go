// from:
// [origin] http://qandwhat.apps.runkite.com/i-failed-a-twitter-interview/
// [CN] http://blog.jobbole.com/50705/

// 11.04: re-design test case struct
//  	: 還是沒有很正確的運作
// 11.05: 再次重寫吧, 好好切分 ...

package main

import (
	"fmt"
)

var showStep bool = true

type TestCase struct {
	Data            []int
	CorrectCapacity int
}

func main() {
	testCases := []TestCase{
		TestCase{[]int{5, 5, 5}, 0},
		// TestCase{[]int{5, 1, 5}, 4},
		// TestCase{[]int{5, 4, 3, 2, 1, 2, 3, 4, 5}, 16},
		// TestCase{[]int{5, 1, 2, 3, 4, 7}, 10},
		// TestCase{[]int{1, 2, 3, 4, 5}, 0}, // waste case1
		// TestCase{[]int{5, 4, 3, 2, 1}, 0}, // waste case2
		// TestCase{[]int{5, 2, 4, 3, 1}, 0},
		// TestCase{[]int{2, 5, 1, 2, 3, 4, 7, 7, 6}, 10},
		// TestCase{[]int{2, 5, 1, 3, 1, 2, 1, 7, 7, 6}, 17},
		// TestCase{[]int{8, 5, 1, 5}, 4},
		// TestCase{[]int{8, 1, 5}, 4},
		// TestCase{[]int{7, 6, 1, 4, 6}, 7},
		// TestCase{[]int{6, 1, 4, 6, 7, 5, 1, 6, 4}, 13},
	}

	failIndexes, failValues := runTest(testCases)
	errorReport(failIndexes, failValues, testCases)
}

func runTest(testCases []TestCase) (failIndexes, failValues []int) {
	failIndexes = make([]int, 0, len(testCases))
	failValues = make([]int, 0, len(testCases))

	for index, tc := range testCases {
		capacity := findTotalCapacity(tc.Data)

		if capacity != tc.CorrectCapacity {
			fmt.Println(capacity, tc.CorrectCapacity)
			failIndexes = append(failIndexes, index)
			failValues = append(failValues, capacity)
		}
	}

	return
}

func errorReport(failIndexes, failValues []int, testCases []TestCase) {
	for _, failIndex := range failIndexes {
		fmt.Printf("Test fail at case #%d data=%v, correct=%d, fail=%d",
			failIndexes, testCases[failIndex].Data, testCases[failIndex].CorrectCapacity, failValues[failIndex])
	}
}

// OMG ... it's mass block ...
func findTotalCapacity(heights []int) int {
	leftIndex := 0
	rightIndex := 1
	modilfyHeight := 0

	localHeighest := 0

	localCapacityTemp := 0
	totalCapacity := 0

	limitedTimes := 100
	limitedCount := 0

	for {
		if limitedCount > limitedTimes {
			break
		}

		leftHeight := heights[leftIndex] + modilfyHeight
		rightHeight := heights[rightIndex]

		if showStep {
			fmt.Printf("compare #%d(h:%d) and #%d(h:%d) | ", leftIndex, leftHeight, rightIndex, rightHeight)
		}

		if rightHeight > localHeighest {
			localHeighest = rightHeight
		}

		capacity := leftHeight - rightHeight

		// + capacity or + to total
		if capacity > 0 {
			localCapacityTemp += capacity
			rightIndex = rightIndex + 1

			if showStep {
				fmt.Println("cap:", capacity, "locCap:", localCapacityTemp)
			}
		} else {
			totalCapacity += localCapacityTemp
			localCapacityTemp = 0

			if leftIndex == rightIndex-1 {
				leftIndex = rightIndex
			} else {
				leftIndex = rightIndex + 1
			}
			rightIndex = leftIndex + 1

			if showStep {
				fmt.Println("succee find capacity, loc:", localCapacityTemp, "to total", totalCapacity)
			}
		}

		if rightIndex >= len(heights) {
			if localCapacityTemp > 0 {
				if modilfyHeight == 0 { // 未執行過區域降低修正過, 將 height 降低, 返回 leftIndex + 1 處重新搜索,
					rightIndex = leftIndex + 1
					modilfyHeight = localHeighest - heights[rightIndex]

					if showStep {
						fmt.Println("modilfy Height left back to:", leftIndex,
							", set base height:", heights[rightIndex]+modilfyHeight)
					}
				} else { // 已執行過區域降低修正過
					rightIndex = rightIndex + 1
					leftIndex = rightIndex + 1
					modilfyHeight = 0

					if showStep {
						fmt.Println("have modilfied Height left back to:", leftIndex,
							", set base height:", heights[rightIndex]+modilfyHeight)
					}
				}

				localCapacityTemp = 0

			} else {
				break
			}
		}

		if leftIndex >= len(heights) {
			if showStep {
				fmt.Println("l to end:", leftIndex, "break")
			}

			break
		}

		limitedCount++
	}

	return totalCapacity
}

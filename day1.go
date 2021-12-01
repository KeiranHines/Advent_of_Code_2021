package main

// CountIncreasing counts the number of times the value n in the data array is greater than
// the preceeding value at n-1
// It returns the count.
func CountIncreasing(data []int) int {
	last := data[0]
	count := 0
	for _, num := range data {
		if num > last {
			count++
		}
		last = num
	}
	return count
}

// CountSlidingWindowIncrease counts the number of times the sum of a window size y in the data array is greater than
// the preceeding window sum
// It returns the count.
func CountSlidingWindowIncrease(data []int, window int) int {
	lastSum := 0
	count := 0
	for i := 0; i < window; i++ {
		lastSum += data[i]
	}
	for i := window; i < len(data); i++ {
		sum := lastSum - data[i-window] + data[i]
		if sum > lastSum {
			count++
		}
		lastSum = sum
	}
	return count
}

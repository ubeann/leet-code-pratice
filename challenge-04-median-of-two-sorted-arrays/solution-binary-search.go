package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// NOTE: This solution uses binary search, and has a time complexity of O(log(min(m,n))).
	//       That's because we are performing binary search on the shorter array (nums1),
	//       and the number of elements we need to search is reduced by half in each iteration.
	//       The space complexity is O(1), because we are not using any extra space.

	// Get length of both arrays
	x, y := len(nums1), len(nums2)

	// Ensure nums1 is the shorter array
	if x > y {
		nums1, nums2 = nums2, nums1
		x, y = y, x
	}

	// Perform binary search on shorter array (nums1)
	low, high := 0, x

	// Loop until low is less than or equal to high
	for low <= high {
		// Calculate the partition points for nums1 and nums2
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		// Get the maximum and minimum values on the left and right sides of the partitions
		maxLeftX, minRightX := getMaxMin(nums1, partitionX)
		maxLeftY, minRightY := getMaxMin(nums2, partitionY)

		// Check if the current partition points satisfy the conditions for the median
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// Calculate and return the median based on whether the total number of elements is even or odd
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// If maxLeftX is too large, adjust the partitionX to move towards the left side
			high = partitionX - 1
		} else {
			// If maxLeftY is too large, adjust the partitionX to move towards the right side
			low = partitionX + 1
		}
	}

	// We should never reach here
	return 0.0
}

// `getMaxMin` is helper function to find maximum of left partition and minimum of right partition
func getMaxMin(nums []int, partition int) (maxLeft, minRight int) {
	// Check partition is 0, it means nothing is there on left side.
	if partition == 0 {
		// Use -INF for maxLeft, equivalent to math.MinInt
		maxLeft = ^int(^uint(0) >> 1)
	} else {
		// Use last element on left side as maxLeft
		maxLeft = nums[partition-1]
	}

	// Check partition is length of input, it means nothing is there on right side.
	if partition == len(nums) {
		// Use +INF for minRight, equivalent to math.MaxInt
		minRight = int(^uint(0) >> 1)
	} else {
		// Use first element on right side as minRight
		minRight = nums[partition]
	}

	// Return maxLeft and minRight
	return maxLeft, minRight
}

// `max` is helper function to find maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// `min` is helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

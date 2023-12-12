package main

func twoSumTraditional(nums []int, target int) []int {
	// NOTE: The time complexity of this solution is O(n^2) and the space complexity is O(1)
	//       where n is the length of the array. This solution uses a nested loop to check if the sum of two values
	//       in the array is equal to the target. If it is, the index of the two values is returned.
	//       If it isn't, the next two values are checked. This continues until the sum of two values is equal to the target.

	// Get the length of the array
	n := len(nums)

	// Loop through the array
	for i := 0; i < n-1; i++ {

		// Loop through the array starting from the next value of the current value
		for j := i + 1; j < n; j++ {

			// Check if the sum of the two values is equal to the target
			if nums[i]+nums[j] == target {

				// Return the index of the two values
				return []int{i, j}
			}
		}
	}

	// Return nil if no solution found
	return nil
}

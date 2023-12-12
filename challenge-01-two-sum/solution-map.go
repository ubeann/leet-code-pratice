package main

func twoSum(nums []int, target int) []int {
	// NOTE: The time complexity of this solution is O(n) and the space complexity is O(n)
	//       where n is the length of the array. This solution uses a map to store the value and index of the array.
	//       The map is used to check if the difference between the target and the current value exists in the array.
	//       If it does, the index of the current value and the index of the difference is returned.
	//       If it doesn't, the value and index of the array is stored in the map.

	// Create a map to store the value and index of the array
	m := make(map[int]int)

	// Loop through the array
	for i, v := range nums {

		// Check if the difference between the target and the current value
		if j, ok := m[target-v]; ok {

			// Return the index of the current value and the index of the difference
			return []int{j, i}
		}

		// Store the value and index of the array
		m[v] = i
	}

	// Return nil if no solution found
	return nil
}

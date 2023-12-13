package main

func lengthOfLongestSubstring(s string) int {
	// NOTE: The time complexity is O(n) where n is the length of the string,
	//       which same as the solution using a hashmap.
	//       This solution is faster than the solution using a hashmap.
	//       Instead of using a map to store the index of each character,
	//       we can use an array of size 128 to represent ASCII characters.
	//       This reduces the memory overhead associated with using a map.

	// Set max and start to 0
	max, start := 0, 0

	// Create an array to store the last index of each character
	var lastIdx [128]int

	// Initialize all elements of lastIdx to -1 (indicating unseen characters)
	for i := range lastIdx {
		lastIdx[i] = -1
	}

	// Loop through the string
	for i, char := range s {
		// Check if the character has been seen and its index is greater than or equal to start
		if lastIdx[char] >= start {
			start = lastIdx[char] + 1
		}

		// Update the index of the character
		lastIdx[char] = i

		// Update max if needed
		if i-start+1 > max {
			max = i - start + 1
		}
	}

	// Return max
	return max
}

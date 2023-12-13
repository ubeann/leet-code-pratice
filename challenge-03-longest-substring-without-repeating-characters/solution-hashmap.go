package main

func lengthOfLongestSubstringMap(s string) int {
	// Set max, start, and end to 0
	max, start, end := 0, 0, 0

	// Create a map to store the index of each character
	var m map[byte]int = make(map[byte]int)

	// Loop through the string
	for i := 0; i < len(s); i++ {

		// Check the character is in the map
		if _, ok := m[s[i]]; ok {

			// Check the index of the character is greater than or equal to start
			if m[s[i]] >= start {

				// Set start to the index of the character + 1
				start = m[s[i]] + 1
			}
		}

		// Set the index of the character to i
		m[s[i]] = i

		// Set end to i
		end = i

		// Check if end-start+1 is greater than max
		if end-start+1 > max {

			// Set max to new value
			max = end - start + 1
		}
	}

	// Return max
	return max
}

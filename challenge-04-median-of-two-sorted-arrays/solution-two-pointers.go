package main

func findMedianSortedArraysPointers(nums1 []int, nums2 []int) float64 {
	// NOTE: This solution is not as efficient as the binary search solution,
	//       but it is easier to understand. This algorithm has a time complexity
	//       of O(m+n), where m and n are the lengths of the two arrays (nums1 and nums2).
	//       That's because each element from both arrays is visited exactly once during the merging process.
	//       It's possible because the two arrays are already sorted.

	// Create a merged array, and merge the two arrays into it
	var merged []int

	// Initialize two pointers, one for each array
	var i, j int

	// Loop until we reach the end of either array
	for i < len(nums1) && j < len(nums2) {

		// Compare the two elements at the current pointers,
		// and append the smaller one to the merged array
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	// If there are remaining elements in `nums1`, append them to the `merged` array.
	for i < len(nums1) {
		merged = append(merged, nums1[i])
		i++
	}

	// If there are remaining elements in `nums2`, append them to the `merged` array.
	for j < len(nums2) {
		merged = append(merged, nums2[j])
		j++
	}

	// Return the median of the merged array
	if len(merged)%2 == 0 {
		return float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / 2
	} else {
		return float64(merged[len(merged)/2])
	}
}

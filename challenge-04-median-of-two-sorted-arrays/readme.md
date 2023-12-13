# Median of Two Sorted Arrays

![Difficulty: Hard](https://img.shields.io/badge/Difficulty-Hard-red)
![Likes: 26.9K](https://img.shields.io/badge/Likes-26.9K-blue)
![Dislikes: 3K](https://img.shields.io/badge/Dislikes-3K-red)

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

**Follow up:** The overall run time complexity should be `O(log (m+n))`.

## Example 1

> **Input:** nums1 = [1,3], nums2 = [2]
>
> **Output:** 2.00000
>
> **Explanation:** merged array = [1,2,3] and median is 2.

## Example 2

> **Input:** nums1 = [1,2], nums2 = [3,4]
>
> **Output:** 2.50000
>
> **Explanation:** merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

## Constraints

* `nums1.length == m`
* `nums2.length == n`
* $0 \leq m \leq 1000$
* $0 \leq n \leq 1000$
* $1 \leq m + n \leq 2000$
* $-10^6 \leq \text{nums1}[i], \text{nums2}[i] \leq 10^6$

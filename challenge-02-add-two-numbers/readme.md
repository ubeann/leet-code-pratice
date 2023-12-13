# Add Two Numbers

![Difficulty: Medium](https://img.shields.io/badge/Difficulty-Medium-yellow)
![Likes: 29.2K](https://img.shields.io/badge/Likes-29.2K-blue)
![Dislikes: 5.7K](https://img.shields.io/badge/Dislikes-5.7K-red)

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return *the sum as a linked list*.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

## Example 1

[![Example 1](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)](https://leetcode.com/problems/add-two-numbers/)

> **Input:** l1 = [2,4,3], l2 = [5,6,4]
>
> **Output:** [7,0,8]
>
> **Explanation:** 342 + 465 = 807.

## Example 2

> **Input:** l1 = [0], l2 = [0]
>
> **Output:** [0]
>
> **Explanation:** 0 + 0 = 0.

## Example 3

> **Input:** l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
>
> **Output:** [8,9,9,9,0,0,0,1]
>
> **Explanation:** 9999999 + 9999 = 10009998.

## Constraints

* The number of nodes in each linked list is in the range $[1, 100]$.
* $0 \leq \text{Node.val} \leq 9$
* It is guaranteed that the list represents a number that does not have leading zeros.

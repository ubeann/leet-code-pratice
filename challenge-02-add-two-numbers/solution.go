package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented by linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// NOTE: This solution is not the most efficient, but it is the most readable and easy to understand.
	//       The time complexity of this solution is O(n) and the space complexity is O(n) where n is the length of the lists.
	//       This solution loops through the lists and adds the values of the current nodes.
	//       If the sum is greater than 9, the 1 is carried to the next node.
	//       If there are more nodes in the lists, a new node is created.
	//       The result list is returned.

	// Create a result list with a dummy node
	result := &ListNode{}

	// Create a pointer to the current node in the result list
	dummy := result

	// Loop until both lists are empty
	for l1 != nil || l2 != nil {
		// Get the values of the current nodes in the lists
		if l1 != nil {
			dummy.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			dummy.Val += l2.Val
			l2 = l2.Next
		}

		// Check if the sum is greater than 9
		if dummy.Val > 9 {

			// Carry the 1 to the next node
			dummy.Val -= 10
			dummy.Next = &ListNode{Val: 1}

		} else if l1 != nil || l2 != nil {
			// Check if there are more nodes in the lists,
			// if so, create a new node
			dummy.Next = &ListNode{}
		}

		// Move the pointer to the next node
		dummy = dummy.Next
	}

	// Return the result list
	return result
}

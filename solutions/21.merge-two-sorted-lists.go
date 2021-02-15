/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (55.93%)
 * Total Accepted:    1.3M
 * Total Submissions: 2.3M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a sorted list. The list
 * should be made by splicing together the nodes of the first two lists.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: l1 = [1,2,4], l2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: l1 = [], l2 = []
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: l1 = [], l2 = [0]
 * Output: [0]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both l1 and l2 are sorted in non-decreasing order.
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// edge cases
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// initialise node and head
	var node, head *ListNode
	if l1.Val < l2.Val {
		node = l1
		head = l1
		l1 = l1.Next
	} else {
		node = l2
		head = l2
		l2 = l2.Next
	}

	// while loop 
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	// end result
	if l1 == nil {
		node.Next = l2
	}
	if l2 == nil {
		node.Next = l1
	}
	return head
}

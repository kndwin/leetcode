/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (39.83%)
 * Total Accepted:    1.3M
 * Total Submissions: 3.3M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: s = "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: s = "{[]}"
 * Output: true
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 * 
 * 
 */

type Stack struct {
	items []byte
}

func (s *Stack) push(b byte) {
	s.items = append(s.items, b)
}

func (s *Stack) pop() byte {
	l := len( s.items )
	var popped byte = byte( 0 )
	if l > 0 {
		popped = s.items[ l-1 ]
		s.items = s.items[:l-1]
	}
	return popped
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func ( s *Stack ) len() int {
	return len( s.items )
}

func ( s *Stack ) last() byte {
	l := len( s.items )
	var ans byte = byte(0)
	if l > 0 {
		ans = s.items[ l - 1 ]
	}
	return ans
}

func isValid(s string) bool {

	var stack Stack

	for i:=0; i<len(s); i++ {
		if s[i]==')' && stack.last()=='(' || 
			 s[i]=='}' && stack.last()=='{' || 
			 s[i]==']' && stack.last()=='['  {
			stack.pop()	
		} else {
			stack.push( s[i] )
		}
	}


	return stack.isEmpty()
}


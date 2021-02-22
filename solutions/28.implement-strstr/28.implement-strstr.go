/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (35.15%)
 * Total Accepted:    824.4K
 * Total Submissions: 2.3M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 * 
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 * 
 * Clarification:
 * 
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 * 
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 * 
 * 
 * Example 1:
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 * Example 2:
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 * Example 3:
 * Input: haystack = "", needle = ""
 * Output: 0
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= haystack.length, needle.length <= 5 * 10^4
 * haystack and needle consist of only lower-case English characters.
 * 
 * 
 */
func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	} 

	var a int = -1

	for i:=0; i<=len(haystack)-len(needle); i++ { 
		if needle[0] == haystack[i] { 
			isMatch := true
			for j:=0; j<len(needle); j++ { 
				if needle[j] != haystack[i+j] { 
					isMatch = false
				}
			}

			if isMatch == true { 
				a = i
				break
			}

		}
	}

	return a
}

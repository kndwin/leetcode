/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (36.08%)
 * Total Accepted:    943.6K
 * Total Submissions: 2.6M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 * 
 * If there is no common prefix, return an empty string "".
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] consists of only lower-case English letters.
 * 
 * 
 */

func longestCommonPrefix(strs []string) string {
	// edge cases of 0
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// expected cases
	var ans string = ""
	var shortStr string = strs[0]

	// find the shortest
	for i:=0; i<len(strs); i++ {
		if len( shortStr ) > len(strs[i]) {
			shortStr = strs[i]
		}
	}
	ans = shortStr

	// compare all combination of shortest word to other words
	for i:=0; i<len(strs); i++ {
		var hasPrefix bool = false
		for j:=len(ans); j>0; j-- {
			if shortStr == strs[i] {
				hasPrefix = true
				break
			}
			if strings.HasPrefix(strs[i], ans[0:j] ) == true {
				ans = ans[0:j]
				hasPrefix = true
				break
			}
		}

		if hasPrefix == false {
			return ""
		}

	}
	return ans
}

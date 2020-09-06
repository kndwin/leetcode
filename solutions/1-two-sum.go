
// O(n) + Maps
func twoSum(nums []int, target int) []int {
    // initialize map and compliment
    m := make(map[int]int)
    comp := int(0);
    // search thru array 
    for i := int(0); i<int(len(nums)); i++ {
        // if match (return pair)
        comp = target - nums[i]
        if index, isExist := m[comp]; isExist == true {
            return []int{index, i}
        }
        // else add to map
        m[nums[i]] = i
    }
    return []int{0};
}

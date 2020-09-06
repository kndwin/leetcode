// O(n^2) + Immutable
function twoSum(nums: number[], target: number): number[] {
    let ans: number[] = []
    for (let i=0; i<nums.length; i++) {
        const tmp:number[] =[...nums]
        delete tmp[i]
        for (let j=0; j<tmp.length; j++) {
            if (tmp[j]+nums[i] == target) {
                ans.push(j, i) 
                return ans
            }
        }
    }   
    return ans
};

// O(n) + Maps
function twoSum(nums: number[], target: number): number[] {
    // initialise tables
    let map = new Map()
    let comp : number
    // loop through array
    for (let i=0; i<nums.length; i++) {
        // if comp is found, return index pair
        comp = target - nums[i]
        if (map.get(comp) != undefined) {
            return [i, map.get(comp)]
        }
        // else add to map
        map.set(nums[i], i)
    }
    return [0]
};

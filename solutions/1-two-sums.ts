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

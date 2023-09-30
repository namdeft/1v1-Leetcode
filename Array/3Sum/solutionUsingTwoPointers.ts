function threeSum(nums: number[]): number[][] {
  nums.sort((a, b) => a - b);
  const result: number[][] = [];

  nums.forEach((value, i) => {
      if(i > 0 && nums[i] === nums[i - 1]) return;
      let left = i + 1;
      let right = nums.length - 1;
      let sum: number = value + nums[left] + nums[right];

      while(left < right) {
          if(sum < 0) left++;
          if(sum > 0) right--;
          if(sum === 0) {
              result.push([value, nums[left], nums[right]])
              left++;
              right--;
          }

      }
  })

  return result;
};

// Input: nums = [-1,0,1,2,-1,-4] --> [-1,-1,0,1,2,4]
// Output: [[-1,-1,2],[-1,0,1]]
class Solution {
   public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> numsMap;
        int size = nums.size();

        for (int i = 0; i < size; i++) {
            numsMap[nums[i]] = i;
        }

        for (int i = 0; i < size; i++) {
            int numNeeded = target - nums[i];

            if (numsMap.count(numNeeded) && numsMap[numNeeded] != i) {
                return {i, numsMap[numNeeded]};
            }
        }
        return {};
    }
};
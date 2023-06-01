class Solution {
   public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        vector<vector<int>> results;
        int size = nums.size();

        for (int i = 0; i < size - 1; i++) {
            int left = i + 1;
            int right = size - 1;

            while (left < right) {
                if (nums[i] + nums[left] + nums[right] < 0) {
                    left++;
                } else if (nums[i] + nums[left] + nums[right] > 0) {
                    right--;
                } else {
                    results.push_back({nums[i], nums[left], nums[right]});
                    int leftIndexTemp = left;
                    int rightIndexTemp = right;

                    while (left < right && nums[left] == nums[leftIndexTemp]) {
                        left++;
                    }

                    while (left < right &&
                           nums[right] == nums[rightIndexTemp]) {
                        right--;
                    }
                }
            }

            while (i < size - 1 && nums[i] == nums[i + 1]) {
                i++;
            }
        }

        return results;
    }
};

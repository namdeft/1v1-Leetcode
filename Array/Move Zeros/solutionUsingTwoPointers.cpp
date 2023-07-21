class Solution {
   public:
    void moveZeroes(vector<int>& nums) {
        int left = 0;
        int right = 0;
        int size = nums.size();

        while (right < size) {
            if (nums[right] == 0) {
                right++;
            }
            left++;
        }
    }
};
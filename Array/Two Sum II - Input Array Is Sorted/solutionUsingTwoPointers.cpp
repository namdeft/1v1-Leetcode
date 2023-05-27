class Solution {
   public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        int left = 0;
        int right = numbers.size() - 1;
        vector<int> results;

        while (left < right) {
            int sum = numbers[left] + numbers[right];

            if (sum < target) {
                left++;
            }

            if (sum > target) {
                right--;
            }

            if (sum == target) {
                break;
            }
        }

        results.push_back(left + 1);
        results.push_back(right + 1);

        return results;
    }
};
class Solution {
    public List<String> summaryRanges(int[] nums) {
        List<String> results = new ArrayList<>();
        String str;
        if (nums.length == 0) {
            return results;
        }

        int l = 0;
        for (int r = 0; r < nums.length; r++) {
            if (r == nums.length - 1 || nums[r + 1] != nums[r] + 1) {
                if (l == r) {
                    results.add(Integer.toString(nums[r]));
                } else {
                    str = Integer.toString(nums[l]) + "->" + Integer.toString(nums[r]);
                    results.add(str);
                }

                l = r + 1;
            }
        }

        return results;
    }
}
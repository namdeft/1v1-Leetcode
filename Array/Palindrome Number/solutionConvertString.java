class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        String strNum = String.valueOf(x);
        int l = 0;
        int r = strNum.length() - 1;
        while (l < r) {
            if (strNum.charAt(l) != strNum.charAt(r)) {
                return false;
            }

            l++;
            r--;
        }

        return true;
    }
}
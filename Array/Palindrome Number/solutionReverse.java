class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        int numReverse = 0;
        int numClone = x;
        while (numClone > 0) {
            int numTemp = 0;
            numTemp = numClone % 10;
            numReverse = numReverse * 10 + numTemp;
            numClone = numClone / 10;
        }

        if (numReverse != x) {
            return false;
        }

        return true;
    }
}
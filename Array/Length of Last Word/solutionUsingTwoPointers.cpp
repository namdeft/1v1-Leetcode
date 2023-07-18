class Solution {
   public:
    int lengthOfLastWord(string s) {
        int flag;
        int count = 0;

        for (int i = s.length() - 1; i >= 0; i--) {
            if (flag && s[i] == ' ') break;

            if (s[i] != ' ') {
                flag = 1;
                count++;
            }
        }

        return count;
    }
};
class Solution {
   public:
    bool isAnagram(string s, string t) {
        // CHeck the length of two strings
        if (s.length() != t.length()) {
            return false;
        }

        // Put single letter into vector as an index ('a' = 0, 'c' = 2)
        vector<int> freq(26, 0);
        for (int i = 0; i < s.length(); i++) {
            freq[s[i] - 'a']++;
            freq[t[i] - 'a']--;
        }

        for (int i = 0; i < freq.size(); i++) {
            if (freq[i] != 0) {
                return false;
            }
        }

        return true;
    }
};
class Solution {
   public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> results;
        unordered_map<string, vector<string>> map;
        int size = strs.size();

        for (int i = 0; i < size; i++) {
            string s = strs[i];
            sort(strs[i].begin(), strs[i].end());
            map[strs[i]].push_back(s);
        }

        for (auto i : map) {
            results.push_back(i.second);
        }

        return results;
    }
};
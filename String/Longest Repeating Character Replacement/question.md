# 424. Longest Repeating Character Replacement

# Medium

You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:

<pre>
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
</pre>

Example 2:

<pre>
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
</pre>

Constraints:

-   `1 <= s.length <= 105`
-   `s` consists of only uppercase English letters.
-   `0 <= k <= s.length`

<details>
<summary> Similar Questions </summary>

-   `Max Consecutive Ones III - Medium`
-   `Maximize the Confusion of an Exam - Medium`
-   `Minimum Number of Operations to Make Array Continuous - Hard`
-   `Longest Substring of One Repeating Character - Hard`

</details>

<details>
<summary> Related Topics </summary>

-   `Hash Table`
-   `String`
-   `Sliding Window`

</details>

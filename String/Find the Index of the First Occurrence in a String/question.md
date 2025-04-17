# 28. Find the Index of the First Occurrence in a String

# Easy

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

Example 1:

<pre>
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
</pre>

Example 2:

<pre>
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
</pre>

Constraints:

-   `1 <= haystack.length, needle.length <= 104`
-   `haystack` and `needle` consist of only lowercase English characters.

<details>
<summary> Similar Questions </summary>

-   `Repeated Substring Pattern - Easy`
-   `Shortest Palindrome - Hard`

</details>

<details>
<summary> Related Topics </summary>

-   `Two Pointers`
-   `String`
-   `String Matching`

</details>

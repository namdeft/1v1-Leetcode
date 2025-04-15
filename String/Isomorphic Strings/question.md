# 205. Isomorphic Strings

# Easy

Given two strings `s` and `t`, determine if they are isomorphic.

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:

<pre>
Input: s = "egg", t = "add"
Output: true
Explanation:
The strings s and t can be made identical by:

Mapping 'e' to 'a'.
Mapping 'g' to 'd'.
</pre>

Example 2:

<pre>
Input: s = "foo", t = "bar"
Output: false
Explanation:
The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.
</pre>

Example 3:

<pre>
Input: s = "paper", t = "title"
Output: true
</pre>

Constraints:

-   `1 <= s.length <= 5 * 104`
-   `t.length == s.length`
-   `s` and `t` consist of any valid ascii character.

<details>
<summary> Similar Questions </summary>

-   `Word Pattern - Easy`
-   `Find and Replace Pattern - Medium`

</details>

<details>
<summary> Related Topics </summary>

-   `Hash Table`
-   `String`

</details>

# 290. Word Pattern

# Easy

Given a `pattern` and a string `s`, find if `s` follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in `pattern` and a non-empty word in `s`. Specifically:

Each letter in `pattern` maps to exactly one unique word in `s`.
Each unique word in `s` maps to exactly one letter in `pattern`.
No two letters map to the same word, and no two words map to the same letter.

Example 1:

<pre>
Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Explanation:
The bijection can be established as:
`'a'` maps to `"dog"`.
`'b'` maps to `"cat"`.
</pre>

Example 2:

<pre>
Input: pattern = "abba", s = "dog cat cat fish"
Output: falseq
</pre>

Example 3:

<pre>
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false
</pre>

Constraints:

-   `1 <= pattern.length <= 300`
-   `pattern` contains only lower-case English letters.
-   `1 <= s.length <= 3000`
-   `s` contains only lowercase English letters and spaces `' '`.
-   `s` does not contain any leading or trailing spaces.
-   All the words in `s` are separated by a single space.

<details>
<summary> Similar Questions </summary>

-   `Isomorphic Strings - Easy`
-   `Word Pattern II - Medium`
-   `Find and Replace Pattern - Medium`

</details>

<details>
<summary> Related Topics </summary>

-   `Hash Table`
-   `String`

</details>

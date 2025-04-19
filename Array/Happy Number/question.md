# 202. Happy Number

# Easy

Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return `true` if n is a happy number, and `false` if not.


Example 1:

<pre>
Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

</pre>

Example 2:

<pre>
Input: n = 2
Output: false
</pre>

Constraints:

-   `1 <= n <= 231 - 1`

<details>
<summary> Similar Questions </summary>

-   `Linked List Cycle - Easy`
-   `Minimum Addition to Make Integer Beautiful - Easy`
-   `Smallest Value After Replacing With Sum of Prime Factors - Medium`

</details>

<details>
<summary> Related Topics </summary>

-   `Hash Table`
-   `Math`
-   `Two Pointers`

</details>

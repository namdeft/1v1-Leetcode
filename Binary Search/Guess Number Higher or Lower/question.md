# 374. Guess Number Higher or Lower

# Easy

We are playing the Guess Game. The game is as follows:

I pick a number from `1` to `n`. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API `int guess(int num)`, which returns three possible results:

- `-1`: Your guess is higher than the number I picked (i.e. `num > pick`).
- `1`: Your guess is lower than the number I picked (i.e. `num < pick`).
- `0`: your guess is equal to the number I picked (i.e. `num == pick`).
Return the number that I picked.


Example 1:

<pre>
Input: n = 10, pick = 6
Output: 6
</pre>

Example 2:

<pre>
Input: n = 1, pick = 1
Output: 1
</pre>

Example 3:

<pre>
Input: n = 2, pick = 1
Output: 1
</pre>

Constraints:

-   `1 <= n <= 231 - 1`
-   `1 <= pick <= n`

<details>
<summary> Similar Questions </summary>

-   `First Bad Version - Medium`
-   `Guess Number Higher or Lower II - Medium`
-   `Find K Closest Elements - Medium`

</details>

<details>
<summary> Related Topics </summary>

-   `Binary Seach`
-   `Interactive`

</details>

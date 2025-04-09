# 414. Third Maximum Number

# Easy

Given an integer array `nums`, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.


Example 1:

<pre>
Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.
</pre>

Example 2:

<pre>
Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.
</pre>

Example 3:

<pre>
Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.
</pre>

Constraints:

-   `1 <= nums.length <= 104`
-   `231 <= nums[i] <= 231 - 1`

Follow up: Can you find an O(n) solution?

<details>
<summary> Similar Questions </summary>

-   `Neither Minimum nor Maximum - Easy`
-   `Kth Largest Element in an Array - Medium`

</details>

<details>
<summary> Related Topics </summary>

-   `Array`
-   `Sorting`

</details>

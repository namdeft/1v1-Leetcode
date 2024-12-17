# 238. Product of Array Except Self

# Medium

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in `O(n)` time and without using the division operation.

Example 1:

<pre>
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
</pre>

Example 2:

<pre>
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
</pre>

Constraints:

- `2 <= nums.length <= 105`
- `30 <= nums[i] <= 30`
- The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

<details>
<summary> Similar Questions </summary>

-   `Trapping Rain Water - Hard`
-   `Maximum Product Subarray - Medium`
-   `Paint House II - Hard`

</details>

<details>
<summary> Related Topics </summary>

-   `Array`
-   `Prefix Sum`

</details>

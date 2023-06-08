# 143. Reorder List

# Medium

You are given the head of a singly linked-list. The list can be represented as:
`L0 → L1 → … → Ln - 1 → Ln`
Reorder the list to be on the following form:
`L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …`
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example 1:

<pre>
Input: head = [1,2,3,4]
Output: [1,4,2,3]
</pre>

Example 2:

<pre>
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
</pre>

Constraints:

-   `The number of nodes in the list is in the range [1, 5 * 104]`.
-   `1 <= Node.val <= 1000`

<details>
<summary> Similar Questions </summary>

-   `Delete the Middle Node of a Linked List - Medium`
-   `Take K of Each Character From Left and Right - Medium`

</details>

<details>
<summary> Related Topics </summary>

-   `Linked List`
-   `Two Pointers`
-   `Stack`
-   `Recursion`

</details>

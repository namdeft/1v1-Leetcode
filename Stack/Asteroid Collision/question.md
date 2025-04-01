# 735. Asteroid Collision

# Medium

We are given an array asteroids of integers representing `asteroids` in a row. The indices of the asteriod in the array represent their relative position in space.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

Example 1:

<pre>
Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
</pre>

Example 2:

<pre>
Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.
</pre>

Example 3:

<pre>
Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
</pre>

Constraints:

-   `2 <= asteroids.length <= 104`
-   `1000 <= asteroids[i] <= 1000`
-   `asteroids[i] != 0`

<details>
<summary> Similar Questions </summary>

-   `Can Place Flowers - Easy`
-   `Destroying Asteroids - Medium`
-   `Count Collisions on a Road - Medium`
-   `Robot Collisions - Hard`

</details>

<details>
<summary> Related Topics </summary>

-   `Array`
-   `Stack`
-   `Simulation`

</details>

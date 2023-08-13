/**
 Do not return anything, modify nums in-place instead.
 */

function rotate(nums: number[], k: number): void {
  k = k % nums.length;
  let l = 0;
  let r = nums.length - 1;

  reverse(nums, l, r);

  l = 0;
  r = k - 1;
  reverse(nums, l, r);

  l = k;
  r = nums.length - 1;
  reverse(nums, l, r);
};

function reverse(nums: number[], l: number, r: number) {
  let temp: any;

  while (l <= r) {
    temp = nums[l];
    nums[l] = nums[r];
    nums[r] = temp;

    l++;
    r--;
  }
}
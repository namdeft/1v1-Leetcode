/**
 Do not return anything, modify nums1 in-place instead.
 */
 function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    for (let n = nums2.length - 1; n >= 0; n--) {
      nums1[m + n] = nums2[n];
    }

    nums1.sort((a, b) => a - b);
 };
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var longestOnes = function(nums, k) {
    const len = nums.length;
    let maxLen = 0;
    let left = 0;
    let flipped = 0;
    let i = 0;

    while (flipped < k && i < len) {
        if (nums[i] === 0) flipped++;
        i++;
        maxLen++;
    }

    let curr = maxLen;
    for (i; i < len; i++) {
        if (nums[i] === 0) {
            while (nums[left] === 1 && left < len) {
                left++;
            }

            left++;
            curr = i - left;
        }

        curr++;
        if (curr > maxLen) maxLen = curr;
    }


    return maxLen;
};
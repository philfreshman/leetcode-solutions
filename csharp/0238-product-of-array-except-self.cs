public class Solution {
    public int[] ProductExceptSelf(int[] nums) {
        var res = new int[nums.Length];

        var prefix = 1;
        for (int i = 0; i < nums.Length; i++) {
            res[i] = prefix;
            prefix *= nums[i];
        }

        var postfix = 1;
        for (int i = nums.Length - 1; i >= 0; i--) {
            res[i] *= postfix;
            postfix *= nums[i];
        }

        return res;
    }
}
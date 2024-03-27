public partial class Solution {
    public int LongestConsecutive(int[] nums) {
        var set = new HashSet<int>(nums);
        var longest = 0;

        foreach (var n in nums) {
            if (!set.Contains(n - 1)) {
                var sequence = 1;
                while (set.Contains(n + sequence)) {
                    sequence++;
                }

                longest = Math.Max(longest, sequence);
            }
        }

        return longest;
    }
}

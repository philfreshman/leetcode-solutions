public partial class Solution {
    public bool ContainsDuplicate(int[] nums) {
        return nums.Distinct().Count() < nums.Length;
    }
}

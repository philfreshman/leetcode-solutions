public partial class Solution {
    public int[] NextGreaterElement(int[] nums1, int[] nums2) {
        var stack = new Stack<int>();
        var nextGreater = new Dictionary<int, int>();

        foreach (var num in nums2) {
            while (stack.Count > 0 && num > stack.Peek()) {
                nextGreater[stack.Pop()] = num;
            }
            stack.Push(num);
        }

        foreach (var num in stack) {
            nextGreater.Add(num, -1);
        }

        for (var i = 0; i < nums1.Length; i++) {
            nums1[i] = nextGreater[nums1[i]];
        }

        return nums1;
    }
}



public class Program {
    public static void Main() {
        var nums1 = new int[]{4, 1, 2};
        var nums2 = new int[]{1,3,4,2};

        var s = new Solution();
        var res = s.NextGreaterElement(nums1, nums2);
        foreach (var r in res) {
            Console.WriteLine(r);
        }
    }
}

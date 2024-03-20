public class Solution {
    public int[] TopKFrequent(int[] nums, int k) {
        var count = new Dictionary<int, int>();
        var freq = new List<int>[nums.Length + 1];

        for (int i = 0; i < freq.Length; i++) {
            freq[i] = new List<int>();
        }

        foreach (int n in nums) {
            if (count.ContainsKey(n)) {
                count[n]++;
            } else {
                count[n] = 1;
            }
        }

        foreach (var pair in count) {
            freq[pair.Value].Add(pair.Key);
        }

        var res = new List<int>();
        for (int i = freq.Length - 1; i > 0; i--) {
            foreach (int n in freq[i]) {
                res.Add(n);
                if (res.Count == k) {
                    return res.ToArray();
                }
            }
        }

        return res.ToArray();
    }
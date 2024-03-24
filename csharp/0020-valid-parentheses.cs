public class Solution {
    public bool IsValid(string s) {

        var stack = new List<char>();
        var map = new Dictionary<char, char>() {
            {')', '('},
            {'}', '{'},
            {']', '['},
        };

        foreach (var bracket in s) {
            var hasKey = map.ContainsKey(bracket);

            // if it's not a closing bracket, add it to the stack
            if (!hasKey) {
                stack.Add(bracket);
                continue;
            }

            // if opening pair exist, but not in the stack => exit
            if (stack.Count == 0) return false;

            // if last element is different that the pair
            if (stack[^1] != map[bracket]) return false;

            stack.RemoveAt(stack.Count - 1);
        }

        return stack.Count == 0;
    }
}
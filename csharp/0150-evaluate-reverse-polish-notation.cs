public partial class Solution {
    public int EvalRPN(string[] tokens) {

        var stack = new List<int>();

        foreach (var token in tokens) {
            if (token is not ("+" or "-" or "*" or "/")) {
                stack.Add(Convert.ToInt32(token));
                continue;
            }

            var num2 = stack.Last();
            stack.RemoveAt(stack.Count - 1);
            var num1 = stack.Last();
            stack.RemoveAt(stack.Count - 1);

            switch (token) {
                case "+":
                    stack.Add(num1 + num2);
                    break;
                case "-":
                    stack.Add(num1 - num2);
                    break;
                case "*":
                    stack.Add(num1 * num2);
                    break;
                case "/":
                    stack.Add(num1 / num2);
                    break;
            }
        }

        return stack[0];
    }
}

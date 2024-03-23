// time: O(1)
// space O(n)

public class MinStack {
    Stack<(int val, int minVal)> _stack;

    public MinStack() {
        _stack = new Stack<(int, int)>();
    }

    public void Push(int val) {
        var min = val;

        if (_stack.Count > 0 && GetMin() < val) {
            min = GetMin();
        }

        _stack.Push((val, min));
    }

    public void Pop() {
        _stack.Pop();
    }

    public int Top() {
        return _stack.Peek().val;
    }

    public int GetMin() {
        return _stack.Peek().minVal;

    }
}


public static class Program {

    private static void Main() {
        var x = new MinStack();
        x.Push(2);
    }

}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.Push(val);
 * obj.Pop();
 * int param_3 = obj.Top();
 * int param_4 = obj.GetMin();
 */
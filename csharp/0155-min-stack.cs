public class MinStack {
    Stack<(int val, int minVal)> _stack;

    public MinStack() => _stack = new Stack<(int, int)>();

    public void Push(int val) {
        var min = val;
        if (_stack.Count > 0 && GetMin() < val) {
            min = GetMin();
        }
        _stack.Push((val, min));
    }

    public void Pop() => _stack.Pop();
    public int Top() => _stack.Peek().val;
    public int GetMin() => _stack.Peek().minVal;
}

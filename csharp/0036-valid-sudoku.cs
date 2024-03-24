// time: O(1)
// space: O(1)

public class Solution {
    public bool IsValidSudoku(char[][] board) {
        var rows = new Dictionary<byte, bool>[9];
        var cols = new Dictionary<byte, bool>[9];
        var squares = new Dictionary<byte, bool>[9];

        for (var i = 0; i < 9; i++)
        {
            rows[i] = new Dictionary<byte, bool>();
            cols[i] = new Dictionary<byte, bool>();
            squares[i] = new Dictionary<byte, bool>();
        }

        for (var r = 0; r < 9; r++)
        {
            for (var c = 0; c < 9; c++)
            {
                if (board[r][c] == '.')
                {
                    continue;
                }
				var field = (byte)(board[r][c] - '0');
                var square = r / 3 * 3 + c / 3;
                if (rows[r].ContainsKey(field) || cols[c].ContainsKey(field) || squares[square].ContainsKey(field))
                {
                    return false;
                }
                rows[r][field] = true;
                cols[c][field] = true;
                squares[square][field] = true;
            }
        }

        return true;
    }
}

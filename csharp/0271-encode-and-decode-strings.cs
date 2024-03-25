public partial class Solution {
    public string Encode(List<string> strs) {
        var output = "";

        foreach (var word in strs) {
            output += $"{word.Length}#{word}";
        }

        return output;
    }

    public List<string> Decode(string s) {
        var output = new List<string>();
        var i = 0;

        while (i < s.Length) {
            var j = i;
            while (s[j] != '#') {
                j++;
            }
            var length = Int32.Parse(s.Substring(i, j - i));
            i = j + 1;
            j = i + length;
            output.Add(s.Substring(i, length));
            i = j;
        }

        return output;
    }
}
public partial class Solution {
    public bool IsIsomorphic(string s, string t) {
        if (s.Length != t.Length) {
            return false;
        }

        var hashSt = new Dictionary<char, char>();
        var hashTs = new Dictionary<char, char>();

        for (var i = 0; i < s.Length; i++) {
            var sChar = s[i];
            var tChar = t[i];
            if (hashSt.ContainsKey(sChar) && hashSt[sChar] != tChar) {
                return false;
            }
            if (hashTs.ContainsKey(tChar) && hashTs[tChar] != sChar) {
                return false;
            }
            hashSt[sChar] = tChar;
            hashTs[tChar] = sChar;
        }

        return true;
    }
}
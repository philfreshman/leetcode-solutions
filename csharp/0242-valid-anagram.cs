public partial class Solution {
    public bool IsAnagram(string s, string t) {
        var mapS = MapLetters(s);
        var mapT = MapLetters(t);

        if (mapT.Count != mapS.Count){
            return false;
        }

        foreach (var key in mapS.Keys) {
            if(!mapT.ContainsKey(key) || mapS[key] != mapT[key]){
                return false;
            }
        }

        return true;
    }

    private Dictionary<char, int> MapLetters(string s) {
        var map = new Dictionary<char, int>();
        foreach (var c in s) {
            if (map.ContainsKey(c)){
                map[c]++;
            } else {
                map[c] = 1;
            }
        }
        return map;
    }
}
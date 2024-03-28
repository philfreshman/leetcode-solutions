public partial class Solution {
    public bool IsPalindrome(string s) {
        int lp = 0;
        int rp = s.Length - 1;

        while (lp < rp) {
            char l = Char.ToLower(s[lp]);
            char r = Char.ToLower(s[rp]);

            if (!char.IsLetterOrDigit(l)){
                lp++;
                continue;
            }

            if (!char.IsLetterOrDigit(r)){
                rp--;
                continue;
            }

            if (l != r)
            {
                return false;
            }

            lp++;
            rp--;
        }

        return true;
    }
}
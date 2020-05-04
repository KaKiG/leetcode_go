#include <iostream>
#include <string>
#include <unordered_map>
using namespace std;

int max2(int, int);

class Solution
{
public:
    int lengthOfLongestSubstring(string s)
    {
        unordered_map<char, int> letterMap;
        int maxLen = 0, left = 0, right = 0;
        if (s.length() == 0)
        {
            return 0;
        }

        while (right < s.length())
        {
            auto it = letterMap.find(s[right]);
            if (it != letterMap.end() && it->second >= left)
            {
                left = it->second + 1;
            }
            maxLen = max2(maxLen, right - left + 1);
            letterMap[s[right]] = right;
            right++;
        }
        return maxLen;
    }
};

int max2(int a, int b)
{
    return a > b ? a : b;
}
int main()
{
    return 0;
}
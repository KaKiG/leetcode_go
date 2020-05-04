#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

// unordered_map is the hash map used in golang, basically equals map
class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        unordered_map<int, int> dict;
        for (int i = 0; i < nums.size(); i++)
        {
            auto res = dict.find(target - nums[i]);
            if (res != dict.end())
            {
                return vector<int>{res->second, i};
            }
            dict[nums[i]] = i;
        }
        return vector<int>();
    }
};

int main()
{
    vector<int> nums = {2, 7, 11, 15};
    vector<int> res;
    int target = 9;

    Solution sol;
    res = sol.twoSum(nums, target);
    for (int i = 0; i < res.size(); i++)
    {
        cout << res.at(i) << ' ';
    }
}
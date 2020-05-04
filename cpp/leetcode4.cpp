#include <iostream>
#include <vector>
using namespace std;

class Solution
{
public:
    double findMedianSortedArrays(vector<int> &nums1, vector<int> &nums2)
    {
        int count = 0, num1 = 0, num2 = 0, l1 = 0, l2 = 0;
        int index = (nums1.size() + nums2.size()) / 2 + 1;
        bool odd = (nums1.size() + nums2.size()) % 2 == 1;
        while (l1 < nums1.size() || l2 < nums2.size())
        {
            if ((l1 < nums1.size() && l2 < nums2.size() && nums1[l1] < nums2[l2]) || l2 == nums2.size())
            {
                num1 = nums1[l1];
                l1++;
            }
            else if ((l1 < nums1.size() && l2 < nums2.size() && nums1[l1] >= nums2[l2]) || l1 == nums1.size())
            {
                num1 = nums2[l2];
                l2++;
            }

            count++;

            if (odd && count == index)
            {
                return double(num1);
            }
            if (!odd && count == index - 1)
            {
                num2 = num1;
            }
            if (!odd && count == index)
            {
                return double(num1 + num2) / 2.0;
            }
        }
        return 0.0;
    }
};

int main()
{
    vector<int> nums1 = {1, 2};
    vector<int> nums2 = {3, 4};
    double res = 0.0;

    Solution sol;
    res = sol.findMedianSortedArrays(nums1, nums2);
    cout << res << ' ' << double(1) << endl;
    return 0;
}
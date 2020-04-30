#include <iostream>
using namespace std;

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution
{
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
    {
        ListNode res(0);
        ListNode *curr = &res;

        int num = 0;

        while (!l1 || !l2)
        {
            num /= 10;
            if (!l1)
            {
                num += l1->val;
                l1 = l1->next;
            }

            if (!l2)
            {
                num += l2->val;
                l2 = l2->next;
            }

            ListNode *tmp = new ListNode(num % 10);
            curr->next = tmp;
            curr = curr->next;
        }

        if (num == 10)
        {
            ListNode *tmp = new ListNode(1);
            curr->next = tmp;
        }
        return res.next;
    }
};

int main()
{
    cout << "Hello, world!" << endl;
    return 0;
}

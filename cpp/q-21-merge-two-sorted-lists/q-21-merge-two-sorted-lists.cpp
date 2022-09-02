#include "iostream"
using namespace std;

//21.合并两个有序链表.merge-two-sorted-lists

struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {}

    ListNode(int x) : val(x), next(nullptr) {}

    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        if (!list1) return list2;
        if (!list2) return list1;
        ListNode *t, *ans;
        if (list1->val < list2->val) {
            t = list1;
            list1 = list1->next;
        } else {
            t = list2;
            list2 = list2->next;
        }
        ans = t;
        while (list1 && list2) {
            if (list1->val < list2->val) {
                t->next = list1;
                list1 = list1->next;
            } else {
                t->next = list2;
                list2 = list2->next;
            }
            t = t->next;
        }
        if (list1) t->next = list1;
        if (list2) t->next = list2;
        return ans;
    }
};
//leetcode submit region end(Prohibit modification and deletion)


int main(){
    ListNode *l1 = nullptr, *l2 = nullptr;
    cout << (new Solution())->mergeTwoLists(l1, l2) << endl;
    return 0;
}

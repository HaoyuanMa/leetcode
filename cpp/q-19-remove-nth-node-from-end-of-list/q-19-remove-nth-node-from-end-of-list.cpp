#include "iostream"
#include "utils.cpp"
using namespace std;

//19.删除链表的倒数第 N 个结点.remove-nth-node-from-end-of-list

typedef SListNode ListNode;

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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *cur = head, *des = head;
        for (int i = 0; i <= n; ++i) {
            if (!cur) return head->next;
            cur = cur->next;
        }
        while (cur){
            cur = cur->next;
            des = des->next;
        }
        des->next = des->next->next;
        return head;
    }
};
//leetcode submit region end(Prohibit modification and deletion)


int main(){
    vector<int> v = {1};
    auto list = utils::CreateSingleList(v);
    utils::PrintList((new Solution())->removeNthFromEnd(list,1));
    return 0;
}

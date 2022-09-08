#include <vector>
#include "utils.cpp"
using namespace std;

//23.合并K个升序链表.merge-k-sorted-lists

#ifndef UTILS
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};
#endif

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
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        int n = lists.size();
        if (lists.empty()) {
            return nullptr;
        } else if (n == 1){
            return lists.at(0);
        } else if (n == 2){
            ListNode *head = new ListNode(), *cur = head;
            ListNode *left = lists.at(0), *right = lists.at(1);
            while (left && right){
                if (left->val < right->val){
                    cur->next = left;
                    left = left->next;
                } else {
                    cur->next = right;
                    right = right->next;
                }
                cur = cur->next;
            }
            if (left) cur->next = left;
            if (right) cur->next = right;
            return head->next;
        } else {
            int mid = n / 2;
            vector<ListNode*> vl(lists.begin(), lists.begin()+mid);
            vector<ListNode*> vr(lists.begin()+mid, lists.end());
            ListNode *left = mergeKLists(vl);
            ListNode *right = mergeKLists(vr);
            vector<ListNode*> vf = {left, right};
            return mergeKLists(vf);
        }
    }
};
//leetcode submit region end(Prohibit modification and deletion)


int main(){
    vector<int> va = {1,4,5};
    vector<int> vb = {1,3,4};
    vector<int> vc = {2,6};

    auto a = utils::CreateSingleList(va);
    auto b = utils::CreateSingleList(vb);
    auto c = utils::CreateSingleList(vc);

    vector<ListNode*> lists;
    lists.push_back(a);
    lists.push_back(b);
    lists.push_back(c);

    utils::PrintList((new Solution())->mergeKLists(lists));
    return 0;
}

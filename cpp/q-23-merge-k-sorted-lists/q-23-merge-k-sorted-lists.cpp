#include <queue>
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
    struct cmp {
        bool operator ()(ListNode *p, ListNode *q){
            return p->val > q->val;
        }
    };

//    class ccmp {
//    public:
//        bool operator ()(ListNode *p, ListNode *q){
//            return p->val > q->val;
//        }
//    };



    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, deque<ListNode*>, cmp> heap;
        ListNode *head = new ListNode(), *cur = head;
        for (auto list : lists) {
            if (list) heap.push(list);
        }
        while (!heap.empty()) {
            auto node = heap.top();
            heap.pop();
            cur->next = node;
            cur = cur->next;
            if (node->next) {
                heap.push(node->next);
            }
        }
        return head->next;
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

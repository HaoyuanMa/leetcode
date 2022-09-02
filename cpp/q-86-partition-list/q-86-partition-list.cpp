#include "iostream"
#include "vector"
using namespace std;

//86.分隔链表.partition-list

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
    ListNode* partition(ListNode* head, int x) {
        ListNode *p = new ListNode(), *q = new ListNode();
        ListNode *l = p, *r = q;
        while (head){
            ListNode *t = head;
            if (head->val < x){
                p->next = head;
                p = p->next;
            } else{
                q->next = head;
                q = q->next;
            }
            head = head->next;
            t->next = nullptr;
        }
        p->next = r->next;
        return l->next;
    }
};
//leetcode submit region end(Prohibit modification and deletion)


int main(){
    vector<int> v = {1, 4, 3, 2, 5, 2};
    ListNode list, *p = &list;

    for (int i = 0; i < v.size(); ++i,p = p->next) {
        p->next = new ListNode(v[i], nullptr);
    }
    cout << (new Solution())->partition(list.next, 3) << endl;
    return 0;
}

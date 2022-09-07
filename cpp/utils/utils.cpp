#include "utils.h"
using namespace std;

#ifndef UTILS
#define UTILS
#include "vector"
#include "iostream"
namespace utils {
    SListNode* CreateSingleList(vector<int> &data){
        int l = data.size();
        SListNode *head = new SListNode(), *cur = head;
        for (int i = 0; i < l; ++i) {
            cur->next = new SListNode(data[i]);
            cur = cur->next;
        }
        return head->next;
    }

    void PrintList(SListNode* list){
        auto cur = list;
        while (cur){
            std::printf("%d  ",cur->val);
            cur = cur->next;
        }
        std::printf("\n");
    }
}

#endif
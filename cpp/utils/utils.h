
#ifndef UTILS_H
#define UTILS_H

struct SListNode {
    int val;
    SListNode *next;
    SListNode() : val(0), next(nullptr) {}
    SListNode(int x) : val(x), next(nullptr) {}
    SListNode(int x, SListNode *next) : val(x), next(next) {}
};


#endif

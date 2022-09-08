#include "iostream"

using namespace std;

int fun_1(int a, int b, int (*add)(int x, int y)) {
    return add(a, b);
}

int fun_add(int x, int y) {return x + y;}

class Add{
public:
    int operator()(int x, int y){
        return x + y;
    }
};

int fun_2(int a, int b, Add){
    return Add()(a, b);
}

int main() {
    cout << fun_1(5, 3, fun_add) << endl;
    //compile fail
    //cout << fun_1(1, 2, Add)
    cout << fun_2(1, 2, Add()) << endl;
    cout << fun_2(1, 2, fun_add) << endl;
}

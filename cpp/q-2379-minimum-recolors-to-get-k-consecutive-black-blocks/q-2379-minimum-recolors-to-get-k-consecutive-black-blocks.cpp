#include "iostream"
using namespace std;

//2379.得到 K 个黑块的最少涂色次数.minimum-recolors-to-get-k-consecutive-black-blocks

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    int minimumRecolors(string blocks, int k) {
        int ans = k;
        int count = 0;
        int l = blocks.length();
        for (int i = 0; i < k; ++i)
            if (blocks[i] == 'W') count++;
        if (count < ans) ans = count;
        for (int i = 1; i < l - k + 1; ++i) {
            if (blocks[i - 1] == 'W')   count--;
            if (blocks[i + k - 1] == 'W')   count++;
            if (count < ans) ans = count;
        }
        return ans;
    }
};
//leetcode submit region end(Prohibit modification and deletion)


int main(){
    string str = "BBBBBBBWBW";
    cout << (new Solution())->minimumRecolors(str, 7) << endl;
    return 0;
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//151.反转字符串中的单词.reverse-words-in-a-string

// leetcode submit region begin(Prohibit modification and deletion)
type bytes []byte

func (bs bytes) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs bytes) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs bytes) Len() int {
	return len(bs)
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	ans := ""
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			ans += strs[i] + " "
		}
	}
	return strings.TrimSpace(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//s := scanner.Text()
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	result := reverseWords(s)
	fmt.Println(result)
}

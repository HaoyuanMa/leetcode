package main

import "fmt"

//6.Z 字形变换.zigzag-conversion

// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 {
		return s
	}
	var ans string
	unit := 2*numRows - 2
	index := 0
	for i := 0; i < numRows; i++ {
		index = i
		//tmp := ""
		if i == 0 || i == numRows-1 {
			for index < l {
				ans += string(s[index])
				index += unit
			}
		} else {
			for j := 0; j <= l/unit; j++ {
				index = j*unit + i
				if index >= l {
					break
				}
				ans += string(s[index])
				index += 2 * (numRows - i - 1)
				if index >= l {
					break
				}
				ans += string(s[index])
			}

		}
		//tmp += " "
		//ans += tmp
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s := "PAYPALISHIRING"
	result := convert(s, 4)
	fmt.Println(result)
}

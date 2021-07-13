// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

//runes int32
//wise 聪明  rune-wise 符文智慧  偏了
// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	//explain this   i = 0, j = len(r) -1。 退出条件 i到一半，这个时候字符串就遍历完了
	//i, j = i+1, j-1 高级写法  j = j-1  i = i + 1
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

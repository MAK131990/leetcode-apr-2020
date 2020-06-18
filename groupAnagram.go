//Given an array of strings, group anagrams together.
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]

package main

import (
	"fmt"
	"strconv"
)

func main() {
	groupAnagrams([]string{"hos", "boo", "nay", "deb", "wow", "bop", "bob", "brr", "hey", "rye", "eve", "elf", "pup", "bum", "iva", "lyx", "yap", "ugh", "hem", "rod", "aha", "nam", "gap", "yea", "doc", "pen", "job", "dis", "max", "oho", "jed", "lye", "ram", "pup", "qua", "ugh", "mir", "nap", "deb", "hog", "let", "gym", "bye", "lon", "aft", "eel", "sol", "jab"})
	fmt.Println(getCharTable("tea"))
}

func getArrToStr(arr [26]int) string {
	var res string
	for i := 0; i < len(arr); i++ {
		res = res + strconv.Itoa(arr[i])
	}
	return res
}

func getCharTable(str string) string {
	var t [26]int
	for i := 0; i < len(str); i++ {
		t[str[i]-97] += 1
	}
	strArr := getArrToStr(t)
	return strArr
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	length := len(strs)
	mapRes := make(map[string][]string)
	for i := 0; i < length; i++ {
		key := getCharTable(strs[i])
		mapRes[key] = append(mapRes[key], strs[i])
		fmt.Println(key, mapRes[key])
	}

	for _, v := range mapRes {
		result = append(result, v)
	}
	return result
}

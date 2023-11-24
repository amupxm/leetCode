package groupanagrams

import "fmt"

func groupAnagrams(strs []string) [][]string {
	// {97 122}//
	hashMap := map[string][]string{}
	for _, str := range strs {
		name := "" // 97:2
		for i := 97; i <= 122; i++ {
			count := 0
			for _, letter := range str {
				if int(letter) == i {
					count++
				}
			}

			if count != 0 {
				name = fmt.Sprintf("%s;%d:%d", name, count, i)
			}
		}

		hashMap[name] = append(hashMap[name], str)

	}

	resp := [][]string{}
	for _, v := range hashMap {
		resp = append(resp, v)
	}
	return resp
}

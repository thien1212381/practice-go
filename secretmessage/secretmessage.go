package secretmessage

import "sort"

// Encode func
func Encode(encoded string) string {
	countData := map[int32]int{}

	for _, c := range encoded {
		if count,ok := countData[c]; ok {
			countData[c] = count + 1
		} else {
			countData[c] = 0
		}
	}

	mapList := map[int][]int32{}
	countArr := []int{}

	for k,v := range countData {
		mapList[v] = append(mapList[v], k)
	}

	for count := range mapList {
		countArr = append(countArr, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(countArr)))

	var result string
	for _, count := range countArr {
		for _, c := range mapList[count] {
			if string(c) == "_" {
				return result
			}
			result += string(c)
		}
	}
	return result

}

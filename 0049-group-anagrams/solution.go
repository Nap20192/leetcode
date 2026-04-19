package groupanagrams

// https://leetcode.com/problems/group-anagrams/
// Difficulty: Medium

func groupAnagrams(strs []string) [][]string {
	storage := make(map[[26]int]int)
	var result [][]string
	index := 0
	for i := range strs {
		h := serialize(strs[i])
		curr, ok := storage[h]
		if !ok {
			storage[h] = index
			result = append(result, []string{})
			result[index] = append(result[index], strs[i])
			index++
			continue
		}
		result[curr] = append(result[curr], strs[i])

	}
	return result
}

func serialize(s string) [26]int {
	var count [26]int
	for i := range s {
		count[s[i]-'a']++
	}
	return count
}

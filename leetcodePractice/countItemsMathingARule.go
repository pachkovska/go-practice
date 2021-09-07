package leetcodePractice

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	keyMap := map[string]int{"type": 0, "color": 1, "name": 2}
	ruleIndex := keyMap[ruleKey]

	var counter int

	for i:=0; i < len(items); i++ {
		currentItem := items[i];
		if currentItem[ruleIndex] == ruleValue {
			counter++
		}
	}

	return counter
}

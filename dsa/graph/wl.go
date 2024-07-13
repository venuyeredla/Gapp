package graph

func ladderLength(beginWord string, endWord string, wordList []string) int {

	patMap := make(map[string][]string)
	L := len(beginWord)
	for i := 0; i < L; i++ {
		for _, word := range wordList {
			pat := word[:i] + "*" + word[i+1:]
			if _, ok := patMap[pat]; !ok {
				patMap[pat] = make([]string, 0)
			}
			patMap[pat] = append(patMap[pat], word)
		}
	}
	visited := make(map[string]bool)
	visited[beginWord] = true
	Queue := make([][]interface{}, 0, len(wordList))
	Queue = append(Queue, []interface{}{beginWord, 1})

	for len(Queue) > 0 {
		node := Queue[0]
		Queue = Queue[1:]
		word := node[0].(string)
		level := node[1].(int)
		for i := 0; i < L; i++ {
			nextPat := word[:i] + "*" + word[i+1:]

			for _, adjacent := range patMap[nextPat] {
				if adjacent == endWord {
					return level + 1
				}
				if !visited[adjacent] {
					visited[adjacent] = true
					Queue = append(Queue, []interface{}{adjacent, level + 1})
				}
			}

		}

	}

	return 0

}

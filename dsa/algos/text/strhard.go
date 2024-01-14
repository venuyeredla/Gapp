package text

import "strings"

func TextJustify(words []string, maxWidth int) []string {
	solution := make([]string, 0)
	sb := make([]string, 0)
	available := maxWidth
	for _, word := range words {
		if available < len(word) {
			solution = append(solution, justify(sb, available, false))
			available = maxWidth
			sb = make([]string, 0)
		}
		sb = append(sb, word)
		available = available - len(word) - 1
	}
	if len(sb) > 0 {
		solution = append(solution, justify(sb, available, true))
	}
	return solution
}

func justify(strs []string, scount int, isLast bool) string {
	scount = scount + len(strs)
	spaces := make([]int, len(strs))
	if !isLast {
		j := 0
		for ; scount > 0; scount-- {
			spaces[j]++
			j++
			if j >= len(spaces)-1 {
				j = 0
			}

		}
	} else {
		i := 0
		for ; i < len(spaces)-1; i++ {
			spaces[i] = 1
			scount--
		}
		spaces[i] = scount
	}

	var sb strings.Builder
	for i, ws := range strs {
		sb.WriteString(ws)
		for i := spaces[i]; i > 0; i-- {
			sb.WriteString("*")
		}

	}
	return sb.String()
}

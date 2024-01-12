package text

import "strings"

/*
Greedy algorith.
	Max=16
	avail=
	16-5 => 11-3 => 8-3 =>5-8
	4,1 => 2,1 => 2,1  = Available 5

*/

func TextJustify(words []string, maxWidth int) []string {
	strs := make([]string, 0)
	sb := make([]*WordSpace, 0)
	available := maxWidth
	for _, word := range words {
		if available < len(word) {
			strs = append(strs, justify(sb, available, false))
			sb = make([]*WordSpace, 0)
			available = maxWidth
		}
		ws := NewWS(word, 1)
		sb = append(sb, ws)
		available = available - len(word) - 1
	}
	if len(sb) > 0 {
		strs = append(strs, justify(sb, available, true))
	}
	return strs
}

type WordSpace struct {
	W string
	S int
}

func NewWS(w string, i int) *WordSpace {
	return &WordSpace{W: w, S: i}
}

func justify(WS []*WordSpace, lastspaces int, last bool) string {

	if last {
		var sb strings.Builder
		for _, ws := range WS {
			sb.WriteString(ws.W)
			sb.WriteString(" ")
		}
		sb.WriteString("|")
		return sb.String()
	}
	j := 0
	wsc := len(WS) - 1
	for i := lastspaces + 1; i >= 0; i-- {
		if j == wsc {
			j = 0
		}
		WS[j].S = WS[j].S + 1
		j++
	}
	var sb strings.Builder
	for _, ws := range WS {
		sb.WriteString(ws.W)
		for i := ws.S; i >= 0; i-- {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("|")
	return sb.String()
}

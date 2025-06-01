package regex

import (
	"unicode/utf8"

	"github.com/eapache/queue/v2"
)

type element struct {
	stateId      int
	indexToMatch int
}

func MatchString(adj map[int][]Edge, nfa *NFA, stringToMatch string) bool {
	start := nfa.Start
	n := utf8.RuneCountInString(stringToMatch)
	stringToMatchSlice := make([]rune, 0, n)
	for _, character := range stringToMatch {
		stringToMatchSlice = append(stringToMatchSlice, character)
	}
	q := queue.New[element]()
	q.Add(element{stateId: start, indexToMatch: 0})
	vis := make(map[element]bool)


	for q.Length() != 0 {
		curr_element := q.Remove()
		vis[curr_element] = true

		if curr_element.indexToMatch == n && curr_element.stateId == nfa.End {
			return true
		}

		for _, nextEdge := range adj[curr_element.stateId] {
			nextStateId := nextEdge.To
			nextIndexToMatch := -1 
			if nextEdge.Char == Epsilon {
				nextIndexToMatch = curr_element.indexToMatch
			} else {
				if curr_element.indexToMatch != n && nextEdge.Char == stringToMatchSlice[curr_element.indexToMatch] {
					nextIndexToMatch = curr_element.indexToMatch + 1
				}
			}

			next_element := element{stateId: nextStateId, indexToMatch: nextIndexToMatch}
			if next_element.indexToMatch != -1 && !vis[next_element] {
				q.Add(next_element)
			}
		}
	}

	return false
}

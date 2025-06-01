package regex

import (
	"github.com/golang-collections/collections/stack"
	"github.com/madhu102938/regex-engine/utils"
)

type NFA struct {
	Start int
	End   int
}

type Edge struct {
	To   int
	Char rune
}

const Epsilon rune = 'ε'

func BuildNFA(expression string) (*NFA, map[int][]Edge) {
	start, end := 0, 1
	s := stack.New()
	adj := map[int][]Edge{}

	for _, character := range expression {
		if !utils.IsOperator(character) {
			if s.Len() != 0 {
				start = s.Peek().(*NFA).End + 1
				end = start + 1
			}
			nfa := &NFA{Start: start, End: end}
			adj[start] = append(adj[start], Edge{To: end, Char: character})
			s.Push(nfa)
		} else {
			switch character {
			case '+':
				nfa_old := s.Pop().(*NFA)
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: nfa_old.Start, Char: Epsilon})
				new_end := nfa_old.End + 1
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: new_end, Char: Epsilon})
				nfa := &NFA{Start: nfa_old.Start, End: new_end}
				s.Push(nfa)
			case '?':
				nfa_old := s.Pop().(*NFA)
				adj[nfa_old.Start] = append(adj[nfa_old.Start], Edge{To: nfa_old.End, Char: Epsilon})
				new_start := nfa_old.End + 1
				new_end := new_start + 1
				adj[new_start] = append(adj[new_start], Edge{To: nfa_old.Start, Char: Epsilon})
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: new_end, Char: Epsilon})
				nfa := &NFA{Start: new_start, End: new_end}
				s.Push(nfa)
			case '*':
				nfa_old := s.Pop().(*NFA)
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: nfa_old.Start, Char: Epsilon})
				adj[nfa_old.Start] = append(adj[nfa_old.Start], Edge{To: nfa_old.End, Char: Epsilon})
				new_start := nfa_old.End + 1
				new_end := new_start + 1
				adj[new_start] = append(adj[new_start], Edge{To: nfa_old.Start, Char: Epsilon})
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: new_end, Char: Epsilon})
				nfa := &NFA{Start: new_start, End: new_end}
				s.Push(nfa)
			case utils.CONCAT_CHARACTER:
				nfa_top := s.Pop().(*NFA)
				nfa_bottom := s.Pop().(*NFA)
				adj[nfa_bottom.End] = append(adj[nfa_bottom.End], Edge{To: nfa_top.Start, Char: Epsilon})
				nfa := &NFA{Start: nfa_bottom.Start, End: nfa_top.End}
				s.Push(nfa)
			case '|':
				nfa_top := s.Pop().(*NFA)
				nfa_bottom := s.Pop().(*NFA)
				new_start := nfa_top.End + 1
				new_end := new_start + 1
				adj[new_start] = append(adj[new_start], Edge{To: nfa_bottom.Start, Char: Epsilon})
				adj[new_start] = append(adj[new_start], Edge{To: nfa_top.Start, Char: Epsilon})
				adj[nfa_bottom.End] = append(adj[nfa_bottom.End], Edge{To: new_end, Char: Epsilon})
				adj[nfa_top.End] = append(adj[nfa_top.End], Edge{To: new_end, Char: Epsilon})
				nfa := &NFA{Start: new_start, End: new_end}
				s.Push(nfa)
			}
		}
	}

	return s.Pop().(*NFA), adj
}

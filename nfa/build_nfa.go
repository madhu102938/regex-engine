package nfa

import (
	"github.com/golang-collections/collections/stack"
	"github.com/madhu102938/regex-engine/token"
	"github.com/madhu102938/regex-engine/utils"
)

type NFA struct {
	Start int
	End   int
}

type Character struct {
	Type int
}

type Edge struct {
	To         int
	RegexToken token.RegexToken
}

func BuildNFA(expression []token.RegexToken) (*NFA, map[int][]Edge) {
	start, end := 0, 1
	s := stack.New()
	adj := map[int][]Edge{}

	for _, character := range expression {
		if !utils.IsOperator(character.Type) {
			if s.Len() != 0 {
				start = s.Peek().(*NFA).End + 1
				end = start + 1
			}
			nfa := &NFA{Start: start, End: end}
			adj[start] = append(adj[start], Edge{To: end, RegexToken: character})
			s.Push(nfa)
		} else {
			epsilon := token.RegexToken{Type: token.Epsilon}
			switch character.Type {
			case token.Plus:
				nfa_old := s.Pop().(*NFA)
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: nfa_old.Start, RegexToken: epsilon})
				new_end := nfa_old.End + 1
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: new_end, RegexToken: epsilon})
				nfa := &NFA{Start: nfa_old.Start, End: new_end}
				s.Push(nfa)
			case token.QuestionMark:
				nfa_old := s.Pop().(*NFA)
				adj[nfa_old.Start] = append(adj[nfa_old.Start], Edge{To: nfa_old.End, RegexToken: epsilon})
				new_start := nfa_old.End + 1
				new_end := new_start + 1
				adj[new_start] = append(adj[new_start], Edge{To: nfa_old.Start, RegexToken: epsilon})
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: new_end, RegexToken: epsilon})
				nfa := &NFA{Start: new_start, End: new_end}
				s.Push(nfa)
			case token.Star:
				nfa_old := s.Pop().(*NFA)
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: nfa_old.Start, RegexToken: epsilon})
				adj[nfa_old.Start] = append(adj[nfa_old.Start], Edge{To: nfa_old.End, RegexToken: epsilon})
				new_start := nfa_old.End + 1
				new_end := new_start + 1
				adj[new_start] = append(adj[new_start], Edge{To: nfa_old.Start, RegexToken: epsilon})
				adj[nfa_old.End] = append(adj[nfa_old.End], Edge{To: new_end, RegexToken: epsilon})
				nfa := &NFA{Start: new_start, End: new_end}
				s.Push(nfa)
			case token.ConcatCharacter:
				nfa_top := s.Pop().(*NFA)
				nfa_bottom := s.Pop().(*NFA)
				adj[nfa_bottom.End] = append(adj[nfa_bottom.End], Edge{To: nfa_top.Start, RegexToken: epsilon})
				nfa := &NFA{Start: nfa_bottom.Start, End: nfa_top.End}
				s.Push(nfa)
			case token.Or:
				nfa_top := s.Pop().(*NFA)
				nfa_bottom := s.Pop().(*NFA)
				new_start := nfa_top.End + 1
				new_end := new_start + 1
				adj[new_start] = append(adj[new_start], Edge{To: nfa_bottom.Start, RegexToken: epsilon})
				adj[new_start] = append(adj[new_start], Edge{To: nfa_top.Start, RegexToken: epsilon})
				adj[nfa_bottom.End] = append(adj[nfa_bottom.End], Edge{To: new_end, RegexToken: epsilon})
				adj[nfa_top.End] = append(adj[nfa_top.End], Edge{To: new_end, RegexToken: epsilon})
				nfa := &NFA{Start: new_start, End: new_end}
				s.Push(nfa)
			}
		}
	}

	return s.Pop().(*NFA), adj
}

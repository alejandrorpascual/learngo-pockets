package gordle

import "strings"

// hint describes the validity of a character in a word.
type hint int

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "-"
	case wrongPosition:
		return "?"
	case correctPosition:
		return "+"
	default:
		// This should never happen
		return "💔"
	}
}

// Feedback is a list of hints, one per character of the word.
type Feedback []hint

// String implements the Stringer interface for a slice of hints
func (fb Feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

func (fb Feedback) Equal(other Feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for i, v := range fb {
		if v != other[i] {
			return false
		}
	}

	return true
}

func (fb Feedback) GameWon() bool {
	for _, c := range fb {
		if c != correctPosition {
			return false
		}
	}

	return true
}

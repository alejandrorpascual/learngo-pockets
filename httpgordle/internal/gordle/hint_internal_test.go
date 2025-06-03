package gordle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_feedback_String(t *testing.T) {
	testCases := map[string]struct {
		fb   Feedback
		want string
	}{
		"three correct": {
			fb:   Feedback{correctPosition, correctPosition, correctPosition},
			want: "+++",
		},
		"one of each": {
			fb:   Feedback{correctPosition, wrongPosition, absentCharacter},
			want: "+?-",
		},
		"different order for one of each": {
			fb:   Feedback{wrongPosition, absentCharacter, correctPosition},
			want: "?-+",
		},
		"unknown position": {
			fb:   Feedback{404},
			want: "💔",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := tc.fb.String(); got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFeedbackGameWon(t *testing.T) {
	tt := map[string]struct {
		fb   Feedback
		want bool
	}{
		"game not won": {
			fb:   Feedback{0, 1, 0, 0, 0},
			want: false,
		},
		"game won": {
			fb:   Feedback{2, 2, 2, 2, 2},
			want: true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.fb.GameWon()
			assert.Equal(t, tc.want, got)
		})
	}
}

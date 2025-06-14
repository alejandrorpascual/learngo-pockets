package guess

import (
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	game, _ := gordle.New("pocket")
	handle := Handler(successGameGuesserStub{session.Game{
		ID:           "123456",
		Gordle:       *game,
		AttemptsLeft: 5,
		Status:       session.StatusPlaying,
	}})

	body := strings.NewReader(`{"guess": "pocket"}`)
	req, err := http.NewRequest(http.MethodPut, "/games/", body)
	require.NoError(t, err)

	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()

	handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":4,"guesses":[{"word":"pocket","feedback":"++++++"}],"word_length":6,"status":"Won"}`, recorder.Body.String())
}

type successGameGuesserStub struct {
	game session.Game
}

func (g successGameGuesserStub) Find(id session.GameID) (session.Game, error) {
	return g.game, nil
}

func (g successGameGuesserStub) Update(game session.Game) error {
	return nil
}

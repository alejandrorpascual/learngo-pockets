package getstatus

import (
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/games", nil)
	require.NoError(t, err)

	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()
	handleFunc := Handler(gameGuesserStub{})
	handleFunc(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"solution":"","status":""}`, recorder.Body.String())
}

type gameGuesserStub struct {
	err error
}

func (g gameGuesserStub) Find(id session.GameID) (session.Game, error) {
	return session.Game{ID: id}, g.err
}
func (g gameGuesserStub) Update(session.Game) error {
	return g.err
}

package newgame

import (
	"learngo-pockets/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/games", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	handleFunc := Handler(gameAdderStub{})
	handleFunc(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":[],"word_length":0,"solution":"","status":""}`, recorder.Body.String())

}

type gameAdderStub struct {
	err error
}

func (gr gameAdderStub) Add(game session.Game) error {
	return gr.err
}
